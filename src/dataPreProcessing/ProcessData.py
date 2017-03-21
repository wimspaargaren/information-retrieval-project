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
        globalTweetIdList.append(row[0])
        tweetList.append(row[2].rstrip("\r\n").lower())
    return tweetList
 

def updateTweets(category, idUpdateSTring) :
    connString = "dbname='"+sys.argv[1]+"' user='"+sys.argv[2]+"' host='"+sys.argv[3]+"' password='"+sys.argv[4]+"' port='"+sys.argv[5]+"'"
    updateQuery = "UPDATE data SET category = '"+ category +"' WHERE id IN ("+idUpdateSTring+")"
    print(updateQuery)
    try:
        conn = psycopg2.connect(connString)
    except:
        print "I am unable to connect to the database"

    cur = conn.cursor()
    try:
        cur.execute(updateQuery)
    except:
        print "I can't update for some reason!"

    conn.commit()
    cur.close()

if __name__ == '__main__' :
    tweetList = getTweets()
    print("Initing BM25\n")
    bm25 = bm25.BM25(tweetList, delimiter=' ')
    #Seperate list of words with comma remove spaces!
    
    #category = "fitness"
    #Query = 'fitness,steps,spinning,aerobics,squat,deadlift,bench press,military press, pull up,gym,sportschool,home,thuis,indoorroeier,crosstrainer,hometrainer,calf raises,chin-up,opdrukken,dumbell flexion extension,lunge,bankdrukken,grapevine,zumba,bodystep'
    
    #category = "soccer"
    #Query = 'voetbal,football,soccer,speeldveld,voetbaleveld,voetbalterrein,grasveld,bal,ball,voetbalschoenen,voetbalshirt,voetbalbroek,voetbalschoenen,scheenbeschermers,scheidsrechter,referee,doel,doelpunt,doel,blessuretijd,strafschoppen'

    #category = "running"
    #Query = 'hardlopen,rennen,running,lopen,run,jogging,marathon,gerend,gelopen,hardgelopen,ran,hard lopen,duursport,hardloopschoenen,hardloopkleding'
  
    #category = "swimming"
    #Query = 'swimming,zwemmen,water,pools,zwembad,zwemmend,gezwommen,vrije slag,vlinder slag,rugslag,schoolslag,borstcrawl,rugcrawl,wisselslag,wedstrijdzwemmen,openwaterzwemmen,schoonspringen,synchroonzwemmen,vinzwemmen,waterbasketbal,waterpolo,winterzwemmen,zwemmend redden,elementair zwemmen'

    #category = "fightingsport"
    #Query = 'vechtsport,fighting sport,vechtensporten,boksen,kung fu,worstelen,schermen,karate,MMA,kickboksen,zelfverdediging,judo'

    #category = "cycling"
    #Query = 'wielrennen,wielersport,fietsen,racefiets,wielrenfiets,cycling,fiets,bicycle,baanwielrennen,bmx,mountainbiken,wegwielrennen,wielrennersshirt,wielrennersbroek,wielrennersschoenen,helm'

    #dancing query must be updated.

    #category = "gymnastics"
    #Query = 'Gymnastics,athletics,exercises,aerobics,aerobic,acrobatics,bars,rings,turnen,leotard,turnace,turnace hall,gym,gymen'

   # category = "hockey"
   # Query = 'hockey,hockey stick,hockey veld,Jockstrap,protective cup,cup,puck,unicycle hockey'

    #category = "yoga"
    #Query = 'Yoga,Indische mystiek,meditatie,mediteren,meditation,meditate,concentration,concentreren,bewustzijn,mindfulness,Ascetische leer,Bewegingsleer,Bewustzijnsoefening,dojo,flexibility,stamina,stress,anti-stress,nutrition,Fitstar,daily yoga,down dog'

    category = "bootcamp"
    Query = 'Bootcamp,crossfit,insanity,tabata,circuittraining,move,pull,push,core,intensief,intensive,grenzen verleggen,Core Power,FitDeck,FitMoves,exercises'

    Query = Query.split(",")
    for query in Query :
        print query
    print("Calculating scores\n")
    scores = bm25.BM25Score(Query)
    counter = 0
    scoresCounter = 0
    print("Scores:\n")
    idsToUpdate = ""
    for score in scores:
        if score > 0 :
            scoresCounter = scoresCounter + 1
            print(score)
            idsToUpdate = idsToUpdate + str(globalTweetIdList[counter]) + ", "
         #   print(tweetList[counter])
        
        counter = counter + 1

    updateTweets(category,idsToUpdate[:-2])
    print("TotalFound: ")
    print(scoresCounter)