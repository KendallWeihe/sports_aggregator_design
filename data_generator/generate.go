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

  custom_json := new(JSON)
  read_json("/home/kendall/Development/mfs_data/NBA_play_by_play/2016/20160304-POR-TOR.json", custom_json)

  key := "gameplaybyplay.plays.play"
  value, list := find(custom_json, key)
  fmt.Printf("%s", value)
  fmt.Printf("%v", list)

  //   - build a config file
  //   - iterate through plays:
  //     - if data type exists: add to output file struct
  //     - save file

}
