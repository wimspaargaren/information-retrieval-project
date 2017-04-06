# Meeting 28-3

## Clustering
- Clusters should be spatially continuous. They should have boundaries.
- If you do not identify good districts. The neighborhoods would be clusters 
with a color scheme. Clustering is more than category, it is also the amount of 
tweets in a region. 
- Clusters should be close to on another. Some points are outliers and give 
weird results.
- A cluster would be based on similar cluster attributes. A cluster classifies 
tweets in different categories.
- Clustering points will not give different results than categorising the 
tweets in different sports.
- Spatial autocorrelation should be applied onto the entire data set. 
Neighborhoods should be defined out of that and then with the amount of tweets 
from a region should be used to identify the category of that neighborhood.
- Do the autocorrelation on multiple variables. Otherwise you are correlating 
a value with itself in another location. Add column for a variable and weight 
them that make sense.
- Some activities are related to others and weight them. We could get clusters 
of related activities. Identify neighborhoods with similar activities.
- Rate a tweet based on the BM25 algorithm for all the categories. And use 
these as columns in the Max XP algorithm. (Only 2 or 3 columns)
