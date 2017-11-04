package main

import (
    "fmt"
    "os"
)


func main() {
  // READ CONFIG FILE ------------------------------------
  config_file := "config.json"
  config := new(JSON)
  read_json(config_file, config)

  // WRITE CONFIG FILE AS TEST EXAMPLE
  output_file := "test.json"
  output_json_str := write_json(*config, 0)
  file, err := os.Create(output_file)
  check(err)
  defer file.Close()
  _, err = file.WriteString(output_json_str)
  check(err)
  fmt.Printf("Saved to file!\n")
}
