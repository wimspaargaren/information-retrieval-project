# Export points
# Can be imported in QGIS in order to do Voronoi there

from classes import database
import numpy as np
import time

conn = database.connect()

db = conn.cursor()
db.execute("SELECT * FROM public.strava WHERE category IS NOT NULL ORDER BY id")
points = db.fetchall()

points2 = []

for point in points:
    # factorize category
    running = 0
    cycling = 0
    if point[1] == "running":
        running = 1
    if point[1] == "cycling":
        cycling = 1

    # Check if location already exists
    overlap = [t for t in points2 if t[1] == point[3] and t[2] == point[4]]

    if len(overlap) > 0:
        # Get existing overlapping item
        for idx, item in enumerate(points2):
            if point[3] == item[1] and point[4] == item[2]:
                # Add values to existing item
                item[3] += running
                item[4] += cycling

                # Put item back
                points2[idx] = item
    else:
        # Get ID and position
        reformat = [point[0], point[3], point[4], running, cycling]
        points2.append(reformat)

np.savetxt("E:/information_retrieval/strava" + str(time.time()) + ".csv", points2, delimiter=";", fmt='%s')
