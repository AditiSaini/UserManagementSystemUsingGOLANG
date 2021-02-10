# aditi-entry
The project consists of three major components:
1. Client Web application- React
2. Server code (TCP and HTTP)- Golang
3. Scripts for stress testing and adding data in the DB- Python and bash

## 1. Client Web Application 
1. Go to the root directory of the client web aplication, inside the client folder 
    - `cd client`
2. Build the node project using, 
    - `npm run build `

## 2. Creating DB records using the scripts
1. Go the scripts/addEntryInDB folder 
2. Run the command to create a pickled file for data to be added into the db 
    - `python3 getData.py`
3. To add data into the existing Profile table in the users db
    - `python3 addSQLData.py`

## 2. Startup instructions for the project using the script 
1. Go to the scripts/ folder and make the bash scripts executable, 
    - `chmod +x startServers.bash`
    - `chmod +x stopServers.bash`
    - `chmod +x createBinaries.bash`
2. Run the following command to create an excutable files
    - `./createBinaries.bash`
2. Run the below commands to start and stop the servers
    - `./startServers.bash`
    - `./stopServers.bash`



