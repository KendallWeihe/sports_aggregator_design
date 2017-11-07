package main

import (
    "fmt"
    "strings"
    "strconv"
    // "io/ioutil"
)

var team_abbreviation_path = "enumerate/output/teamAbbreviation.json"
var team_abbreviation_data = new(JSON)
var err_ta = read_json(team_abbreviation_path, team_abbreviation_data)

var players_path = "enumerate/output/players.json"
var players_data = new(JSON)
var err_p = read_json(players_path, players_data)

var shot_type_path = "enumerate/output/shotType.json"
var shot_type_data = new(JSON)
var err_st = read_json(shot_type_path, shot_type_data)

var turnover_type_path = "enumerate/output/turnoverType.json"
var turnover_type_data = new(JSON)
var err_tt = read_json(turnover_type_path, turnover_type_data)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func teamAbbreviation(team_abbr string) float64 {
  enum := 0.0
  max_enum := float64(len(team_abbreviation_data.key_value))
  for k, v := range team_abbreviation_data.key_value {
    if k == team_abbr {
      v_float, err := strconv.ParseFloat(v, 64)
      check(err)
      enum = v_float
      break
    }
  }

  return enum / max_enum
}

func playerID(player_id string) float64 {
  enum := 0.0
  max_enum := float64(len(players_data.json_nested))
  for k, v := range players_data.json_nested {
    if k == player_id {
      enum_val, _, _ := find(v, "enum")
      v_float, err := strconv.ParseFloat(*enum_val, 64)
      check(err)
      enum = v_float
      break
    }
  }

  return enum / max_enum
}

func shotType(shot_type string) float64 {
  enum := 0.0
  max_enum := float64(len(shot_type_data.key_value))
  for k, v := range shot_type_data.key_value {
    if k == shot_type {
      v_float, err := strconv.ParseFloat(v, 64)
      check(err)
      enum = v_float
      break
    }
  }

  return enum / max_enum
}

func turnoverType(turnover_type string) float64 {
  enum := 0.0
  max_enum := float64(len(turnover_type_data.key_value))
  for k, v := range turnover_type_data.key_value {
    if k == turnover_type {
      v_float, err := strconv.ParseFloat(v, 64)
      check(err)
      enum = v_float
      break
    }
  }

  return enum / max_enum
}

func foulType(foul_type string) float64 {
  return -0.4
}

func true_false(stat_value string) float64 {
  var enumerated float64
  switch stat_value {
    case "false": enumerated = 0.5
    case "true": enumerated = 1.0
    default: enumerated = 0.0
  }
  return enumerated
}

func enumerate_play_attr(stat_type_key string, stat_value string) string {
  // (gdb) p stat_type_key
  // $3 = "fieldGoalAttempt.teamAbbreviation"
  // (gdb) p *stat_value
  // $4 = "NYK"

  // fmt.Printf("stat_type_key: %s, stat_value: %s\n", stat_type_key, stat_value)

  if stat_value == "" {
    // fmt.Printf("stat_value == ''\n")
    return "0.0"
  }

  enumerated := 0.0
  keys := strings.Split(stat_type_key, ".")
  if stringInSlice("teamAbbreviation", keys){
    enumerated = teamAbbreviation(stat_value)
  } else if stringInSlice("ID", keys) {
    enumerated = playerID(stat_value)
  } else if stringInSlice("shotType", keys) {
    enumerated = shotType(stat_value)
  } else if stringInSlice("turnoverType", keys) {
    enumerated = turnoverType(stat_value)
  } else if stringInSlice("foulType", keys) {
    enumerated = foulType(stat_value)
  } else if (stringInSlice("shotLocation", keys) ||
              stringInSlice("foulLocation", keys)) {
    location, err := strconv.ParseFloat(stat_value, 64)
    check(err)
    if stringInSlice("x", keys) {
      enumerated = location / 940
    } else {
      enumerated = location / 500
    }
  } else if stringInSlice("Points", keys) {
    points_float, err := strconv.ParseFloat(stat_value, 64)
    check(err)
    enumerated = points_float / 3.0
  } else if stringInSlice("outcome", keys) {
    switch stat_value {
      case "BLOCKED": enumerated = 0.33
      case "MISSED": enumerated = 0.66
      case "SCORED": enumerated = 1.0
      default: enumerated = 0.0
    }
  } else if stringInSlice("offensiveOrDefensive", keys) {
    switch stat_value {
      case "DEF": enumerated = 0.5
      case "OFF": enumerated = 1.0
      default: enumerated = 0.0
    }
  } else if (stringInSlice("isStolen", keys) ||
              stringInSlice("isPersonal", keys) ||
              stringInSlice("isTechnical", keys) ||
              stringInSlice("isFlagrant1", keys) ||
              stringInSlice("isFlagrant2", keys)) {
    enumerated = true_false(stat_value)
  } else if (stringInSlice("totalAttempts", keys) ||
              stringInSlice("attemptNum", keys)) {
    switch stat_value {
      case "1": enumerated = 0.33
      case "2": enumerated = 0.66
      case "3": enumerated = 1.0
      default: enumerated = 0.0
    }
  } else {
    fmt.Printf("DEBUG: %s, %s\n", stat_type_key, stat_value)
    enumerated = 0.0
  }

  // fmt.Printf("stat_value: %s, enumerated: %f\n", stat_value, enumerated)
  return fmt.Sprintf("%f", enumerated)
}
