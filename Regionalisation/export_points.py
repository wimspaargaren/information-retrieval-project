# Export points
# Can be imported in QGIS in order to do Voronoi there

from classes import database
import numpy as np
import time

conn = database.connect()

db = conn.cursor()
db.execute("SELECT * FROM public.data WHERE category IS NOT NULL ORDER BY id")
tweets = db.fetchall()

tweets2 = []

for tweet in tweets:
    replaced = [0 if v is None else v for v in tweet[8:]]

    # Check if location already exists
    overlap = [t for t in tweets2 if t[1] == tweet[5] and t[2] == tweet[6]]

    if len(overlap) > 0:
        # Get existing overlapping item
        for idx, item in enumerate(tweets2):
            if tweet[5] == item[1] and tweet[6] == item[2]:
                # Add values to existing item
                item[3:] = [sum(x) for x in zip(item[3:], replaced)]

                # Put item back
                tweets2[idx] = item
    else:
        # Get ID and position
        reformat = [tweet[0], tweet[5], tweet[6]]
        # Add values
        reformat.extend(replaced)
        tweets2.append(reformat)

np.savetxt("E:/information_retrieval/tweets" + str(time.time()) + ".csv", tweets2, delimiter=";", fmt='%s')
