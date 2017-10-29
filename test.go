package main

import (
  "fmt"
)

type Test struct {
  key_value map[string]string
  json_nested map[string]*Test
}

func main() {
  test := new(Test)
  test.key_value = make(map[string]string)
  test.key_value["test1"] = "test11"
  test.key_value["test2"] = "test22"

  test2 := new(Test)
  test.json_nested = make(map[string]*Test)
  test.json_nested["test3"] = test2
  test.json_nested["test3"].key_value = make(map[string]string)
  test.json_nested["test3"].key_value["test4"] = "test44"

  test3 := new(Test)
  test.json_nested["test3"].json_nested = make(map[string]*Test)
  test.json_nested["test3"].json_nested["test5"] = test3

  fmt.Printf("%+v\n", test)
}
