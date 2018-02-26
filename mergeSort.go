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
func (st *MergeSort) Sort(a []string,lo int,hi int){
	if(hi<=lo){
		return
	}
	mid := lo +(hi-lo)/2

	st.Sort(a,lo,mid)
	st.Sort(a,mid+1,hi)
	st.Merge(a,lo,mid,hi)
}

func main()  {
	element := MergeSort{}
	var a []string
	a = []string{"M","E","R","G","E","S","O","R","T","E","X","A","M","P","L","E"};
	element.SortInit(a)
	element.Show(a)
}
