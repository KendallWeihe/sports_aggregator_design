package main

import (
    "fmt"
    "io/ioutil"
    // "reflect"
    "encoding/csv"
    "os"
    // "color"
    "log"
    "strconv"
    "strings"
)

type StatType struct {
  name string
  keys []string
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
  config_file := "config_pbp.json"
  config := new(JSON)
  read_json(config_file, config)

  input_path, _, _ := find(config, "input_path")
  box_score_path, _, _ := find(config, "box_score_path")
  output_path, _, _ := find(config, "output_path")
  plays_key, _, _ := find(config, "plays_key")
  _, f_stat_types, _ := find(config, "stat_types")

  stat_types, num_columns := get_stat_types(f_stat_types)
  fmt.Printf("\nstat_types: %s\n", stat_types)

  // READ DATA FILES ------------------------------------
  input_files, err := ioutil.ReadDir(*input_path)
  check(err)
  for _, file := range input_files { // ITERATE OVER FILES
    // TODO: read all files at once -- possible speed improvement
    fmt.Printf("Generating for file [%s]...\n", file.Name())
    file_path := *input_path + "/" + file.Name()
    input_data := new(JSON)
    read_json(file_path, input_data) // READ THE PLAY JSON OBJECT

    _, _, plays := find(input_data, *plays_key)
    var table [][]string
    for _, play := range plays.json_objs { // ITERATE OVER PLAYS
      quarter := play.key_value["quarter"]
      time := play.key_value["time"]
      // fmt.Printf("Quarter: [%s] Time [%s]\n", quarter, time)
      minutes, err := strconv.ParseFloat(strings.Split(time, ":")[0], 64)
      check(err)
      seconds, err := strconv.ParseFloat(strings.Split(time, ":")[1], 64)
      check(err)
      time_in_seconds := (minutes * 60.0) + seconds

      quarter_nomarlized, err := strconv.ParseFloat(quarter, 64)
      check(err)
      quarter_nomarlized = quarter_nomarlized / 4.0
      time_normalized := time_in_seconds / (12.0 * 60.0)

      row := make([]string, num_columns+2)
      row[0] = fmt.Sprintf("%f", quarter_nomarlized)
      row[1] = fmt.Sprintf("%f", time_normalized)
      // TODO: add a play_type column
      index := 2
      for index < (num_columns+2) {
        row[index] = "0.0"
        index += 1
      }
      index = 2

      for name, _ := range play.json_nested { // ITERATE OVER THE PLAY TYPE ATTRIBUTES
        stat_type_keys := get_stat_type_keys(name, stat_types)
        for _, stat_type_key := range stat_type_keys {
          // ENUMERATE THE STAT VALUE
          stat_value, _, _ := find(play, stat_type_key)
          enumerated := enumerate_play_attr(stat_type_key, *stat_value)
          row[index] = enumerated
          index += 1
        }
      }

      // colorized_output := fmt.Sprintf("Num columns: %d, Len row: %d, Row: [%v]\n", num_columns, len(row) row)
      // color.Cyan(colorized_output)
      table = append(table, row)
      // colorized_output = fmt.Sprintf("Table: [%v]\n\n", table)
      // color.Magenta(colorized_output)
    }

    // read box score result
    box_score_file_path := *box_score_path + "/" + file.Name()
    box_score_data := new(JSON)
    read_json(box_score_file_path, box_score_data)
    home_team_score, _, _ := find(box_score_data, "gameboxscore.quarterSummary.quarterTotals.homeScore")
    away_team_score, _, _ := find(box_score_data, "gameboxscore.quarterSummary.quarterTotals.awayScore")
    home_team_score_f, _ := strconv.ParseFloat(*home_team_score, 64)
    away_team_score_f, _ := strconv.ParseFloat(*away_team_score, 64)
    spread := home_team_score_f - away_team_score_f
    spread_str := fmt.Sprint(spread)
    row := make([]string, num_columns+2)
    row[0] = spread_str
    // fmt.Printf("Spread: %s\n", spread_str)
    index := 1
    for index < (num_columns+2) {
      row[index] = "0.0"
      index += 1
    }
    table = append(table, row)

    // WRITE TO OUTPUT FILE
    output_path := *output_path + "/" + file.Name() + ".csv"
    output_file, err := os.Create(output_path)
    check(err)
    w := csv.NewWriter(output_file)
  	w.WriteAll(table) // calls Flush internally

  	if err := w.Error(); err != nil {
  		log.Fatalln("error writing csv:", err)
  	}
    output_file.Close()
  }
}
