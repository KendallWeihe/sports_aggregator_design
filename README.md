# sports_aggregator_design
This repo houses the design plans for collecting sports statistics

# DevOps thangs...
  - AWS Lambda
    - code written in Go
  - CloudWatch Event (CRON)
  - S3 to store data
    - maybe a database instead...?

# S3 storage outline:
  - League_1
    - Date_1
      - Game_1
        - team_stats
          - time_stamp_1.csv
          - time_stamp_2.csv
          - time_stamp_N.csv
        - player_1_stats
          - time_stamp_1.csv
          - time_stamp_2.csv
          - time_stamp_N.csv          
        - player_2_stats
        - player_N_stats
      - Game_2
      - Game_N
    - Date_2
    - Date_N
  - League_2
  - League_N

# Go Lambda design
  - structs:
    - League:
      - string name
      - array<Game>
    - Game:
      - Team home_team
      - Team away_team
      - time start_time
    - Team:
      - string name
      - array <Player>
    - Player:
      - string name
      - array <Stats> [cummulative over the season]
      - array <Stats> [at current "live" time stamp]
    - Stats:
      - "template" based
      - load data members from config files?

  - ...
  - Logic flow:
    - read todays_games.csv file
    - iterate through games
      - if game is live:
        - get HTML
        - parse HTML (& any JS actions???)
          - construct stats csv's
            - team stats
            - player stats
        - save stats to S3

# Questions:
  - can I load a struct based on a JSON config in Go?
    - i.e. ...
      ```python
        "bball_team": {
          "string": "name",
          "string": "timestamp",
          "string": "game_timestamp",
          "string": "points",
          "string": "rebounds"
          "...": "..."
        }
      ```
  - can I run concurrent processes on an AWS Lambda written in Go?
    - i.e. each game would start their own process
