#Tasks
A five-pages Project idea document. In this small document, you will put in a written
form your pitch presentation. In details, you are required to formalise the project topic and
to motivate, in a written form:

– the relevance to the course; - Wim

– the innovative and challenging nature of the project; - Ruben

– the requirements/specifications of the application to develop (e.g. using MOSCOW
requirement prioritisation); - Rick

– a coarse-grained execution plan, with responsibilities clearly assigned to each group
member ;

– a coarse-grained business plan;

– Flow diagram of major system components, and mockups of any user-facing interfaces
(crowdworkers and end-users)

– your planned evaluation and success metrics.

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

#The relevance to the course

#The innovative and challenging of the project

##The requirements/specifications of the application

### MOSCOW

#### Must haves

- Data crawler on Twitter
- Data crawler on Strava
- BM25 analytics with known taxonomies
- Visual representation of categorized data on Mapbox

#### Should haves

- Dyamic selection on day of week to visualize data
- Dynamic selection on time of day to visualize data
- Neighbourhood selection to represent several attributes regarding the neighbourhood

#### Could haves

- Real-time input of new taxonomies
- Use of face++ to retrieve additional attributes
- Use of Genderize to retrieve additional attributes
- Display multiple maps next to eachother for easy comparison of different filters

#### Won't haves