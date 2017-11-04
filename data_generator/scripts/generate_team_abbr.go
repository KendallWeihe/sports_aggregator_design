package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
  unique_abbrs := make(map[string]int)
  count := 0
  input_path := "/home/kendall/Development/mfs_data/NBA_play_by_play/2017"

  // READ DATA FILES ------------------------------------
  input_files, err := ioutil.ReadDir(input_path)
  check(err)
  for _, file := range input_files { // ITERATE OVER FILES
    file_path := input_path + "/" + file.Name()
    input_data := new(JSON)
    read_json(file_path, input_data) // READ THE PLAY JSON OBJECT

    home_team_abbr_key := "gameplaybyplay.game.homeTeam.Abbreviation"
    away_team_abbr_key := "gameplaybyplay.game.awayTeam.Abbreviation"

    home_team_abbr, _, _ := find(input_data, home_team_abbr_key)
    away_team_abbr, _, _ := find(input_data, away_team_abbr_key)

    if _, ok := unique_abbrs[*home_team_abbr]; !ok {
      fmt.Printf("Unique team: [%s] ID: [%d]\n", *home_team_abbr, count)
      unique_abbrs[*home_team_abbr] = count
      count += 1
    }

    if _, ok := unique_abbrs[*away_team_abbr]; !ok {
      fmt.Printf("Unique team: [%s] ID: [%d]\n", *away_team_abbr, count)
      unique_abbrs[*away_team_abbr] = count
      count += 1
    }

    if count == 30 {
      // TODO: save map to file & exit

    }

    // fmt.Printf("Unique list: [%v]\n", unique_abbrs)
  }
}
