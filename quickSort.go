/*
	快速排序
 */

package main

import (
	"elementarySort"
	"fmt"
)

type QuickSort struct {
	elementarySort.ElementarySort
}

func (st *QuickSort) SortInit(a []string){
	st.Sort(a,0,len(a)-1)

}

func (st *QuickSort) Sort(a []string,lo int,hi int){
	if(hi<=lo){
		return
	}
	fmt.Print("传入",lo," ",hi," ")
	var j int
	j = st.Partition(a,lo,hi)
	fmt.Println(j)
	st.Sort(a,lo,j-1)
	st.Sort(a,j+1,hi)
}

//规则
func (st *QuickSort) Partition(a []string,lo int,hi int) int{
	var i , j int
	var v string
	i = lo
	j = hi+1
	fmt.Print("i和j：",i," ",j,"  ")

	v = a[lo]

	for{
		for{
			i++
			if(!st.Less(a[i],v)){
				break
			}
			if(i == hi){
				break
			}
		}
		for{
			j--
			if(!st.Less(v,a[j])){
				break
			}
			if(j == lo){
				break
			}
		}
		if(i >= j){
			break
		}
		st.Exch(a,i,j)
	}
	st.Exch(a,lo,j)
	return j
}

func main()  {
	element := QuickSort{}
	var a []string
	a = []string{"K","R","A","T","E","L","E","P","U","I","M","Q","C","X","O","S"};

	//自上而下的归并排序
	element.SortInit(a)
	element.Show(a)
	//A C E E I K L M O P Q R S T U X
}
