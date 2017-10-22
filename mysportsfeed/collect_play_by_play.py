import pdb
import os
import requests
import time

from ohmysportsfeedspy import MySportsFeeds
import json
import sys

msf = MySportsFeeds(version="1.0")
msf.authenticate("kendallweihe", os.environ["PWD_MSF"])

def get_dates():
    game_dates = []
    year = "2014"
    for i in range(11, 13):
        for j in range(1, 32):
            date = "{}{}{}".format(year, i, str(j).zfill(2))
            game_dates.append(date)
    year = "2015"
    for i in range(1, 8):
        for j in range(1, 32):
            date = "{}{}{}".format(year, str(i).zfill(2), str(j).zfill(2))
            game_dates.append(date)

    return game_dates

def get_daily_schedule(date):
    time.sleep(5)
    schedule = msf.msf_get_data(
                                league="nba",
                                feed="daily_game_schedule",
                                season="2014-2015-regular",
                                format="json",
                                fordate=date,
                                force=True,
                                store_file=None
                                )
    return schedule

def get_play_by_play(game_id):
    time.sleep(5)
    play_by_play = msf.msf_get_data(
                                    league="nba",
                                    feed="game_playbyplay",
                                    season="2014-2015-regular",
                                    format="json",
                                    gameid=game_id,
                                    force=True,
                                    store_file=None
                                    )
    return play_by_play

def main():
    requests_count = 0
    game_dates = get_dates()
    for date in game_dates:

        found_error = False
        while not found_error:
            try:
                schedule = get_daily_schedule(date)
                break
            except Exception as e:
                error = str(e)
                print(error)
                if "429" in error:
                    # sleep for a minute and try again
                    print("Error getting schedule for {}...\nSleeping for 60 seconds and trying again".format(date))
                    time.sleep(60)
                elif "400" in error:
                    print("400 Error...")
                    found_error = True
                else:
                    pdb.set_trace()
                    sys.exit()

        if found_error:
            continue

        if (not schedule.get("dailygameschedule")) or (not schedule["dailygameschedule"].get("gameentry")):
            print("No games found for {}".format(date))
            continue

        for game in schedule["dailygameschedule"]["gameentry"]:
            game_id = "{}-{}-{}".format(date, game["awayTeam"]["Abbreviation"], game["homeTeam"]["Abbreviation"])

            found_error = False
            while not found_error:
                try:
                    play_by_play = get_play_by_play(game_id)
                    break
                except Exception as e:
                    error = str(e)
                    print(error)
                    if "429" in error:
                        # sleep for a minute and try again
                        print("Error getting game for {}...\nSleeping for 60 seconds and trying again".format(game_id))
                        time.sleep(60)
                    elif "400" in error:
                        print("400 Error...")
                        found_error = True
                    else:
                        pdb.set_trace()
                        sys.exit()

            if found_error:
                continue

            f = open("{}/{}.json".format(os.environ["OUT_DIR"], game_id), "w")
            f.write(json.dumps(play_by_play, indent=4))
            f.close()

            print("Collected {}".format(game_id))
            os.system("rm results/*")

if __name__ == "__main__":
    main()
