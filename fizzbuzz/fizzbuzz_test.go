package fizzbuzz

import (
  "testing"
  "strings"
)

func slicesEq(s1, s2 []string) bool{
  if s1 == nil && s2 == nil { return true }

  if s1 == nil || s2 == nil { return false }

  if len(s1) != len(s2) { return false }

  for i := range s1 {
    if (s1[i] != s2[i]) { return false }
  }

  return true
}

func TestBasicFizzBuzz(t *testing.T){
  actual := Generate(3, 5, 20, "Fizz", "Buzz")
  expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "Fizz Buzz", "16", "17", "Fizz", "19", "Buzz"}

  if slicesEq(actual, expected) == false {
    t.Fatalf("\r\nexpected [%s]\r\n but got [%s] instead", strings.Join(expected, ","), strings.Join(actual, ","))
  }
}
