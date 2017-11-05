package main

import (
    "fmt"
    "strings"
    "strconv"
    // "io/ioutil"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func teamAbbreviation(team_abbr string) float64 {
  return 0.5
}

func playerID(player_id string) float64 {
  return 0.5
}

func shotType(shot_type string) float64 {
  return 0.5
}

func turnoverType(turnover_type string) float64 {
  return 0.5
}

func foulType(foul_type string) float64 {
  return 0.5
}

func enumerate_play_attr(stat_type_key string, stat_value string) string {
  // (gdb) p stat_type_key
  // $3 = "fieldGoalAttempt.teamAbbreviation"
  // (gdb) p *stat_value
  // $4 = "NYK"

  fmt.Printf("stat_type_key: %s, stat_value: %s\n", stat_type_key, stat_value)

  if stat_value == "" {
    return "0.1"
  }

  keys := strings.Split(stat_type_key, ".")

  enumerated := 0.0
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
  } else if stringInSlice("shotLocation", keys) {
    if stringInSlice("x", keys) {
      enumerated = 0.0
    } else {
      enumerated = 0.0
    }
  } else if stringInSlice("Points", keys) {
    points_float, err := strconv.ParseFloat(stat_value, 64)
    check(err)
    enumerated = points_float / 3.0
  }

  fmt.Printf("stat_value: %s, enumerated: %f\n", stat_value, enumerated)
  return "0.2"
}
