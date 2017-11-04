package main

import (
    "fmt"
    "io/ioutil"
    // "reflect"
)

type Row struct {
  elements []string
}

type DataFile struct {
  columns []string
  rows []Row
}

type StatType struct {
  name string
  keys []string
}

func enumerate_play_attr(play_type string, play_value string) string {
  return ""
  // TODO:
  //   - basically a switch statement that invokes a functions in a separate program file
  //   - ...
  //   - if "teamAbbreviation":
  //   - else if "Player.ID":
  //   - else if "Location.":
  //   - else use full play_type in default function
}

func get_stat_type_keys(name string, stat_types []StatType) []string {
  for _, stat_type := range stat_types {
    if stat_type.name == name {
      return stat_type.keys
    }
  }
  var empty_str []string
  return empty_str
}

func get_stat_types(f_stat_types *JSON) ([]StatType, int) {
  var stat_types []StatType
  num_columns := 0

  for key, json_list := range f_stat_types.json_list {
    fmt.Printf("key[%s]\nvalues[%s]\n", key, json_list)

    var stat_type StatType
    stat_type.name = key

    num_attributes := len(json_list.values)
    index := 0
    for index < num_attributes {
      json_key := key + "." + json_list.values[index]
      stat_type.keys = append(stat_type.keys, json_key)
      index += 1
    }

    stat_types = append(stat_types, stat_type)

    if num_attributes > num_columns {
      num_columns = num_attributes
    }
  }

  return stat_types, num_columns
}

func main() {

  // READ CONFIG FILE ------------------------------------
  config_file := "config.json"
  config := new(JSON)
  read_json(config_file, config)

  input_path, _, _ := find(config, "input_path")
  // output_path, _, _ := find(config, "output_path")
  plays_key, _, _ := find(config, "plays_key")
  _, f_stat_types, _ := find(config, "stat_types")

  stat_types, num_columns := get_stat_types(f_stat_types)
  fmt.Printf("\nstat_types: %s\n", stat_types)

  // READ DATA FILES ------------------------------------
  input_files, err := ioutil.ReadDir(*input_path)
  check(err)
  for _, file := range input_files { // ITERATE OVER FILES
    file_path := *input_path + "/" + file.Name()
    input_data := new(JSON)
    read_json(file_path, input_data) // READ THE PLAY JSON OBJECT

    _, _, plays := find(input_data, *plays_key)
    for _, play := range plays.json_objs { // ITERATE OVER PLAYS
      quarter := play.key_value["quarter"]
      time := play.key_value["time"]
      fmt.Printf("Quarter: [%s] Time [%s]\n", quarter, time)

      index := 0
      row := make([]string, num_columns)
      row[index] = quarter
      row[index] = time

      for name, _ := range play.json_nested { // ITERATE OVER THE PLAY TYPE ATTRIBUTES
        stat_type_keys := get_stat_type_keys(name, stat_types)
        for _, stat_type_key := range stat_type_keys {
          play_value, _, _ := find(play, stat_type_key)
          // TODO: enumerate the play_value
          row[index] = *play_value
          index += 1
        }
      }

      fmt.Printf("Row: [%v]\n", row)
    }
  }

}
