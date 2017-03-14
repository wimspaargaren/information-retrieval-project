# Meeting 14-3
- Document looks quite good. Reflects what we have talked about.
- How is the progress going? How is the data collection going?
- Tweets need to be preprocessed -> Filter to see what is in there.
- Do we already have the taxonomies? Start on BM25 with taxonomies.
- What attributes are we going to use later on in the analysis. Define which
ones.

- Attributes:
  - geolocation
  - type of sports activity -> topics that are sports -> do this with a
  threshold.

- Clustering:
  - There is no single recepy.
  - Different techniques might cater better to what we want to do.
  - Do not cluster them in administrative units. Neighborhoods cross those
  boundaries.
  - Going upwards from individual data is fine, other way not so much. Do not
  disaggregate the data.
  - Because the data is spatially referenced there is a pitfall to skew the
  data based on that.
  - KDE: Kernel Density Estimation -> Make structure in the spatially distributed
  data. Invalidate that the data is randomly distributed.
  - Spatial dependence, data that is close is more related.

- Get acquainted with pycel as soon as possible.
- Lecture about spatial clustering will tell how to find out if there is a
local or a global autocorrelation. Local/Global formula's.
  - Contiguity Matrix
- See wether there is correlation between the data collected from twitter and
strava. KDE can help with that -> Multi-correlation.
- Switch to Python in general -> One single language
- Create additional columns in the data for the days/time.
  - Separate time frame into hours. Perhaps some ranges from morning,
  midday, afternoon, evening. This is an aggregation upwards.
- Validate the data by comparing different techniques.
