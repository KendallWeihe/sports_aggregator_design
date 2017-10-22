# Machine learning design...

# TODO:
  - get all unique play types
  - rank all data types...
    - i.e. each player is ranked -- based on influence -- and given a unique identifier based on rank
    - ... extend this thought process to the other data types
  - create players document
  - gather other statistics...
    - team stats at time of game
    - ...?


# What the data will look like...
  - columns:
    - team1_*
      - cummulative stats at time of game
        - ...?
      - time
      - team identifier (ranked)
      - player identifier (ranked)
      - data type (numerical identifier based on rank)
      - data value (outcome -- normalized)
      - ...NOTE: various data types will require extended data points
        - i.e. sometimes there are multiple players involved

  - NOTE on normalization...
    - try normalizing data relative to the game stats
    - try normalizing data relative to the season stats
    - ... the "in the moment" vs "the long term"

  - NOTE on ranking...
    - try ranking relative to the game stats
    - try ranking relative to the entire...

    - PLAYERS:
      - try ranking relative to the individual game
      - try ranking relative to the entire league

    - DATA TYPE:
      - there is only one type of ranking for data types

    - RANKING ALGORITHM...
      - ...?
      - regression?
