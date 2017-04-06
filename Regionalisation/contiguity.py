import pysal

w = pysal.weights.Contiguity.Rook.from_shapefile("E:/information_retrieval/voronoi.shp", "field_1")