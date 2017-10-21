import pdb
import os
import requests
import time

from ohmysportsfeedspy import MySportsFeeds
import json
import datetime

msf = MySportsFeeds(version="1.0")

msf.authenticate("kendallweihe", os.environ["PWD_MSF"])

game_dates = []
for i in range(20161025, 20170701):
    game_dates.append(str(i))

requests_count = 0
time_start = datetime.datetime.now()
pdb.set_trace()
for date in game_dates:

    schedule = msf.msf_get_data(
                                league="nba",
                                feed="daily_game_schedule",
                                season="2016-2017-regular",
                                format="json",
                                fordate=date,
                                force=True
                                )
    requests_count += 1

    if (not schedule.get("dailygameschedule")) or (not schedule["dailygameschedule"].get("gameentry")):
        continue

    for game in schedule["dailygameschedule"]["gameentry"]:
        game_id = "{}-{}-{}".format(date, game["awayTeam"]["Abbreviation"], game["homeTeam"]["Abbreviation"])
        play_by_play = msf.msf_get_data(
                                        league="nba",
                                        feed="game_playbyplay",
                                        season="2016-2017-regular",
                                        format="json",
                                        gameid=game_id,
                                        force=True
                                        )
                                        
        f = open("/Users/kendallweihe/Development/msf_data/NBA_play_by_play/{}.json".format(game_id), "w")
        f.write(json.dumps(play_by_play, indent=4))
        f.close()

        os.system("rm results/*")

        requests_count += 1
        if requests_count > 250:
            time_end = datetime.datetime.now()
            seconds_passed = (time_start - time_end).total_seconds()
            wait_time = (60*5 - seconds_passed)
            time.sleep(wait_time)

            requests_count = 0
            time_start = datetime.datetime.now()
