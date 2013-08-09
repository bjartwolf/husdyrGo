package main

import "fmt"
import "github.com/cznic/mathutil"
import "sort"

func main () {
	persons := sort.StringSlice{"arne", "beb", "cecilie"}
	drikke := sort.StringSlice{"kaffe", "te", "pils"}
/*	for _, person := range persons {
		fmt.Println(person)
	}
*/
	mathutil.PermutationFirst(persons)
//	mathutil.PermutationFirst(drikke)
	for personPerm, p := true, 0; personPerm ; personPerm, p = mathutil.PermutationNext(persons), p + 1 {
//		if (persons[1] == "arne") {
//			continue;
//		}
		mathutil.PermutationFirst(drikke)
		for drikk, drikkePerm := 0, true; drikkePerm; drikk, drikkePerm = drikk + 1,  mathutil.PermutationNext(drikke){
			fmt.Println("******************")
			fmt.Println(persons)
			fmt.Println(drikke)
		}
	}
}
