package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

/*
  structures:
    - Config: holds configuration data
    - InputData: dynamically generated based on JSON fields
    - Row: holds data for each row
*/

type Config struct {
  date, era string
  time int
  alive bool
}

type InputData struct {

}

type Row struct {

}

/*
  functions:
    - check
    - read_json
    - generate_row
    - write_output_file
*/

func check(e error) {
  if e != nil {
    panic(e)
  }
}

/*
  INPUTS:
    - path to JSON file
  OUTPUTS:
    - json encoded object (could be Config or it could be Row)
*/
func read_json(path string) string {
  raw, err := ioutil.ReadFile("./config.json")
  check(err)

  m := map[string]string{}
  err = json.Unmarshal([]byte(raw), &m)
  check(err)
  fmt.Println(m)

  return ""

  // dec := json.NewDecoder(strings.NewReader(raw))
  // for {
	// 	var m Message
	// 	if err := dec.Decode(&m); err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%s: %s\n", m.Name, m.Text)
	// }
}

/*
  INPUTS:
    - json object
  OUTPUTS:
    - Row object
*/
// func generate_row() Row {
//   var r []Row
//   //...
//   return r
// }

/*
  INPUTS:
    - output path
  OUTPUT:
    - boolean on success
*/
func save_file() bool {
  return true
}

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

  json_data := read_json("config.json")
  fmt.Print(string(json_data))

}
