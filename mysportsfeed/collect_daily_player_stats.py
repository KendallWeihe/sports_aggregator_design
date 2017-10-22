import pdb
import os
import time

from ohmysportsfeedspy import MySportsFeeds
import json
import sys

msf = MySportsFeeds(version="1.0")
msf.authenticate("kendallweihe", os.environ["PWD_MSF"])

def get_dates():
    game_dates = []
    year = "2015"
    for i in range(10, 13):
        for j in range(1, 32):
            date = "{}{}{}".format(year, i, str(j).zfill(2))
            game_dates.append(date)
    year = "2016"
    for i in range(1, 8):
        for j in range(1, 32):
            date = "{}{}{}".format(year, str(i).zfill(2), str(j).zfill(2))
            game_dates.append(date)

    return game_dates

def get_daily_player(date):
    time.sleep(5)
    stats = msf.msf_get_data(
                                league="nba",
                                feed="daily_player_stats",
                                season="2015-2016-regular",
                                format="json",
                                fordate=date,
                                force=True,
                                store_file=None
                                )
    return stats

def main():
    dates = get_dates()
    for date in dates:

        found_error = False
        while not found_error:
            try:
                stats = get_daily_player(date)
                break
            except Exception as e:
                error = str(e)
                print(error)
                if "429" in error:
                    # sleep for a minute and try again
                    print("429 Error for {}...\nSleeping for 60 seconds".format(date))
                    time.sleep(60)
                elif "Max retries" in error:
                    print("Max retries error for {}...\nSleeping for 60 seconds".format(date))
                    time.sleep(180)
                elif "400" in error:
                    print("400 Error...")
                    found_error = True
                else:
                    pdb.set_trace()
                    sys.exit()

        if found_error:
            continue

        if (not stats.get("dailyplayerstats")) or (not stats["dailyplayerstats"].get("playerstatsentry")):
            print("No stats found for {}".format(date))
            continue

        print("FOUND stats for {}... writing to file".format(date))
        f = open("{}/{}.json".format(os.environ["OUT_DIR"], date), "w")
        f.write(json.dumps(stats, indent=4))
        f.close()

if __name__ == "__main__":
    main()
