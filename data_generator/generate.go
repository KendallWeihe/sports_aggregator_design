package main

import (
    "fmt"
)

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

  config_file := "config.json"
  config := new(JSON)
  read_json(config_file, config)

  key := "gameplaybyplay.plays.play"
  value, list := find(custom_json, key)
  fmt.Printf("%s", value)
  fmt.Printf("%v", list)

  // TODO:
  //   - read config
  //   - define output struct
  //   - list input files
  //   - iterate over files
  //     - read file
  //       - iterate over list
  //         - add to output struct

}
