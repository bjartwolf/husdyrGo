package main
// TODO: Could probably reuse variables and save on GC
import "fmt"
import "github.com/cznic/mathutil"
import "sort"
import "testing"

func Abs (num int) int {
  if (num < 0) {
    return -num
  } else {
    return num
  }
}

func indexOf (array sort.StringSlice, lookingFor string) int {
  for i, value := range array {
    if (value == lookingFor) {
      return i
    }
  }
  return -1
}

func solve() {
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
      for d := true; d; d = mathutil.PermutationNext(drinks){
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
        sigarettes := sort.StringSlice{"kools", "chesterfield", "oldgold", "parliament", "luckystrike"}
        mathutil.PermutationFirst(sigarettes)
        for s:= true; s; s = mathutil.PermutationNext(sigarettes){
          // Rule #8 Yellow house smokes kools
          if (indexOf(sigarettes, "kools") != indexOf(colors, "yellow")) {
            continue
          }
          // Rule #13 Lucky strike smoker drinks orange juice 
          if (indexOf(sigarettes, "luckystrike") != indexOf(drinks, "orangejuice")) {
            continue
          }
          animals := sort.StringSlice{"zebra", "snails", "dog", "horse", "fox"}
          mathutil.PermutationFirst(animals)
          for s:= true; s; s = mathutil.PermutationNext(sigarettes){
            for a:= true; a; a = mathutil.PermutationNext(animals) {
              // Rule #11 The man who smokes Chesterfields lives in the house next to the man with the fox.
              if (Abs(indexOf(sigarettes, "chesterfield") - indexOf(animals, "fox")) != 1) {
                continue
              }
              // Rule #15 The Norwegian lives next to the blue house 
              if (Abs(indexOf(persons, "norwegian") - indexOf(colors, "blue")) != 1) {
                continue
              }
              // Rule #12 Kools are smoked in the house next to the house where the horse is kept
              if (Abs(indexOf(sigarettes, "kools") - indexOf(animals, "horse")) != 1) {
                continue
              }
              // Rule #7 The Old Gold smoker owns snails
              if (indexOf(sigarettes, "oldgold") != indexOf(animals, "snails")) {
                continue
              }
              // Rule #3 The Spaniard owns the dog
              if (indexOf(persons, "spaniard") != indexOf(animals, "dog")) {
                continue
              }
              // Rule #14 The Japanese smokes Parliaments
              if (indexOf(persons, "japanese") != indexOf(sigarettes, "parliament")) {
                continue
              }
/*
              fmt.Println("******************")
              fmt.Println(persons)
              fmt.Println(drinks)
              fmt.Println(colors)
              fmt.Println(sigarettes)
              fmt.Println(animals)
*/
            }
          }
        }
      }
    }
  }
}

func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        solve()
    }
}

func main () {
    br := testing.Benchmark(BenchmarkFunction)
    fmt.Println(br)
}
