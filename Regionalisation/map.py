import matplotlib.pyplot as plt
import mplleaflet
import psycopg2
import settings as s

ids = [637, 19422, 19941, 33553, 22398]
lat = []
lon = []

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

for tweet_id in ids:
    db.execute("SELECT * FROM public.data WHERE id = %s", (tweet_id,))
    tweet = db.fetchone()

    lat.append(tweet[5])
    lon.append(tweet[6])

print lat
print lon

plt.hold(True)
plt.plot(lon, lat, 'b')  # Draw blue line

mplleaflet.show()
