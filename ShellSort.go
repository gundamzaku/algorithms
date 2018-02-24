/*
	希尔排序算法
	这个算法开始有点搞脑子了，重点是分块
 */
package main

import (
	"elementarySort"
)

type ShellSort struct {
	elementarySort.ElementarySort
}

func (st *ShellSort) Sort(a []string){
	var n,h,i,j,dfBlockNum int
	//配置变量，控制到底分成几块进行排序查找，可以更改
	dfBlockNum = 3//5
	n = len(a)
	h = 1
	//golang没有while，用for来处理
	for{
		if h < n/dfBlockNum {
			h = dfBlockNum*h+1
		}else{
			break
		}
	}

	for{
		if(h<1){
			break
		}

		for i = h; i<n;i++  {
			for j = i;j>=h && st.Less(a[j],a[j-h]) ;j-=h  {
				st.Exch(a,j,j-h)
			}
		}
		h = h/dfBlockNum
	}

}

func main()  {
	element := ShellSort{}
	var a []string
	a = []string{"S","H","E","L","L","S","O","R","T","E","X","A","M","P","L","E"};
	element.Sort(a)
	element.Show(a)
}
