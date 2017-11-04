// TODO:
//   - if key contains list, return a list of values

package main

import (
    "io/ioutil"
    "strings"
    "fmt"
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
  raw, err := ioutil.ReadFile(path)
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
      custom_json.key_value[key] = value // TODO: lowercase?
      index = value_closing_quote
      // fmt.Printf(key)
      // fmt.Printf(value)
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

// OUTPUTS:
//   - either:
//     - value (string)
//     - embedded json (JSON)
//     - json list (JSONList)
func find_json_recursive(custom_json *JSON, keys []string, key_index int) (*string, *JSON, *JSONList) {

  empty_str := new(string)
  empty_json := new(JSON)
  empty_json_list := new(JSONList)

  // value found
  if value, ok := custom_json.key_value[keys[key_index]]; ok {
    return &value, empty_json, empty_json_list
  }

  // nested json found
  if nested_json, ok := custom_json.json_nested[keys[key_index]]; ok {
    // if the key requested is in fact a JSON object
    if key_index == (len(keys)-1) {
      return empty_str, nested_json, empty_json_list
    }
    // keep searching in nested JSON
    return find_json_recursive(nested_json, keys, key_index+1)
  }

  // json list
  if json_list, ok := custom_json.json_list[keys[key_index]]; ok {
    // if the key is the last key in the key_string
    if key_index == (len(keys)-1) {
      return empty_str, empty_json, json_list
    }
    // if there is more to find ... TODO this is shit?
    json_objs := custom_json.json_list[keys[key_index]].json_objs
    for _, json_obj := range json_objs {
      return find_json_recursive(json_obj, keys, key_index+1)
    }
  }

  return empty_str, empty_json, empty_json_list
}

/*
  INPUTS:
    - key in the form
      - value is from an object inside a list: "path.to.key.with.[list].support"
      - value is an actual a list: "path.to.key.with.[list]"
  OUTPUTS:
    - value
*/
func find(custom_json *JSON, key string) (*string, *JSON, *JSONList) {
  keys := strings.Split(key, ".")
  value, json, list := find_json_recursive(custom_json, keys, 0)
  return value, json, list
}

func get_indent(indent_count int) string {
  indent := ""
  i := 0
  for i < indent_count {
    indent += "\t"
    i += 1
  }
  return indent
}

func write_json_list(list JSONList, indent_count int) string {
  indent := get_indent(indent_count)
  output_str := fmt.Sprintf("%s[\n", indent)
  indent_count += 1
  indent = get_indent(indent_count)

  for _, v := range list.values {
    output_str += fmt.Sprintf("%s\"%s\",\n", indent, v)
  }

  for _, json := range list.json_objs {
    output_str += fmt.Sprintf("%s\n%s,\n", indent, write_json(*json, indent_count+1))
  }

  indent_count -= 1
  indent = get_indent(indent_count)
  output_str += fmt.Sprintf("%s]", indent)
  return output_str
}

func write_json(custom_json JSON, indent_count int) string {

  indent := get_indent(indent_count)
  output_str := fmt.Sprintf("%s{\n", indent)
  indent_count += 1
  indent = get_indent(indent_count)

  for k, v := range custom_json.key_value {
    output_str += fmt.Sprintf("%s\"%s\": \"%s\",\n", indent, k, v)
  }

  for k, json := range custom_json.json_nested {
    output_str += fmt.Sprintf("%s\"%s\":\n%s,\n", indent, k, write_json(*json, indent_count+1))
  }

  for k, list := range custom_json.json_list {
    output_str += fmt.Sprintf("%s\"%s\":\n%s,\n", indent, k, write_json_list(*list, indent_count+1))
  }

  indent_count -= 1
  indent = get_indent(indent_count) // TODO: remove the last comma
  output_str += fmt.Sprintf("%s}", indent)
  return output_str
}
