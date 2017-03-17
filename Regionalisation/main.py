import pysal
import numpy as np
import random

example_path = "C:/Python27/Lib/site-packages/pysal/examples"

f = pysal.open(example_path + "/us_income/usjoin.csv")
pci = np.array([f.by_col[str(y)] for y in range(1929, 2010)])
pci = pci.transpose()
pci.shape

print(pci.shape)

w = pysal.open(example_path + "/us_income/states48.gal").read()

np.random.seed(100)
random.seed(10)
r = pysal.Maxp(w, pci, floor = 5, floor_variable = np.ones((48, 1)), initial = 99)

print(r.regions)