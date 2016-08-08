package zen

import (
  "testing"
  "reflect"
)

func TestContainsString(t *testing.T) {
  intSlice := []int{0, 1, 2, 3}
  stringSlice := IntToStringSlice(intSlice)

  if ContainsString(stringSlice, "1") != true {
    t.Error("Expected slice to include '1'")
  }

  if ContainsString(stringSlice, "0") != true {
    t.Error("Expected slice to include '0'")
  }
}

func TestIntToStringSlice(t *testing.T) {
  intSlice := []int{0, 1, 2, 3}
  stringSlice := []string{"0", "1", "2", "3"}
  transformedSlice := IntToStringSlice(intSlice)

  if !reflect.DeepEqual(stringSlice, transformedSlice) {
    t.Error("Expected slice to be transformed to string slice, but it was", transformedSlice)
  }
}

func TestContains(t *testing.T) {
  intSlice := []int{0, 1, 2, 3}

  if Contains(intSlice, 3 ) != true {
    t.Error("Expected slice to include 3")
  }
}

func TestToString(t *testing.T) {
  intSlice := []int{0, 1, 2, 3}
  expectedString := "0 1 2 3 "

  if toString := ToString(intSlice); toString != expectedString {
    t.Error("Expected processed string, but got", toString)
  }
}

func TestMaxValue(t *testing.T) {
  testHash := map[int]int {1: 20, 2: 15, 3: 10}

  if maxValue := MaxValue(testHash); maxValue != 20{
    t.Error("Expected max value to be 20, but it was", maxValue)
  }
}

func TestMaxValueKey(t *testing.T) {
  testHash := map[int]int {1: 20, 2: 15, 3: 10}

  if maxValueKey := MaxValueKey(testHash); maxValueKey != 1{
    t.Error("Expected max value to be 20, but it was", maxValueKey)
  }
}

func TestMinValue(t *testing.T) {
  testHash := map[int]int {1: 20, 2: 15, 3: 10}

  if minValue := MinValue(testHash); minValue != 10{
    t.Error("Expected max value to be 20, but it was", minValue)
  }
}

func TestMinValueKey(t *testing.T) {
  testHash := map[int]int {1: 20, 2: 15, 3: 10}

  if minValueKey := MinValueKey(testHash); minValueKey != 3{
    t.Error("Expected max value to be 20, but it was", minValueKey)
  }
}

func TestIncludesString(t *testing.T) {
  stringSlice := []string{"0", "1", "2", "3"}

  if IncludesString("1", stringSlice)!= true{
    t.Error("Expected slice to include '1'")
  }
}
