# Can be imported in QGIS

from classes import database
import numpy as np

conn = database.connect()

db = conn.cursor()
db.execute("SELECT * FROM public.data WHERE category IS NOT NULL ORDER BY id")
tweets = db.fetchall()

tweets2 = []

for tweet in tweets:
    tweets2.append((tweet[0], tweet[5], tweet[6], tweet[8], tweet[9], tweet[10], tweet[11],
                    tweet[12], tweet[13], tweet[14], tweet[15], tweet[16]))

np.savetxt("E:/information_retrieval/tweets.csv", tweets2, delimiter=";", fmt='%s')