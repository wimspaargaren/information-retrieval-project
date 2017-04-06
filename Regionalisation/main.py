import matplotlib.pyplot as plt
import pysal
import numpy as np
from sklearn.neighbors import NearestNeighbors
import random
import psycopg2
import sys
import settings as s
import time
import mplleaflet


def connect():
    # Define our connection string
    conn_string = "host='" + s.server + "' port='" + s.port + "' dbname='" + s.database_name + "' user='" + s.user + "' password='" + s.password + "'"

    # get a connection, if a connect cannot be made an exception will be raised here
    conn = psycopg2.connect(conn_string)

    # conn.cursor will return a cursor object, you can use this cursor to perform queries
    print "Connected!\n"
    return conn

# config
write_database = False
write_screen = True
write_map = True
plt.hold(True)

# get stuff
conn = connect()
db = conn.cursor()
db.execute("SELECT * FROM public.data WHERE category IS NOT NULL ORDER BY id")
tweets = db.fetchall()

# Create a set
cluster_config_id = -1
cluster_set_id = -1
timestamp = time.time()

if write_database:
    query = "INSERT INTO public.cluster_sets(cluster_config_id, timestamp) VALUES (%s, %s) RETURNING id;"
    db.execute(query, (cluster_config_id, timestamp))
    cluster_set_id = db.fetchone()[0]

if write_screen:
    print "Cluster " + str(cluster_set_id) + " created at " + str(time.time())

# Filter categories
categories = []
for tweet in tweets:
    if tweet[7] not in categories:
        categories.append(tweet[7])
print categories

# construct lists
id_list = []
position_list = []
categories = []
skipped = []


for tweet in tweets:
    location = (tweet[5], tweet[6])

    # There is an issue with identical coordinates, we filter those out now and keep track of them
    # We can add them to the clusters later
    if location not in position_list:
        id_list.append(tweet[0])  # ID
        position_list.append(location)  # LAT, LON

        category_weights = list(tweet[8:])
        replaced = [0 if v is None else v for v in category_weights]
        categories.append(replaced)

    else:
        identical_index = position_list.index(location)
        skipped.append((tweet[0], id_list[identical_index]))  # Tup(ID, ID of identical)


# perform kNN
for amount_neighbours in [5]:
    # kd = pysal.cg.kdtree.KDTree(np.array(position_list))
    # w = pysal.weights.Distance.KNN(kd, amount_neighbours)
    w = pysal.weights.Contiguity.Rook.from_shapefile("E:/information_retrieval/voronoi.shp", "field_1")

    z = np.random.random_sample((w.n, 2))

    # z = np.array(categories)
    # z.transpose()
    # print z

    p = np.ones((w.n, 1), float)
    floor = 3
    solution = pysal.region.Maxp(w, z, floor, floor_variable=p, initial=100)

    # Set this high (9999) in final product, low in debug mode
    solution.inference(nperm=9999)
    print solution.pvalue

    print str(len(solution.regions)) + " clusters identified "

    print solution.regions

    for region in solution.regions:
        data_ids = []
        data_lats = []
        data_lons = []

        for index in region:
            tweet_id = id_list[index]
            data_ids.append(tweet_id)
            data_lats.append(position_list[index][0])
            data_lons.append(position_list[index][1])

            # Find skipped ones with this ID and also append them
            for item in skipped:
                if item[1] == tweet_id:
                    data_ids.append(item[0])

        if write_database:
            # Insert regions, possible to add properties
            query = "INSERT INTO public.clusters(cluster_set_id, data_ids) VALUES (%s, %s);"
            db.execute(query, (cluster_set_id, data_ids))

        if write_screen:
            print "ID's: " + str(data_ids)

        if write_map:
            plt.plot(data_lons, data_lats, 'b')  # Draw blue line

if write_database:
    conn.commit()

db.close()
conn.close()

print str(time.time() - timestamp) + " elapsed."

if write_map:
    mplleaflet.show()