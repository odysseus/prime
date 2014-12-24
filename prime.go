package main

import (
  "fmt"
  "math"
  "strconv"
  "flag"
)

func main() {
  // Define command line flags
  fact := flag.Bool("f", false, "Displays prime factors for non-prime numbers")

  // Parse the flags
  flag.Parse()

  // Get the non-flag arguments
  args := flag.Args()

  // Run primality checks on each
  for _, v := range args {
    a, err := strconv.Atoi(v)
    if err != nil {
      fmt.Printf("%s cannot be converted to an integer\n", v)
    } else {
      if IsPrime(a) {
        fmt.Printf("%v is prime\n", a)
      } else {
        if *fact {
          factors := PrimeFacts(a)
          fmt.Printf("%v is not prime, prime factors: %v\n", a, factors)
        } else {
          fmt.Printf("%v is not prime\n", a)
        }
      }
    }
  }
}

// Returns a boolean indicating whether the
// number is prime
func IsPrime(n int) bool {
  if n <= 1 { return false }
  if n == 2 || n == 3 { return true }
  if n % 2 == 0  || n % 3 == 0 { return false }
  if n < 9 { return true }
  r := intRoot(n)
  s := TwoThreeSieve()
  for i:=s(); i<=r; i=s() {
    if n % i == 0 { return false }
  }
  return true
}

// Returns the square root, truncated to an int + 1
// For the express purpose of having a stop point
// for testing primality
func intRoot(n int) int {
  return int(math.Sqrt(float64(n))) + 1
}

// Returns all numbers above three that are
// not multiples of two or three
func TwoThreeSieve() func() int {
  i := 5
  two, first := true, true
  return func() int {
    if first {
      first = false
      return 5
    }
    if two {
      two = false; i += 2
      return i
    } else {
      two = true; i += 4
      return i
    }
  }
}

// Finds all prime factors of n and returns
// them in a slice
func PrimeFacts(n int) []int {
  factors := make([]int, 0, 0)
  if (n % 2 == 0) {
    factors = append(factors, 2)
    for n % 2 == 0 { n /= 2 }
  }
  for i:=3; i<=n; i+=2 {
    if n % i == 0 {
      factors = append(factors, i)
      for n % i == 0 { n /= i }
    }
  }
  return factors
}
