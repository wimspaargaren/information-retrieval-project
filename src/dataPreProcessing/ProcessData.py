#!/usr/bin/python2.4
#
# Small script to show PostgreSQL and Pyscopg together
#

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
        cur.execute("""SELECT * from data""")
    except:
        print "I can't select from specified table!"

    rows = cur.fetchall()

    tweetList = []
    for row in rows:
        globalTweetIdList.append(rows[0])
        tweetList.append(row[2].rstrip("\r\n").lower())
    return tweetList
 

if __name__ == '__main__' :
    tweetList = getTweets()
    print("Initing BM25\n")
    bm25 = bm25.BM25(tweetList, delimiter=' ')
    #Seperate list of words with comma remove spaces!
    #Query = 'voetbal,voetballen,voetbalschoen,scheenbeschermers,scheenbeschermer,arena,soccer'
    Query = 'hardlopen,rennen,running,lopen,run,jogging,marathon,gerend,gelopen,hardgelopen,ran,hard lopen'
    #Query = 'tennis,tennisracket,racket,tennisbaan,tennisbal-ball,tennisball,tennis court'
    Query = Query.split(",")
    for query in Query :
        print query
    print("Calculating scores\n")
    scores = bm25.BM25Score(Query)
    counter = 0
    scoresCounter = 0
    print("Scores:\n")
    for score in scores:
        if score > 0 :
            scoresCounter = scoresCounter + 1
            print(score)
            print(tweetList[counter])
        
        counter = counter + 1

    print("TotalFound: ")
    print(scoresCounter)