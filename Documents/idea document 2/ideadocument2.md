# Introduction

# Objective

# Data Collection

To achieve this objective, data must be collected for the application. This data will be collected from two sources, which are social media platforms. 

The first source is [Twitter](https://twitter.com/). Twitter is a good choice to collect data from, since tweets contain mostly text about what people are doing. Along with that, a tweet can contain a geolocation of the user which posted the tweet. Another good reason to use Twitter is that it has a very rich [API](https://dev.twitter.com/rest/public). To collect relevant data from Twitter, Twitter's streaming API will be used. The streaming API can be given a bounding box, to stream tweets which are send from the area of the bounding box. After this, every tweet will be checked if it has a geolocation and if so, this tweet will be collected. This is relevant for the application, since it might be possible to derive from the tweet what a user was doing at an exact location in Amsterdam.

The second source is [Strava](https://www.strava.com/). Strava is a good choice to collect data from, since it is a social media platform for athletes. This platform is used to keep track of sport activities like cycling and running, for this reason, data from Strava contains routes which are represented by geolocations. Just like Twitter, Strava has got a rich [API](https://strava.github.io/api/). To collect data from Strava, the Stava streaming API will be used. This API can also be given a bounding box, such that data from a specific area can be obtained, which in this case is Amsterdam.

With these two sources data which is relevant for the application, can be collected. This data is relevant, since Strava contains data about athletes accompanied by locations. Data from Twitter is relevant, since athletes might use Twitter to tell about their sport activities and if that tweet has a geolocation it can be derived where the athlete was conducting its sport. 
After the collectoin, data can be analyzed which will be discussed in the Methods section.

# Methods
(Thom)

# Speficications & Execution Plan

# Expected Outcomes

# Evaluation & Outlook
