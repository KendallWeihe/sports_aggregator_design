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
  team_abbr_key := "gameplaybyplay.game.homeTeam.Abbreviation"

  file_count := 0
  enum := 0
  input_path := "/home/kendall/Development/mfs_data/NBA_play_by_play/2017"
  output_file := "/home/kendall/Development/sports_aggregator_design/data_generator/enumerate/output/teamAbbreviation.json"
  output_json := new(JSON)
  output_json.key_value = make(map[string]string)
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

    team_abbr, _, _ := find(input_data, team_abbr_key)
    fmt.Printf("Found team_abbr_key...\n")
    if _, ok := output_json.key_value[*team_abbr]; !ok {
      output_json.key_value[*team_abbr] = fmt.Sprintf("%d", enum)
      enum += 1
    }

    output_json_str = write_json(*output_json, 0)
    fmt.Printf("%s\n", output_json_str)

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
