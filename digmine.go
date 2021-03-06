package main

import (
	"fmt"
	"encoding/binary"
	"time"
	"strings"
	"encoding/hex"
	"crypto/sha256"
	"bytes"
)

func swapOrder(in string) string {
	var split []string
	var x string
	x = ""
	split = strings.Split(reverseString(in),"")
	for i:=0;i<len(split) ;i+=2  {
		x+= split[i+1]+split[i]
	}
	return x
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func littleEndian(value int64) string{
	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(value))
	str := fmt.Sprintf("%x",testBytes)
	return str
}

func getUnixTime(toBeCharge string) int64  {

	//待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("UTC")
	//重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
	//使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()
	return sr
}

func hex2bin(value string)[]byte  {
	dstTmp := make([]byte, hex.EncodedLen(len(value)))
	hex.Decode(dstTmp, []byte(value))
	//真他妈坑，和PHP的不一样，这个Byte还要裁剪，否则算出来的值不对
	dst:=bytes.TrimRight(dstTmp,"\x00")
	return dst
}

func enSha256(str []byte) string {

	//b9d751533593ac10cdfb7b8e03cad8babc67d8eaeac0a3699b82857dacac9390
	hash := sha256.New()
	hash.Write(str)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr;
}

//比特币头部的结构参数，用来进行挖矿运算的
type configBitHead struct {
	prevBlockHash string
	rootHash string
	version string
	time string
	bits int64
	nonce int64
}

func main(){

	var version string
	var time string
	var bits string
	var nonce string
	var prevBlockHash string
	var rootHash string
	var headerHex string
	var pass1 string
	var pass2 string
	var finalHash string

	var config configBitHead
	config.version = "20000000"
	config.time = "2018-02-22 08:41:00"
	config.bits = 392009692
	config.nonce = 1111548121
	config.prevBlockHash = "00000000000000000010ddc4fcbf331d5000c431aa6aa474bf38e85f9f1f5788"
	config.rootHash  = "4018fe3dd8f08ba3fcaf26819cbcdc0a6b3cddd624ef6f07897daf29734b5826"

	version = swapOrder(config.version)
	time = littleEndian(getUnixTime(config.time))
	bits = littleEndian(config.bits)
	nonce = littleEndian(config.nonce)
	prevBlockHash = swapOrder(config.prevBlockHash)
	rootHash = swapOrder(config.rootHash)

	headerHex = version+prevBlockHash+rootHash+time+bits+nonce
	dst := hex2bin(headerHex)
	pass1 = string(hex2bin(enSha256(dst)))

	pass2 = enSha256([]byte(pass1))

	finalHash = swapOrder(pass2)

	fmt.Println("最后算出的哈希值：",finalHash)
}
