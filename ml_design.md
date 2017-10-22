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
      - data types... (see below)
        - fill in columns for particular type
        - 0 out all other columns/types
        - example...
          ```csv
            previous_stats || 1:33 || 0,0,0... || 1,0.75,0.3,0,0,1... || 0,0,0...
          ```
  - the idea...
    - write a program that can dynamically generate datasets based on configuration files
    - example config.json...
      ```python
        {
          "include_cummulative_stats": true,
          "play_types": [
            "fieldGoalAttempt",
            "rebound",
            "..."
          ],
          "number_of_plays": 100
        }
      ```

  - NOTE on normalization...
    - try normalizing data relative to the game stats
    - try normalizing data relative to the season stats
    - ... the "in the moment" vs "the long term"

  - NOTE on ranking...
    - PLAYERS:
      - try ranking relative to the individual game
      - try ranking relative to the entire league
    - TEAMS:
      - ...?
    - RANKING ALGORITHM...
      - ...?
      - regression?

# data types:
  - fieldGoalAttempt
    - team
    - player
    - points
    - outcome
    - shotType
    - assist? (boolean)
    - blocked? (boolean)
    - location???
  - rebound
    - team
    - player
    - Off/Def? (boolean)
  - turnover
    - team
    - stolenByPlayer
    - lostByPlayer
    - turnoverType
    - isStolen
  - foul
    - team
    - drawnByPlayer
    - penalizedPlayer
    - foulType
    - isPersonal?
    - isTechnical?
    - isFlagrant1?
    - isFlagrant2?
    - location???
  - freeThrowAttempt
    - team
    - player
    - totalAttempts
    - attemptNum
    - outcome
  - substitution???
    - team
    - outgoingPlayer
    - incomingPlayer
  - jumpBall
