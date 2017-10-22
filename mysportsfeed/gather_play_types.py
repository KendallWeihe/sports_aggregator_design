import json
import pdb

f = open("play_by_play_ex.json", "r")
play_by_play = json.loads(f.read())
f.close()

plays = play_by_play["gameplaybyplay"]["plays"]["play"]

unique_play_types = []
for play in plays:
    for child in play:
        print(child)
        if child not in unique_play_types:
            unique_play_types.append(child)

f = open("play_types.txt", "w")
f.write("\n".join(unique_play_types))
f.close()
