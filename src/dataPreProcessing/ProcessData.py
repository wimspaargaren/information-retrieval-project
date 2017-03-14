#!/usr/bin/python2.4
#
# Small script to show PostgreSQL and Pyscopg together
#
#86.87.235.82 port 8082
#dbname twitter
#login user: rick pass: proost

import psycopg2
import json
import codecs

try:
    conn = psycopg2.connect("dbname='twitter' user='rick' host='86.87.235.82' password='proost' port='8082'")
except:
    print "I am unable to connect to the database"

cur = conn.cursor()
try:
    cur.execute("""SELECT * from data limit 1""")
except:
    print "I can't select from specified table!"

rows = cur.fetchall()
for row in rows:
    #reader = codecs.getreader("utf-8")
    #obj = json.load(reader(row[1]))
    array = json.dumps(row[1])
    a = json.loads(array)
    print(a["text"])
 