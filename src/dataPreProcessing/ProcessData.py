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
import sys

globalTweetIdList = []

def getTweets() : 
    print("Retrieving tweets: \n")
    connString = "dbname='"+sys.argv[1]+"' user='"+sys.argv[2]+"' host='"+sys.argv[3]+"' password='"+sys.argv[4]+"' port='"+sys.argv[5]+"'"
    try:
        conn = psycopg2.connect(connString)
    except:
        print "I am unable to connect to the database"

    cur = conn.cursor()
    try:
        cur.execute("""SELECT * from data limit 5000""")
    except:
        print "I can't select from specified table!"

    rows = cur.fetchall()

    tweetList = []
    for row in rows:
        array = json.dumps(row[1])
        a = json.loads(array)
      #  print("Tweet:")
        globalTweetIdList.append(rows[0])
      #  print(a["text"].encode("utf-8").rstrip("\r\n"))
        tweetList.append(a["text"].encode("utf-8").rstrip("\r\n").lower())
    return tweetList
 

if __name__ == '__main__' :
   
    tweetList = getTweets()
    print("Initing BM25\n")
    bm25 = bm25.BM25(tweetList, delimiter=' ')
    #Query = 'voetbal voetballen voetbalschoen scheenbeschermers scheenbeschermer arena soccer'
    Query = 'hardlopen rennen running lopen run jogging marathon gerend gelopen hardgelopen ran'
    Query = Query.split()
    print("Calculating scores\n")
    scores = bm25.BM25Score(Query)
    counter = 0
    print("Scores:\n")
    for score in scores:
        if score > 0 :
            print(score)
            print(tweetList[counter])
        
        counter = counter + 1
 #   tfidf = bm25.TFIDF()
  #  print bm25.Items()
 #   for i, tfidfscore in enumerate(tfidf):
 #       print("\n")
 #       print i, tfidfscore