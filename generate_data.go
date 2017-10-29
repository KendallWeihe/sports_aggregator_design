package main

import (
    "fmt"
    "io/ioutil"
)

type JSONList struct {
  values []string // list of values
  json_objs []*JSON // list of json objects
}

type JSON struct {
  key_value map[string]string // key/value map
  json_nested map[string]*JSON // key/json object map
  json_list map[string]*JSONList // key/json list map
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func read_json(path string, custom_json *JSON) bool {
  raw, err := ioutil.ReadFile("./config.json")
  check(err)

  start_index := 0
  end_index := len(raw)
  raw_str := string(raw)
  construct_json(custom_json, &raw_str, start_index, end_index)
  return true
}

func find_closing_bracket(raw *string, opening_bracket byte, closing_bracket byte, start_index int) int {
  index := start_index
  num_opening := 0
  num_closing := 0
  for index < len(*raw) {
    if (*raw)[index] == opening_bracket {
      num_opening += 1
    } else if (*raw)[index] == closing_bracket {
      num_closing += 1
      if num_opening == num_closing {
        return index
      }
    }
    index += 1
  }
  return -1
}

func find_specific_delim(raw *string, delim byte, start_index int) int {
  index := start_index
  for index < len(*raw) {
    if (*raw)[index] == delim {
      return index
    }
    index += 1
  }
  return len(*raw)
}

func construct_json_list(json_list *JSONList, raw *string, start_index int, end_index int) interface{} {
  index := start_index
  for index < end_index {
    value_opening_quote := find_specific_delim(raw, '"', index+1)

    if value_opening_quote >= end_index {
      return true
    }

    next_curly_bracket := find_specific_delim(raw, '{', index+1)

    if value_opening_quote < next_curly_bracket {
      value_closing_quote := find_specific_delim(raw, '"', value_opening_quote+1)
      value := (*raw)[value_opening_quote+1:value_closing_quote]
      json_list.values = append(json_list.values, value)
      index = value_closing_quote
    } else if next_curly_bracket < value_opening_quote {
      closing_bracket := find_closing_bracket(raw, '{', '}', next_curly_bracket)
      json_objs := new(JSON)
      construct_json(json_objs, raw, next_curly_bracket, closing_bracket)
      json_list.json_objs = append(json_list.json_objs, json_objs)
      index = closing_bracket
    } else {
      index += 1
    }
  }
  return true
}

func construct_json(custom_json *JSON, raw *string, start_index int, end_index int) interface{} {
  custom_json.key_value = make(map[string]string)
  custom_json.json_nested = make(map[string]*JSON)
  custom_json.json_list = make(map[string]*JSONList)
  index := start_index
  for index < end_index {
    key_opening_quote := find_specific_delim(raw, '"', index+1)

    if key_opening_quote >= end_index {
      return true
    }

    key_closing_quote := find_specific_delim(raw, '"', key_opening_quote+1)
    key := (*raw)[key_opening_quote+1:key_closing_quote]
    colon := find_specific_delim(raw, ':', key_closing_quote+1)

    next_quote := find_specific_delim(raw, '"', colon+1)
    next_curly_bracket := find_specific_delim(raw, '{', colon+1)
    next_sq_bracket := find_specific_delim(raw, '[', colon+1)

    // simple key/value
    if next_quote < next_curly_bracket && next_quote < next_sq_bracket {
      value_closing_quote := find_specific_delim(raw, '"', next_quote+1)
      value := (*raw)[next_quote+1:value_closing_quote]
      custom_json.key_value[key] = value
      index = value_closing_quote
    } else if next_curly_bracket < next_sq_bracket {
      closing_bracket := find_closing_bracket(raw, '{', '}', next_curly_bracket)
      nested_json := new(JSON)
      construct_json(nested_json, raw, next_curly_bracket, closing_bracket)
      custom_json.json_nested[key] = nested_json
      index = closing_bracket
    } else if next_sq_bracket < next_curly_bracket {
      closing_bracket := find_closing_bracket(raw, '[', ']', next_sq_bracket)
      json_list := new(JSONList)
      construct_json_list(json_list, raw, next_sq_bracket, closing_bracket)
      custom_json.json_list[key] = json_list
      index = closing_bracket
    } else {
      index += 1
    }
  }

  return true
}

/*
  flow of control:
    - read config file
    - iterate over input files:
      - read input file
      - iterate over data list:
        - generate row
      - write rows to output file
*/
func main() {

  custom_json := new(JSON)
  read_json("config.json", custom_json)
  fmt.Printf("%+v\n", custom_json)

}
