#Tasks
A five-pages Project idea document. In this small document, you will put in a written
form your pitch presentation. In details, you are required to formalise the project topic and
to motivate, in a written form:

<s>– the relevance to the course; - Wim</s>

<s>– the innovative and challenging nature of the project; - Ruben</s>

– the requirements/specifications of the application to develop (e.g. using MOSCOW
requirement prioritisation); - Rick

<s>– a coarse-grained execution plan, with responsibilities clearly assigned to each group
member ; - Wim</s>

<s>– a coarse-grained business plan;</s>

<s>– Flow diagram of major system components, and mockups of any user-facing interfaces
(crowdworkers and end-users) - Thom</s>

– your planned evaluation and success metrics. - Ruben  
– Further expand on project idea with more detailed explanation of uses for different stakeholders(See company story in pitch folder) - Daan

The project idea documents must be delivered by Week 4.

NOTE 1: the requirement specifications must be complete. The design and implementation
may address only a subset of the specifications, for time/effort reasons. Identify some system
functions (5/6) – a core set – keeping in mind that functions should support user needs.

NOTE 2: while the complete specification should be ready by Week 4, you are strongly
invited to start development as soon as possible
The Project idea document will be assessed according to the Rubric reported in Table 2

# Amsterdam in motion

## Introduction

A lot of data is gathered on the web and with that data you can do amazing things. Many approaches to analyze data can be taken for all kind of things. Conventionally, data was gathered for specific reasons and only used for those specific reasons. Nowadays a lot of social data is gathered and with out-of-the-box thinking, this data can be used in an enormous amount of ways. With API-crawlers [ref-for-definition-needed], web-scrapers[ref-for-definition-needed], you can retrieve this data and analyze it in any way possible to for example visualize trends, activities and life-style differences in any neighbourhood.

## Project idea
The goal of 'Amsterdam in motion' is to visualize different neighbourhoods in amsterdam, based on the type of activity going on there, geographical as well as in time. This then can be used to know where certain needs are, like events and shops. This can also lead to more social interaction between people with the same activity, because they will know where that activity is most popular.

##The relevance to the course

The relevance of our project to the course is as follows. Information retrieval deals with the representation, storage, organization of, and access to information items. The final product which will be implemented during this course will cover all of these topics.

First of all the representation of data. To give the users of "Amsterdam in motion" a user-friendly presentation of data, a map will be used to display areas where people carry out certain sport activities.

Secondly the storage and organization of the data is covered. For storage of data the product uses two crawlers which gather information from Twitter and Strava and store this in a database.  This data is unstructured, but certain documents will satisfy the information need. To find the relevant documents from the information which is gathered, the data will be organized in a second database per sport category. This will be done by using the BM25 algorithm to find relevant tweets, for example, for a certain sports category.

And last the access to information is covered, since social media data is accessed on the background, users of the product "Amsterdam in motion" can access organized data displayed on a map.

In short the product "Amsterdam in motion" organizes the overload of social media data and categorizes this into different sport categories. To achieve this goal all topics of the course information retrieval are covered, so it is pretty clear that the project is very relevant.


## The innovative and challenging nature of the project
Tracking individuals on various social media platforms is no challenge anymore.
With the multitude of platforms available every kind of activity and personal
updates can be tracked.

What still lacks online is an option to easily bundle
all this individual data into accurate and up to date information on when and
where these activities happen. Because of that our project aims to innovate in
the way the data will be modeled onto the map of Amsterdam. The data will be
analysed and categorised based on several factors.

First of all the data will of course be mapped into different neighborhoods
based on the type of sport activity. Next our project splits the data based on
the time the activity takes place. What we are curious to see is on what types
of activity happen at which time of the day and on which day of the week.
Furthermore our plan for this project is to analyse the data based on age and
gender and map wether or not this influences the outcome of information on the
map. This aspect of the project is quite challenging as no age information is
directly bundled in the crawled data and in twitter also no gender information
is supplied. In order to still gather data with age and gender we want to
analyse profile images with external libraries such as Face++ and genderize.

## The requirements/specifications of the application

### MOSCOW

#### Must haves

- Data crawler on Twitter
- Data crawler on Strava
- BM25 analytics with known taxonomies
- Visual representation of categorized data on Mapbox

#### Should haves

- Dyamic selection on day of week to visualize data
- Dynamic selection on time of day to visualize data
- Neighborhood selection to represent several attributes regarding the neighborhood

#### Could haves

- Real-time input of new taxonomies
- Use of face++ to retrieve additional attributes
- Use of Genderize to retrieve additional attributes
- Display multiple maps next to each other for easy comparison of different filters

#### Won't haves

## Evalutation and Succes metrics
The way we will evaluate the project is going to be based on the MOSCOW
definitions. This is a useful and structured listing of the project in which
the importance of each feature is represented. The must haves in this list
are the bare minimum that should be working at the end of the project. The
should haves and the could haves after that are the metrics that are used to
further evaluate the success of the project. How well does a feature work and
what is the experienced usefulness of a feature.

Besides the metrics defined by MOSCOW the project will be evaluated by how well
we have reached our goals set in the beginning of the project. These goals can
be evaluated by analysing how well the different stakeholders can use the end
product for their respective goals.

## Overall architecture and UI mockups
 /Flow diagram of major system components, and mockups of any user-facing interfaces (crowdworkers and end-users) - Thom/

![Amsterdam in motion - pipeline](pipeline.png "The information pipeline")

In order to be able to provide a nice overview of distinct neighbourhoods in Amsterdam several steps need to be taken. A data source providing information about activities (bounded in space and time) is needed. Twitter and Strava are being used in this project, but similar sources can be used as well. A crawler gathers every tweet from Amsterdam and stores this in a PostgreSQL database. 

These tweets can be used for both the processing and identification of neighbourhoods and direct display on the user interface as illustration.

The tweets are processed to determine which topics are discussed. Relevant (i.e. mathing with a sport related word list) tweets are used to classify neighbourhoods. This is done with the CNN- and/or kNN-algorithms. The clusters are also used in the user interface.

This user interface is being used to display the neighbourhoods, some characteristics of the neighbourhoods and some tweets for illustration. This interface is displayed using Mapbox. The different stakeholders might get different features (which also enables for a price differentiation), but the basics are the same. The most extensive version of the interface is shown here. Variants for different stakeholders will not be distinguished currently.

![User interface](mockup.png "The end-user interface")

This mockup shows how end-users will interact with the end-result. The interface provides a big graphical interface containing a map with tweets and areas (neighbourhoods) projected on it. A settings panel allows for a little customization in what is shown. Export functionality is provided for stakeholders who want to use the data for their own application (e.g. to calculate where a shop can be opened).



## Execution plan
In the table below our execution plan for the upcoming weeks is described. 

In week four the idea document will be elaborated. The different tasks for the report will be divided among the group members.

In week five the data collection and management of the system will be implemented. The implementation of the collection and management of Twitter, will be done by Wim and Rick and the implementation of the collection and management of Strave, will be done by Daan and Ruben.

In week six the text processing and classification will be implemented. This will be done by Wim, Ruben and Thom. This means that the data collection will be ready to be displayed on the map.

In week seven the user interface will be implemented by Daan and Thom. This means that the first sport activity groups will be displayed on the map.

In week eight the product will be finished. Everyone has the responsibility that the different parts of the system work together.

|  |  Week 4 	|  Week 5 	| Week 6  	| Week 7  	| Week 8  	| Week 9	|
| --- | ---		| ---		| ---		| ---		| ---		| ---		|
| What | Idea document|Data collection & Management|Text processing & Classification|User interface|Finishing product|Presentation|
| By whom | Everyone  |Twitter: Rick & Wim Strava: Daan & Ruben|Wim & Ruben & Thom|Daan & Thom |Everyone|Everyone   |