/*
题目：
解开隐秘的等式WWWDOT - GOOGLE = DOTCOM，每个字母代表十个阿拉伯数字中的一个且各不相同，每个六位数字最高位都不是0

结果：
777589 - 188106 = 589483
777589 - 188103 = 589486
穷举次数：
738111110
 */

package main

import "fmt"

var times int64
var val1 int32
var val2 int32
var val3 int32

func main()  {
	//WWWDOT-GOOGLE = DOTCOM
	var data[] int32
	times = 0;
	//W0 G1 D2 O3 T4 L5 E6 C7 M8
	data = []int32{0,0,0,0,0,0,0,0,0}

	loop(data,8)
	//find(data)
	fmt.Println("穷举次数：",times)
}

func loop(data[] int32,len int8) bool {
	var i int32
	var offset int32
	if(len<=3){
		offset = 1
	}else{
		offset = 0
	}
	if(len<0){
		return false;
	}
	for i=offset;i<=9 ;i++  {
		data[len] = i
		loop(data,len-1)
		find(data)
		//fmt.Println(data)
	} 
	return true;
}

func find(data[] int32) bool {
	times++
	if(data[0] == 0 || data[1] == 0 || data[2] == 0){
		return false;
	}

	val1 = data[0]*100000 + data[0]*10000 + data[0]*1000 + data[2]*100 + data[3]*10 +data[4]
	val2 = data[1]*100000 + data[3]*10000 + data[3]*1000 + data[1]*100 + data[5]*10 +data[6]
	val3 = data[2]*100000 + data[3]*10000 + data[4]*1000 + data[7]*100 + data[3]*10 +data[8]
	if(val1 - val2 == val3){

		for index, value := range data {
			for i:=0;i<=8 ;i++  {
				if(i == index) {
					continue;
				}
				if (value == data[i]) {
					return false
				}
			}
		}

		fmt.Println(val1,"-",val2,"=",val3)
		return true
	}
	return false
}
