# Introduction

# Objective

# Data Collection

To achieve this objective, data must be collected for the application. This data will be collected from two sources, which are social media platforms.

The first source is [Twitter](https://twitter.com/). Twitter is a good choice to collect data from since it has a very rich [API](https://dev.twitter.com/rest/public). Another good reason to use Twitter is that tweets contain mostly text about what people are doing and along with that, a tweet can contain a geolocation of the user who posted the tweet. To collect relevant data from Twitter, Twitter's streaming API will be used. The streaming API can be given a bounding box, to stream tweets which are sent from the area of the bounding box. After this every tweet will be checked if it has a geolocation and if so, this tweet will be collected. This is a relevant for the application, since there is a high change to derive from the tweet what a user was doing at an exact location in Amsterdam.

The second source is [Strava](https://www.strava.com/). Strava is a good choice to collect data from, since it is a social media platform for athletes. Since this platform is used to keep track of sport activities like cycling and running, the data contains routes which are represented by geolocations. Just like Twitter, Strava has got a rich [API](https://strava.github.io/api/). To collect data from Strava, the Stava streaming API will be used. This API can also be given a bounding box, such that data from a specific area can be obtained.

With these sources a lot of data which is relevant for the application, can be collected. This data can than be analyzed which will be discussed in the Methods section.

# Methods
(Thom)

# Speficications & Execution Plan

# Expected Outcomes

# Evaluation & Outlook
