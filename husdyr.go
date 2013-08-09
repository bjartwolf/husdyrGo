package main

import "fmt"
import "github.com/cznic/mathutil"
import "sort"

func main () {
	persons := sort.StringSlice{"arne", "beb", "cecilie"}
/*	drikke := []string{"kaffe", "te", "pils"}
	for _, person := range persons {
		fmt.Println(person)
	}
*/
	mathutil.PermutationFirst(persons)
	for morePermutations := true ;morePermutations; morePermutations = mathutil.PermutationNext(persons) {
		fmt.Println(persons)
	}
}
