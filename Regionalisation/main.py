import pysal
import numpy as np
from sklearn.neighbors import NearestNeighbors
import random
import psycopg2
import sys
import settings as s
import time


def connect():
    # Define our connection string
    conn_string = "host='" + s.server + "' port='" + s.port + "' dbname='" + s.database_name + "' user='" + s.user + "' password='" + s.password + "'"

    # get a connection, if a connect cannot be made an exception will be raised here
    conn = psycopg2.connect(conn_string)

    # conn.cursor will return a cursor object, you can use this cursor to perform queries
    print "Connected!\n"
    return conn


# get stuff
conn = connect()
db = conn.cursor()
db.execute("SELECT * FROM public.data WHERE category IS NOT NULL ORDER BY id")

# construct lists from db
id_list = []
position_list = []
skipped = []

for tweet in db:
    location = (tweet[5], tweet[6])
    # There is an issue with identical coordinates, we filter those out now and keep track of them
    # We can add them to the clusters later
    if location not in position_list:
        id_list.append(tweet[0])  # ID
        position_list.append(location)  # LAT, LON
    else:
        identical_index = position_list.index(location)
        skipped.append((tweet[0], id_list[identical_index]))  # Tup(ID, ID of identical)

print position_list

# perform kNN
for amount_neighbours in [3]:
    kd = pysal.cg.kdtree.KDTree(np.array(position_list))
    w = pysal.weights.Distance.KNN(kd, amount_neighbours)

    print "Neighbours of first tweet are:"
    print w.neighbors[0]

    z = np.random.random_sample((w.n, 2))
    p = np.ones((w.n, 1), float)
    floor = 3
    solution = pysal.region.Maxp(w, z, floor, floor_variable=p, initial=100)

    # Set this high (9999) in final product, low in debug mode
    solution.inference(nperm=9999)
    print solution.pvalue

    print solution.regions

    cluster_config_id = -1
    query = "INSERT INTO public.cluster_sets(cluster_config_id, timestamp) VALUES (%s, %s) RETURNING id;"
    db.execute(query, (cluster_config_id, time.time()))
    # cluster_set_id = -1
    cluster_set_id = db.fetchone()[0]

    print skipped

    for region in solution.regions:
        data_ids = []
        for index in region:
            tweet_id = id_list[index]
            data_ids.append(tweet_id)

            # Find skipped ones with this ID and also append them
            for item in skipped:
                if item[1] == tweet_id:
                    data_ids.append(item[0])

        # Insert regions, possible to add properties
        query = "INSERT INTO public.clusters(cluster_set_id, data_ids) VALUES (%s, %s);"
        db.execute(query, (cluster_set_id, data_ids))

conn.commit()
db.close()
conn.close()
