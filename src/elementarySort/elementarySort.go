package elementarySort

import (
	"fmt"
	"strings"
)

type ElementarySort struct {
	//a []string
}

func (st *ElementarySort) Less(v string, w string) bool {
	return strings.Compare(v, w) < 0
}

//数据交换
func (st *ElementarySort) Exch(a []string, i int, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func (st *ElementarySort) Show(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i] + " ")
	}
	fmt.Println()
}

func (st *ElementarySort) IsSorted(a []string) bool {
	for i := 1; i < len(a); i++ {
		if st.Less(a[i], a[i-1]) == true {
			return false
		}
	}
	return true
}
