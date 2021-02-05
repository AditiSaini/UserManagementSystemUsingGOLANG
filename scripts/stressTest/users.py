import pickle

#Unpickling the values to be added to the db
filename = '../addEntryInDB/data'
infile = open(filename,'rb')
val = pickle.load(infile)
infile.close()

#Setting the users info value to val
users_info = val
