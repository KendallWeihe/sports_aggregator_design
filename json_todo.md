# JSON package documentation:

### Objectives:
  - read from file into JSON structure
  - find based on key
    - specify type
      - value (returns as string)
      - json object (JSON)
      - json list (JSONList)
      - list of json values ([]string)
  - create new JSON structure
    - just use the structure...?
  - write JSON structure to file





# Write pseudocode:
  - recursive functions...

  - write_json():
    - "{"
    - iterate through key values and write
    - iterate through JSON objects
      - recursive call
    - iterate through JSON lists
      - write_json_list()
    - "}"
  - write_json_list():
    - "["
    - iterate through values and write
    - iterate through JSON objects
      - write_json()
    - "]"
