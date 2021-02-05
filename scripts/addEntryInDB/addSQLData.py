import pymysql
import pickle

con = pymysql.connect(
  host="localhost",
  user="root",
  passwd="",
  db="users"
)

#Unpickling the values to be added to the db
filename = 'data'
infile = open(filename,'rb')
val = pickle.load(infile)
infile.close()

try:
    with con.cursor() as cur:
        cur.executemany('INSERT INTO Profile (Name, Nickname, Password) VALUES(%s, %s, %s)', val) 
        con.commit()
        print("Data added")
finally:
    con.close()