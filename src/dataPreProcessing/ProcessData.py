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
import bm25

def getTweets() : 
    print("Retrieving tweets: \n")
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
 
if __name__ == '__main__' :
    getTweets()
    fn_docs = 'mycorpus.txt'
    bm25 = bm25.BM25(fn_docs, delimiter=' ')
    Query = 'The intersection graph of paths in trees survey Graph'
    Query = Query.split()
    scores = bm25.BM25Score(Query)
    tfidf = bm25.TFIDF()
  #  print bm25.Items()
    for i, tfidfscore in enumerate(tfidf):
        print("\n")
        print i, tfidfscore