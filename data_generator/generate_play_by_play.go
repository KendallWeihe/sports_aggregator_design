package main

import (
    "fmt"
    // "reflect"
)

type Row struct {
  elements []string
}

type DataFile struct {
  columns []string
  rows []Row
}

func get_data_keys(data_types *JSONList) []string {
  var data_keys []string
  for key, json_list := range data_types.json_list {
    fmt.Printf("key[%s]\nvalues[%s]\n", key, json_list)
    num_attributes := len(json_list.values)

    var json_keys []string
    index := 0
    for index < num_attributes {
      json_key := key + "." + json_list.values[index]
      fmt.Printf("%s\n", json_key)
      json_keys = append(json_keys, json_key)
      index += 1
    }

    data_keys = append(data_keys, json_keys)
    fmt.Printf("\n")
  }
  return data_keys
}

func main() {

  config_file := "config.json"
  config := new(JSON)
  read_json(config_file, config)

  input_path, _, _ := find(config, "input_path")
  output_path, _, _ := find(config, "output_path")
  data_list_key, _, _ := find(config, "data_list_key")
  _, data_types, _ := find(config, "data_types")
  fmt.Printf("Input path: %s\n", input_path)
  fmt.Printf("Output path: %s\n", output_path)
  fmt.Printf("Data list key: %s\n", data_list_key)
  fmt.Printf("Data types: %v\n\n", data_types)
  fmt.Printf("data_types.key_value: %v\n", data_types.key_value)
  fmt.Printf("data_types.json_nested: %v\n", data_types.json_nested)
  fmt.Printf("data_types.json_list: %v\n\n", data_types.json_list)

  data_keys := get_data_keys(data_types)
  fmt.Printf("data_keys: %s", data_keys)

  // TODO:
  //   - define DataFile with data_keys

}
