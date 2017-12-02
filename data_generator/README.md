# Data generator

TODO...

### TODO:
  - recollect for foulType
    - re-enumerate
    - regenerate dataset -- only print debug statements
  - add a play type enum
  - ranking system


## Open source thangs:
  - README
    -
      ```
        I didn't like [this](https://golang.org/pkg/encoding/json/)...
        ...I just wanted to Read, Create, Manipulate, and Save JSON data...
        ...but I did like [this](https://docs.python.org/3/library/json.html)...
        ...so I wrote something like [that](https://docs.python.org/3/) for [this](https://golang.org/doc/)
      ````
    -
      ```
        The (current) idea is minimalist abstraction, for the purposes of easy-to-use (and fast) JSON input/output/manipulation for Go. In other words, you can do these things...

          - (Read || Create) a JSON file
          - Search for JSON: values || objects || lists
          - Manipulate JSON: values || objects || lists
          - Save a JSON file
      ```
  - files:
    - input
      - For Reading & Creating JSON structs
    - output
      - For saving to file
    - manipulate
      - For Editing JSON structs

  - ...

  - TODO:
    - refactor
    - add features (to support README claims)
    - optimization? (Go routines?)
    - write documentation
    - get build verifications
    - open source
    - promote :)
