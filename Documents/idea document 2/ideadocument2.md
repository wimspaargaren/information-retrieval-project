# Introduction

# Objective

# Data Collection

To achieve this objective, data must be collected for the application. This data will be collected from two sources, which are social media platforms. 

The first source is [Twitter](https://twitter.com/). Twitter is a good choice to collect data from, since tweets contain mostly text about what people are doing. Along with that, a tweet can contain a geolocation of the user, who posted the tweet. Another good reason to use Twitter is that it has a very rich [API](https://dev.twitter.com/rest/public). To collect relevant data from Twitter, Twitter's streaming API will be used. The streaming API can be given a bounding box, to stream tweets which are send from the area of the bounding box. After this, every tweet will be checked if it has a geolocation and if so, this tweet will be collected. This is relevant for the application, since it might be possible to derive from the tweet what a user was doing at an exact location in Amsterdam.

The second source is [Strava](https://www.strava.com/). Strava is a good choice to collect data from, since it is a social media platform for athletes. This platform is used to keep track of sport activities like cycling and running, for this reason, data from Strava contains routes which are represented by geolocations. Just like Twitter, Strava has got a rich [API](https://strava.github.io/api/), from which the streaming API will be used. This API can also be given a bounding box, such that data from a specific area can be obtained, which in this case is Amsterdam.

With these two sources data which is relevant for the application, can be collected. Data from strava is relevant, because it contains data about athletes accompanied by their locations. Data from Twitter is relevant, since athletes might use Twitter to tell about their sport activities and if that tweet has a geolocation it can be derived where the athlete was conducting its sport. 
To derive this information after the collection, data must be analyzed, which will be discussed in the Methods section.

# Methods
(Thom)

# Specifications & Execution Plan
In the previous sections the objective is presented and the way the data is collected and how this is processed in order to use for the application. In this section the different features as what can be done with this data in the application are described.

## MoSCoW
Introduction MoSCoW..

### Must haves

- Data crawler on Twitter
- Data crawler on Strava
- BM25 analytics with known taxonomies
- Visual representation of categorized data on Mapbox

### Should haves

- Dynamic selection on day of week to visualize data
- Dynamic selection on time of day to visualize data
- Neighborhood selection to represent several attributes regarding the neighborhood

### Could haves

- Real-time input of new taxonomies
- Use of Face++ to retrieve additional attributes (note: related to whom?)
- Use of Genderize to retrieve additional attributes (note: related to whom?)
- Display multiple maps next to each other for easy comparison of different filters

### Won't haves
- Intensity (heat) map

Explanation different features

Timeline, very small explanation of timeline (comment of achilleas, is this really needed?)
|  |  Week 4 	|  Week 5 	| Week 6  	| Week 7  	| Week 8  	| Week 9	|
| --- | ---		| ---		| ---		| ---		| ---		| ---		|
| What | Idea document|Data collection & Management|Text processing & Classification|User interface|Finishing product|Presentation|
| By whom | Everyone  |Twitter: Rick & Wim Strava: Daan & Ruben|Wim & Ruben & Thom|Daan & Thom |Everyone|Everyone   |


# Expected Outcomes
As mentioned earlier in the section objective, the goal of this application is to identify and characterise neighborhoods in Amsterdam. In figure mockup-ui a mockup of the user interface (UI) is shown, it will become visible how the UI could look.
<p align="center">
  <img alt="Workbench electron-browser component" src="../images/mockup.png" width = "70%">
  <br><i>Figure mockup-ui: Mockup of the User Interface</i>
</p>
The user is presented with a large map of Amsterdam. On this map blue dots are shown which are tweets or strava activities that are found to be relevant by the algorithm. Relevant posts are posts that fall into a sports activity based on the taxonomy of that sport. Around those posts the neighborhood boxes are drawn in different colors to quickly make clear how the neighborhoods of Amsterdam are divided based on sport activities. Because only relevant tweets are shown this can give an indication to the user how active an area is with regards to sports.

In the setting panel the user can customize what will be shown. This will contain a filter on the type of activity and a filter on when a tweet or activity has taken place. The setting to filter when activities were carried out will give insight in activities that only take place at a certain day of the week or even a certain time of the day. There are also show options which toggle areas or labels.


# Evaluation & Outlook
