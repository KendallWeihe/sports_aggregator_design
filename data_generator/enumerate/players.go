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
  starting_key := "dailyplayerstats.playerstatsentry"
  id_key := "player.ID"
  first_name_key := "player.FirstName"
  last_name_key := "player.LastName"
  position_key := "player.Position"

  file_count := 0
  enum := 0
  input_path := "/home/kendall/Development/mfs_data/NBA_daily_player/2017"
  output_file := "/home/kendall/Development/sports_aggregator_design/data_generator/enumerate/output/players.json"
  output_json := new(JSON)
  output_json.key_value = make(map[string]string)
  output_json.json_nested = make(map[string]*JSON)
  output_json.json_list = make(map[string]*JSONList)
  output_json_str := ""

  // READ DATA FILES ------------------------------------
  input_files, err := ioutil.ReadDir(input_path)
  check(err)
  for _, file := range input_files { // ITERATE OVER FILES
    file_path := input_path + "/" + file.Name()
    fmt.Printf("Reading file: %s ...\n", file_path)
    input_data := new(JSON)
    read_json(file_path, input_data) // READ THE PLAY JSON OBJECT
    fmt.Printf("JSON file loaded...\n")

    _, _, playerstatsentry := find(input_data, starting_key)
    fmt.Printf("Found starting_key...\n")
    for _, player := range playerstatsentry.json_objs {
      id, _, _ := find(player, id_key)
      if _, ok := output_json.json_nested[*id]; !ok {
        first_name, _, _ := find(player, first_name_key)
        last_name, _, _ := find(player, last_name_key)
        position, _, _ := find(player, position_key)
        fmt.Printf("New player found: %s, %s\n", *first_name, *last_name)

        output_json.json_nested[*id] = new(JSON)
        output_json.json_nested[*id].key_value = make(map[string]string)
        output_json.json_nested[*id].json_nested = make(map[string]*JSON)
        output_json.json_nested[*id].json_list = make(map[string]*JSONList)
        output_json.json_nested[*id].key_value["first_name"] = *first_name
        output_json.json_nested[*id].key_value["last_name"] = *last_name
        output_json.json_nested[*id].key_value["position"] = *position
        output_json.json_nested[*id].key_value["enum"] = fmt.Sprintf("%d", enum)
        enum += 1
      }
    }

    output_json_str = write_json(*output_json, 0)
    // fmt.Printf("%s\n", output_json_str)

    if file_count % 5 == 0 {
      file, err := os.Create(output_file)
      check(err)
      defer file.Close()
      _, err = file.WriteString(output_json_str)
      check(err)
      fmt.Printf("Saved to file!\n")
    }
    file_count += 1
  }
}
