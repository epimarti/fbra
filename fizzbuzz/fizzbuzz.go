package fizzbuzz

import (
  "strings"
  "strconv"
)

func Generate(int1 int, int2 int, limit int, s1 string, s2 string) []string {
  result := make([]string, limit)

  for i := 0; i < limit; i++ {
    current := i + 1
    var sb []string
    if current % int1 == 0 {
      sb = append(sb, s1)
    }
    if current % int2 == 0 {
      sb = append(sb, s2)
    }

    if len(sb) > 0 {
      result[i] = strings.Join(sb, " ")
    } else {
      result[i] = strconv.Itoa(current)
    }
  }

  return result
}
