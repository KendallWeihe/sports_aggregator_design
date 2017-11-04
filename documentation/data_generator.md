# Data generator design:

...write in Go?

# The idea:
  - input config file
  - output ml data

# Config file:
  - input data path
  - output path
  - key to data list
  - keys to data types
    - column number
    - type enum
    - ...
  - normalization


# Technical design:
  - Go program
  - read config file as CLA
  - structures:
    - Row
      - ...

  - psuedocode:
    - read config file
    - initialize variables
    - get input files paths  
    - iterate through files:
      - open file
      - iterate through data list:
        - add to variables


# JSON stuff:
  - find all indices of { and } and [ and ]
  - ORDERING:
    - first "
    - following "
    - colon
      - "
        - ,
      - {
        - " -- RECURSIVE CALL
      - [
        - " -- SEE BEGINNING
        - { -- RECURSIVE CALL
