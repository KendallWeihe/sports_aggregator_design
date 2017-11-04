package main

import (
    "io/ioutil"
    "fmt"
    "os"
)

// INPUTS:
//   - keys to find
//     - if within list, separate ...
//   - otuput file name
// OUTPUT:
//   - file w/ enum values
func main() {
  key1 := "dailyplayerstats.playerstatsentry"
  key2 := "player.ID"
  output_file := "./players.txt"
  unique_values := make(map[string]int)
  count := 0
  input_path := "/home/kendall/Development/mfs_data/NBA_daily_player/2017"

  // READ DATA FILES ------------------------------------
  input_files, err := ioutil.ReadDir(input_path)
  check(err)
  for _, file := range input_files { // ITERATE OVER FILES
    file_path := input_path + "/" + file.Name()
    fmt.Printf("Reading file: %s ...\n\n", file_path)
    input_data := new(JSON)
    read_json(file_path, input_data) // READ THE PLAY JSON OBJECT
    fmt.Printf("JSON file loaded...")

    _, _, playerstatsentry := find(input_data, key1)
    fmt.Printf("Found key1...")
    for _, player := range playerstatsentry.json_objs {
      value, _, _ := find(player, key2)
      if _, ok := unique_values[*value]; !ok {
          fmt.Printf("Unique value: [%s] ID: [%d]\n", *value, count)
          unique_values[*value] = count
          count += 1
      }
    }

    // value, _, _ := find(input_data, key)
    // if _, ok := unique_values[*value]; !ok { // CHECK UNIQUENESS
    //   fmt.Printf("Unique team: [%s] ID: [%d]\n", *value, count)
    //   unique_values[*value] = count
    //   count += 1
    // }

    file, err := os.Create(output_file)
    check(err)
    defer file.Close()
    for k, v := range unique_values {
      line := fmt.Sprintf("%s,%d\n", k, v)
      _, err := file.WriteString(line)
      check(err)
    }

  }

  file, err := os.Create(output_file)
  check(err)
  defer file.Close()
  for k, v := range unique_values {
    line := fmt.Sprintf("%s,%d\n", k, v)
    _, err := file.WriteString(line)
    check(err)
  }
}


// json_obj := input_data
// index := 0
// for index < len(keys) {
//   key := keys[index]
//
//   index += 1
// }
