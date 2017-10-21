# sports_aggregator_design
This repo houses the design plans for collecting sports statistics

# DevOps thangs...
  - AWS Lambda
    - 3 different Lambda's (see below)
    - code written in Go
  - CloudWatch Event (CRON)
  - S3 to store data
    - maybe a database instead...?

# S3 storage outline:
  - Stats_configs
    - nba_team_config
    - nba_player_config
    - nfl_team_config
    - nfl_player_config
    - ...
  - League_1
    - Date_1
      - Event_1
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
      - Event_2
      - Event_N
    - Date_2
    - Date_N
  - League_2
  - League_N

# Three different Lambda's:
  - todays_events:
    - ran every 30 minutes
    - new `<date>_events.csv` generated at 12:01AM every day
    - finds all events scheduled for today
    - stores in `<date>_events.csv`
  - find_live_events:
    - reads from `<date>_events.csv`
    - gets current time
    - finds all events live _right now_
    - invokes `collect_stats` Lambda for live events
  - collect_stats:
    - collects stats for given event

# todays_events Lambda design:
  - structs:
    - ...
  - input:
    - None -- invoked via a CRON
  - output:
    - PUT `<date>_events.csv` to S3
  - logic flow:
    - check's if it is new day:
      - if so, creates new day file
    - GET HTML
    - parse HTML for leagues & events
    - add events to `<data>_events.csv` file
      - no duplicates
    - PUT the S3

# find_live_events Lambda design:
  - structs:
    - ...
  - input:
    - None -- invoked via a CRON
  - output:
    - None -- invokes next Lambda
  - logic flow:
    - GET `<data>_event.csv` from S3
    - iterate over events
      - if event is live
        - invoke `collect_stats`

# collect_stats Lambda design
  - structs:
    - League:
      - string name
      - array<Event>
    - Event:
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
  - input:
    - JSON containing data relevant to event
    ```python
      {
        "event_link": "<link>",
        "league": "<name>",
        "home_team": "<name>",
        "away_team": "<name>",
        "time_start": "<timestamp>",
        "home_team_players": [
          "player_1_name",
          "player_2_name",
          "..."
        ],
        "away_team_players": [
          "player_1_name",
          "player_2_name",
          "..."
        ]
      }
    ```
  - output:
    - stats files to S3
  - logic flow:
    - construct structs from input
    - GET HTML from input link
    - parse HTML
      - add stats to structs
    - save files to S3

# Questions:
  - can I load a struct based on a JSON config in Go?
    - i.e. ...
      ```python
        "bball_team": {
          "string": "name",
          "string": "timestamp",
          "string": "event_timestamp",
          "string": "points",
          "string": "rebounds"
          "...": "..."
        }
      ```
  - can I run concurrent processes on an AWS Lambda written in Go?
    - i.e. each event would start their own process
