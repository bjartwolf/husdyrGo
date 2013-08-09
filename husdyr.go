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
  persons := sort.StringSlice{"ukranian", "norwegian", "japanese", "spaniard", "englishman"}
  drinks := sort.StringSlice{"milk", "tea", "water", "orangejuice", "coffee"}
  colors := sort.StringSlice{"yellow", "blue", "red", "green", "ivory"}
/*  for _, person := range persons {
    fmt.Println(person)
  }
*/
  mathutil.PermutationFirst(persons)
  mathutil.PermutationFirst(drinks)
  mathutil.PermutationFirst(colors)
  for personPerm := true; personPerm ; personPerm = mathutil.PermutationNext(persons) {
    // Rule #1: Norweigan lives in first house 
    if (persons[0] != "norwegian") {
      continue
    }
    for colorsPerm:= true; colorsPerm; colorsPerm= mathutil.PermutationNext(colors){
      // #Rule #2: Englishman lives in red house
      if (indexOf(colors, "red") != indexOf(persons, "englishman")) {
        continue
      }
      // #Rule #6: Green house is to the right of ivory house
      if (indexOf(colors, "green") - indexOf(colors, "ivory") != 1) {
        continue
      }

      for drinksPerm := true; drinksPerm; drinksPerm =  mathutil.PermutationNext(drinks){
      // Person i siste huset (n=2) drinksr ikke te
//      if (drinks[2] == "te") {
//        continue
//      }
      // Arne drinksr te
//      if (indexOf(drinks, "te") != indexOf(persons, "arne")) {
//        continue
//      }
      fmt.Println("******************")
      fmt.Println(persons)
      fmt.Println(drinks)
      fmt.Println(colors)
      }
    }
  }
}
