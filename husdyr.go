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
  mathutil.PermutationFirst(persons)
  for personPerm := true; personPerm ; personPerm = mathutil.PermutationNext(persons) {
    // Rule #1: Norweigan lives in first house 
    if (persons[0] != "norwegian") {
      continue
    }
    colors := sort.StringSlice{"yellow", "blue", "red", "green", "ivory"}
    mathutil.PermutationFirst(colors)
    for colorsPerm:= true; colorsPerm; colorsPerm= mathutil.PermutationNext(colors){
      // Rule #2: Englishman lives in red house
      if (indexOf(colors, "red") != indexOf(persons, "englishman")) {
        continue
      }
      // Rule #6: Green house is to the right of ivory house
      if (indexOf(colors, "green") - indexOf(colors, "ivory") != 1) {
        continue
      }
      drinks := sort.StringSlice{"milk", "tea", "water", "orangejuice", "coffee"}
      mathutil.PermutationFirst(drinks)
      for drinksPerm := true; drinksPerm; drinksPerm =  mathutil.PermutationNext(drinks){
        // Rule #5 Ukranian drinks tea 
        if (indexOf(drinks, "tea") != indexOf(persons, "ukranian")) {
          continue
        }
        // Rule #4 Coffee is drunk in the green house
        if (indexOf(drinks, "coffee") != indexOf(colors, "green")) {
          continue
        }
        // Rule #9 Milk is drunk in the middle house (n=2)
        if (indexOf(drinks, "milk") != 2) {
          continue
        }
        fmt.Println("******************")
        fmt.Println(persons)
        fmt.Println(drinks)
        fmt.Println(colors)
      }
    }
  }
}
