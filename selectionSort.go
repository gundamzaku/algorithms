package main

import (
	"elementarySort"
)

type SelectionSort struct {
	elementarySort.ElementarySort
}

func (st *SelectionSort) Sort(a []string){
	var n int
	var i int
	n = len(a)

	for i=0;i<n ;i++  {
		min := i
		for j := i+1;j<n ;j++  {
			if(st.Less(a[j],a[min])){
				min = j
			}
			st.Exch(a,i,min)
		}
	}

}

func main()  {
	element := SelectionSort{}
	var a []string
	a = []string{"S","O","R","T","E","X","A","M","P","L","E"};
	element.Sort(a)
	element.Show(a)
}
