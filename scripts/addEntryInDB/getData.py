import numpy as np
import pandas as pd
from faker.providers.person.en import Provider
import bcrypt
import pickle

size = 10000000

password = '245'
hashAndSalt = bcrypt.hashpw(password.encode(), bcrypt.gensalt()).decode("utf-8") 

def random_names(name_type, size):
    """
    Generate n-length ndarray of person names.
    name_type: a string, either first_names or last_names
    """
    names = getattr(Provider, name_type)
    return np.random.choice(names, size=size)

df = pd.DataFrame(columns=['Name', 'Nickname', 'Password'])
df['Name'] = random_names('first_names', size)
df['Nickname'] = random_names('last_names', size)
data = np.random.randint(1, 30001, size=size)
df['Num'] = pd.DataFrame(data, columns=['random_numbers'])
df['Nickname'] = df['Nickname'].astype('str')+'_'+ df['Name'].astype('str')+'_'+ df['Num'].astype('str')
df['Password'] = hashAndSalt 
del df['Num']

records = df.to_records(index=False)
result = list(records)

result_process = list([tuple(row) for row in result])

filename = 'data'
outfile = open(filename,'wb')
pickle.dump(result_process,outfile)
outfile.close()