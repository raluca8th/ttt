package zen

import "strconv"

func Contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}

func ToString(slice []int) string {
  var printable string
  for _, value := range slice {
    printable += strconv.Itoa(value) + " "
  }
  return printable
}

func MaxValue(g map[int]int) int{
  i := -10
  for _, value := range g{
    if i < value {
      i = value
    }
  }
  return i
}

func MaxValueKey(g map[int]int) int{
  maxValue := MaxValue(g)
  for key, value := range g{
    if value == maxValue {
      return key
    }
  }
  return -1
}

func MinValue(g map[int]int) int{
  i := 10
  for _, value := range g{
    if i > value {
      i = value
    }
  }
  return i
}

func MinValueKey(g map[int]int) int{
  minValue := MinValue(g)
  for key, value := range g{
    if value == minValue {
      return key
    }
  }
  return -1
}

func IncludesString(s string, collection []string) bool{
  for _, value := range collection {
    if value == s {
      return true
    }
  }
  return false
}
