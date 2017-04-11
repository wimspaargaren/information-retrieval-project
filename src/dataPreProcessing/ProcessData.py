#!/usr/bin/python2.4
#
# Small script to show PostgreSQL and Pyscopg together
#

import psycopg2
import json
import codecs
import bm25
import sys

class Query():
    def __init__(self,category, searchquery):
        # Sets all the properties
        self.category = category
        self.searchquery = searchquery

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
    
class sportTweet():
    def __init__(self,tweetID,category,bm_running,bm_gymnastics,bm_cycling,bm_bootcamp,bm_fightingsport,bm_yoga,bm_soccer,bm_fitness,bm_swimming,bm_dancing,bm_hockey):
        self.category = category
        self.bm_running = bm_running
        self.bm_gymnastics = bm_gymnastics
        self.bm_cycling = bm_cycling
        self.bm_bootcamp = bm_bootcamp
        self.bm_fightingsport = bm_fightingsport
        self.bm_yoga = bm_yoga
        self.bm_soccer = bm_soccer
        self.bm_fitness = bm_fitness
        self.bm_swimming = bm_swimming
        self.bm_dancing = bm_dancing
        self.tweetID = tweetID
        self.bm_hockey = bm_hockey

def getSportTweets() : 
    print("Retrieving sport tweets: \n")
    connString = "dbname='"+sys.argv[1]+"' user='"+sys.argv[2]+"' host='"+sys.argv[3]+"' password='"+sys.argv[4]+"' port='"+sys.argv[5]+"'"
    try:
        conn = psycopg2.connect(connString)
    except:
        print "I am unable to connect to the database"

    cur = conn.cursor()
    try:
        cur.execute("""SELECT id, category, bm_running,bm_gymnastics, bm_cycling, bm_bootcamp, bm_fightingsport, bm_yoga,bm_soccer, bm_fitness, bm_swimming, bm_dancing,bm_hockey FROM data WHERE bm_running IS NOT NULL OR bm_gymnastics IS NOT NULL OR bm_cycling IS NOT NULL OR bm_bootcamp IS NOT NULL OR bm_fightingsport IS NOT NULL OR bm_yoga IS NOT NULL OR bm_soccer IS NOT NULL OR bm_fitness IS NOT NULL OR bm_swimming IS NOT NULL OR bm_dancing IS NOT NULL OR bm_hockey IS NOT NULL;""")
    except:
        print "I can't select from specified table!"

    rows = cur.fetchall()

    tweetList = []
    for row in rows:
        tweetList.append(sportTweet(row[0],row[1],row[2],row[3],row[4],row[5],row[6],row[7],row[8],row[9],row[10],row[11],row[12]))

    return tweetList

def resetCatgories() : 
    print("TODO RESET CATEGORIES")


def updateBmColumns(updateQuerys):
    connString = "dbname='"+sys.argv[1]+"' user='"+sys.argv[2]+"' host='"+sys.argv[3]+"' password='"+sys.argv[4]+"' port='"+sys.argv[5]+"'"
    try:
        conn = psycopg2.connect(connString)
    except:
        print "I am unable to connect to the database"

    cur = conn.cursor()
    for updateQuery in updateQuerys:
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
    queryList = []
    queryList.append(Query("fitness",'fitness,steps,spinning,aerobics,squat,deadlift,bench press,military press, pull up,gym,sportschool,home,thuis,indoorroeier,crosstrainer,hometrainer,calf raises,chin-up,opdrukken,dumbell flexion extension,lunge,bankdrukken,grapevine,zumba,bodystep'))
    queryList.append(Query("soccer",'voetbal,football,soccer,speeldveld,voetbaleveld,voetbalterrein,grasveld,bal,ball,voetbalschoenen,voetbalshirt,voetbalbroek,voetbalschoenen,scheenbeschermers,scheidsrechter,referee,doel,doelpunt,doel,blessuretijd,strafschoppen'))
    queryList.append(Query("running",'hardlopen,rennen,running,lopen,run,jogging,marathon,gerend,gelopen,hardgelopen,ran,hard lopen,duursport,hardloopschoenen,hardloopkleding'))
    queryList.append(Query("swimming",'swimming,zwemmen,water,pools,zwembad,zwemmend,gezwommen,vrije slag,vlinder slag,rugslag,schoolslag,borstcrawl,rugcrawl,wisselslag,wedstrijdzwemmen,openwaterzwemmen,schoonspringen,synchroonzwemmen,vinzwemmen,waterbasketbal,waterpolo,winterzwemmen,zwemmend redden,elementair zwemmen'))
    queryList.append(Query("fightingsport",'vechtsport,fighting sport,vechtensporten,boksen,kung fu,worstelen,schermen,karate,MMA,kickboksen,zelfverdediging,judo'))
    queryList.append(Query("cycling",'wielrennen,wielersport,fietsen,racefiets,wielrenfiets,cycling,fiets,bicycle,baanwielrennen,bmx,mountainbiken,wegwielrennen,wielrennersshirt,wielrennersbroek,wielrennersschoenen,helm'))
    queryList.append(Query("dancing", 'Academisch ballet,Academische dans,Acrobatische rock,Aerobics,American Rhythm,American Smooth,Argentijnse tango,Bachata,Balboa,Balfolk,Ballet,Ballroom,Ballroomtango,Basse danse,Batucada,Biodanza,Bolero,Boogiewoogie,Bostella,Bourree,Breakdance,Buikdans,Cachucha,Cakewalk,Calypso,Cancan,Capoeira,Chachacha,Charleston,Cheerleading,Chunga,Coco,Conga,Congada,Cumbia,Computerdans,Dancehall,Dansexpressie,Danza,Discofox,Dodendans,Dubstep,Engelse wals,Euritmie,Fandango,Frevo,Freestyle,Flamenco,Foxtrot,Gaillarde,Gavotte,Gigue,Hakken,Headbangen,Hiphop,Hora,Hula,Jazzballet,Jive,Jumpstyle,Karakterdans,Kathakali,Klassiek ballet,Klompendans,Krumping,Lambada,Landler,Limbo,Lindy hop,Linedance,Locking,Macarena,Mambo,Mazurka,Melbourne Shuffle,Menuet,Merengue,Moonwalk,Moshen,Moderne dans,Moresca,Paaldansen,Paso doble,Pogo,Polka,Polonaise,Popping,Quadrille,Quickstep,Ragga,Reggaeton,Reidans,Repetitieve dans,Riverdance,Rondedans,Rueda de Casino,Rumba,Salsa,Saltarello,Samba,Sarabande,Showdance,Sirtaki,Skanken,Slowfox,Son,Stijldansen,Streetdance,Striptease,Swing,Swingfox,Tango,Tapdansen,Tecktonik,Twirling,Twist,Volksdans,Vogue,Vogeltjesdans,Wals,Weense wals,Ballet,Concert dance,Contact improvisation,Contemporary dance,Ecstatic dance,Jazz dance,Lyrical dance,Postmodern dance,Performance dance,Tap dance,American Rhythm,Bolero,East Coast Swing,Mambo,Rumba,Cha Cha,Cumbia,Argentine tango,Jive,Merengue,Milonga,Reggaeton,Rueda,Salsa,Samba,ballroom dance,Samba,Brazilian dance,Samba de Gafieira,Zumba,dance,dansen,dans,Zouk,Quebradita,gayunder,Balboa,Blues dance,Carolina Shag,Charleston,Collegiate Shag,East Coast Swing,Hand Jive,Jitterbug,Jive,Leroc,Lindy Hop,Modern Jive,St. Louis Shag,West Coast Swing,Traditional jazz,Foxtrot,Tango,Viennese Waltz,Waltz,Quickstep,Bharatanatyam,Chhau dance,Gaudiya Nritya,Kathak,Kathakali,Kuchipudi,Manipuri dance,Mohiniyattam,Odissi,Sattriya,Crowdsurfing,Moshing,Skanking,Jump up,Footwork,Powermoves,Toprock,Bounce,Clown Walk,Electric boogaloo,Flexing,Floating,Krumping,Litefeet,Robot dance,hip hop,Memphis Jookin,Animation,Hit Dem Folks,Liquiding,Moonwalk,Scarecrow,Strobing,Waving,Snap dance,Turfing,Disco dance,Electro Dance,Free step,Grinding,Hard dance,Hardcore dancing,House dance,Hustle,Jacking,Jumpstyle,Melbourne Shuffle,Slothing,Vogue,Waacking,Hard Rock Pogo,Metal Mosh,Punk Rock Pogo,Ska Dance,Baroque dance,English country dance,Masque,Medieval dance,Regency dance,Vintage dance,Liturgical dance,Acro dance,Belly dance,Bollywood dance,Calypso,Fire dance,Flamenco,Flying Men Dance,Hip Hop,Kizomba,Line dance,Native American,Novelty and fad dances,Reggae,Salsaton,Semba,Soca,hooping,Zumba,The Routledge Dance Studies Reader,The dance; an historical survey of dancing in Europe,The Body, Dance and Cultural Theory,Modern Tap Dance,Jazz,Rhythm,Body and Soul'))
    queryList.append(Query("gymnastics",'Gymnastics,athletics,exercises,aerobics,aerobic,acrobatics,bars,rings,turnen,leotard,turnace,turnace hall,gym,gymen'))
    queryList.append(Query("hockey",'hockey,hockey stick,hockey veld,Jockstrap,protective cup,cup,puck,unicycle hockey'))
    queryList.append(Query("yoga",'Yoga,Indische mystiek,meditatie,mediteren,meditation,meditate,concentration,concentreren,bewustzijn,mindfulness,Ascetische leer,Bewegingsleer,Bewustzijnsoefening,dojo,flexibility,stamina,stress,anti-stress,nutrition,Fitstar,daily yoga,down dog'))
    queryList.append(Query("bootcamp",'Bootcamp,crossfit,insanity,tabata,circuittraining,move,pull,push,core,intensief,intensive,grenzen verleggen,Core Power,FitDeck,FitMoves,exercises'))
    for q in queryList :
        print(q.category)
        Query = q.searchquery.split(",")
        column = "bm_" + q.category
        updateQueryList = []
        scores = bm25.BM25Score(Query)
        counter = 0
        scoresCounter = 0
        for score in scores:
            if score != 0 :
                if q.category == "gymnastics"  :
                    print("UPDATE data SET " + column + " = "+ str(score) +" WHERE id = "+ str(globalTweetIdList[counter]))
                
                if q.category == "fitness" :  
                    print("UPDATE data SET " + column + " = "+ str(score) +" WHERE id = "+ str(globalTweetIdList[counter]))

                updateQueryList.append("UPDATE data SET " + column + " = "+ str(score) +" WHERE id = "+ str(globalTweetIdList[counter]))
                scoresCounter = scoresCounter + 1
            counter = counter + 1
        
        print("Updating bm columns")
        updateBmColumns(updateQueryList) 

    resetCatgories()
    updateSportCategories = []
    sporTweets = getSportTweets()   
    for t in sporTweets:
        maxval = 0.0
        category = "undefined"
        if t.bm_running > maxval :
            maxval = t.bm_running
            category = "running"
        if t.bm_gymnastics > maxval :
            maxval = t.bm_gymnastics
            category = "gymnastics"
        if t.bm_cycling > maxval :
            maxval = t.bm_cycling
            category = "cycling"
        if t.bm_bootcamp > maxval :
            maxval = t.bm_bootcamp
            category = "bootcamp"
        if t.bm_fightingsport > maxval :
            maxval = t.bm_fightingsport
            category = "fightingsport"
        if t.bm_yoga > maxval :
            maxval = t.bm_yoga
            category = "yoga"
        if t.bm_soccer > maxval :
            maxval = t.bm_soccer
            category = "soccer"
        if t.bm_fitness > maxval :
            maxval = t.bm_fitness
            category = "fitness"
        if t.bm_swimming > maxval :
            maxval = t.bm_swimming
            category = "swimming"
        if t.bm_dancing > maxval :
            maxval = t.bm_dancing
            category = "dancing"
        if t.bm_hockey > maxval :
            maxval = t.bm_hockey
            category = "hockey"
        updateSportCategories.append("UPDATE data SET category = '"+ category +"' WHERE id = " + str(t.tweetID))
        
    updateBmColumns(updateSportCategories)
