package main

import "fmt"
import "github.com/cznic/mathutil"
import "sort"

func indexOf (array sort.StringSlice, lookingFor string) int {
	for i, value := range array {
		if (value == lookingFor) {
			return i
		}
	}
	return -1
}

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
		// Arne bor i det første huset
//		if (persons[0] != "arne") {
//			continue
//		}
		mathutil.PermutationFirst(drikke)
		for drikk, drikkePerm := 0, true; drikkePerm; drikk, drikkePerm = drikk + 1,  mathutil.PermutationNext(drikke){
			// Person i siste huset (n=2) drikker ikke te
			if (drikke[2] == "te") {
				continue
			}
			// Arne drikker te
			if (indexOf(drikke, "te") != indexOf(persons, "arne")) {
				continue
			}
			fmt.Println("******************")
			fmt.Println(persons)
			fmt.Println(drikke)
		}
	}
}
