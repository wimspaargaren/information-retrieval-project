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
from numpy import genfromtxt


def connect():
    # Define our connection string
    conn_string = "host='" + s.server + "' port='" + s.port + "' dbname='" + s.database_name + "' user='" + s.user + "' password='" + s.password + "'"

    # get a connection, if a connect cannot be made an exception will be raised here
    conn = psycopg2.connect(conn_string)

    # conn.cursor will return a cursor object, you can use this cursor to perform queries
    print "Connected!\n"
    return conn

# config
write_database = True
write_screen = False
write_map = True
strava = True
plt.hold(True)

# get stuff
conn = connect()
db = conn.cursor()
# db.execute("SELECT * FROM public.data WHERE category IS NOT NULL ORDER BY id")
# tweets = db.fetchall()

if strava:
    tweets = genfromtxt('E:/information_retrieval/strava1492007566.76.csv', delimiter=';')
else:
    tweets = genfromtxt('E:/information_retrieval/tweets1491997871.83.csv', delimiter=';')

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
if strava:
    pass
else:
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

# for tweet in tweets:
#     location = (tweet[5], tweet[6])
#
#     id_list.append(tweet[0])  # ID
#     position_list.append(location)  # LAT, LON
#
#     category_weights = list(tweet[8:])
#     replaced = [0 if v is None else v for v in category_weights]    # Null -> 0
#     categories.append(replaced)

# perform kNN

# kd = pysal.cg.kdtree.KDTree(np.array(position_list))
# w = pysal.weights.Distance.KNN(kd, amount_neighbours)
if strava:
    w = pysal.weights.Contiguity.Rook.from_shapefile("E:/information_retrieval/strava/strava.shp", "field_1")
else:
    w = pysal.weights.Contiguity.Rook.from_shapefile("E:/information_retrieval/v3/shapefile.shp", "field_1")

# z = np.random.random_sample((w.n, 2))

# Get attributes for each item in W, put those in the categories list
if strava:
    for i in w:
        from_csv = [x for x in tweets if x[0] == i[0]]     # get categories for this tweet
        factors = from_csv[0][3:4]
        categories.append(factors)
else:
    for i in w:
        from_csv = [x for x in tweets if x[0] == i[0]]     # get categories for this tweet
        bm25s = from_csv[0][3:12]
        categories.append(bm25s)

# for index, item in enumerate(w):
#     print index
#     print item
#     print categories[index]

z = np.array(categories)
z.transpose()   # prepare for maxp
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
    data = []

    # print region

    for tweet_id in region:
        data_ids.append(tweet_id)
        from_csv = [x for x in tweets if x[0] == tweet_id][0]  # get tweet

        data_lats.append(from_csv[1])
        data_lons.append(from_csv[2])
        data.append(from_csv[3:])


        # Find skipped ones with this ID and also append them
        # for item in skipped:
        #     if item[1] == tweet_id:
        #         data_ids.append(item[0])

    # Determine region category
    if strava:
        cat_sum = [0, 0]
        for i in range(2):
            # aggregate
            for d in data:
                cat_sum[i] = cat_sum[i] + d[i]
    else:
        cat_sum = [0] * 11

        for i in range(11):
            # aggregate
            for d in data:
                cat_sum[i] = cat_sum[i] + d[i]

    highest_index = 999
    highest_value = 0
    for index, item in enumerate(cat_sum):
        if item > highest_value:
            highest_value = item
            highest_index = index

    print highest_index

    if strava:
        categories = ["running", "cycling"]
    else:
        categories = ["running", "gymnastics", "cycling", "bootcamp", "fightingsport",
                      "yoga", "soccer", "fitness", "swimming", "dancing", "hockey"]

    category = categories[highest_index]
    print category
    print highest_value

    if write_database:
        # Insert regions, possible to add properties
        query = "INSERT INTO public.clusters(cluster_set_id, data_ids, category) VALUES (%s, %s, %s);"
        db.execute(query, (cluster_set_id, data_ids, category))

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