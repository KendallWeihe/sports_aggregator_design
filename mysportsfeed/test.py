import pdb
import os
import requests
import time

# # HEADER STUFF
# pwd = os.environ["PWD"]
# authentication = "Authorization: Basic kendallweihe:{}".format(pwd)
#
# compression = "Accept-Encoding: gzip"
#
#
#
# heades = {}
#
# s = requests.Session()
#
# # watch throttling
#     # force=false
# time.sleep(3)

from ohmysportsfeedspy import MySportsFeeds
import json

msf = MySportsFeeds(version="1.0")

msf.authenticate("kendallweihe", os.environ["PWD_MSF"])

game_dates = []
for i in range(20161001, 20170701):
    game_dates.append(str(i))

for date in game_dates:
    schedule = msf.msf_get_data(
                                league="nba",
                                feed="daily_game_schedule",
                                season="2016-2017-regular",
                                format="json",
                                fordate="20161025"
                                )

    for game in schedule:
        pass

# output = msf.msf_get_data(league='nba',season='2016-2017-regular',feed='gameplaybyplay',format='json',game_id="")

print
