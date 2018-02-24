package main

import (
	"elementarySort"
) 

type InsertionSort struct {
	elementarySort.ElementarySort
}

func (st *InsertionSort) Sort(a []string){
	var n int
	var i int
	n = len(a)

	for i=0;i<n ;i++  {
		for j := i;j>0 && st.Less(a[j],a[j-1]) ;j--  {
			st.Exch(a,j,j-1)
		}
	}

}

func main()  {
	element := InsertionSort{}
	var a []string
	a = []string{"S","O","R","T","E","X","A","M","P","L","E"};
	element.Sort(a)
	element.Show(a)
}
