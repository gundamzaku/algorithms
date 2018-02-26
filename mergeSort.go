/*
	归并排序（其实应该就是算法导论里面分治排序）
	有点烧脑子，思路就是要带点二叉树的概念来看这个代码了。
	重点是拆成小块，然后一块一块排序，最后合并成大块数字串
	"M","E","R","G","E","S","O","R","T","E","X","A","M","P","L","E"
	拆成
	"M","E","R","G","E","S","O","R"
	"T","E","X","A","M","P","L","E"
	……后面类推
 */

package main

import (
	"elementarySort"
	"fmt"
)

type MergeSort struct {
	elementarySort.ElementarySort
}

var aux [16]string

func (st *MergeSort) Merge(a []string,lo int,mid int,hi int){

	i:= lo
	j:=mid+1

	for k:=lo;k<=hi ;k++  {
		aux[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		if(i>mid){
			a[k] = aux[j]
			j++
		}else if(j>hi){
			a[k] = aux[i]
			i++
		}else if(st.Less(aux[j],aux[i])){
			a[k] = aux[j]
			j++
		}else{
			a[k] = aux[i]
			i++
		}
	}
}
func (st *MergeSort) SortInit(a []string){
	st.Sort(a,0,len(a)-1)

}

func (st *MergeSort) SortBu(a []string){
	n := len(a)
	for sz := 1; sz < n; sz = sz + sz {
		fmt.Println("sz=>",sz)
		for lo := 0; lo < n-sz; lo += sz + sz {
			fmt.Println("lo=>",lo)
			st.Merge(a,lo,lo+sz-1,Min(lo+sz+sz-1,n-1))
		}
	}
}

func (st *MergeSort) Sort(a []string,lo int,hi int){
	if(hi<=lo){
		return
	}
	mid := lo +(hi-lo)/2

	st.Sort(a,lo,mid)
	st.Sort(a,mid+1,hi)
	st.Merge(a,lo,mid,hi)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main()  {
	element := MergeSort{}
	var a []string
	//a = []string{"E","E","G","M","R","A","C","E","R","T"}
	a = []string{"M","E","R","G","E","S","O","R","T","E","X","A","M","P","L","E"};

	//自上而下的归并排序
	element.SortInit(a)
	element.Show(a)

	//自下而上的归并排序
	element.SortBu(a)
	element.Show(a)
}
