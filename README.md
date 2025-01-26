# Organs Demo

## This is a demo of a Go/Gin REST api that delivers national Organ Donor waitlist data for use in visualizations
- The data is from [OPTN](https://optn.transplant.hrsa.gov/data/view-data-reports/national-data/)
- It has been dumped to CSV and imported into Firebase and is not real-time
- To run this application requires the Firebase organs-demo-api-key.json which is not provided here: email brandon.pliska@gmail.com for the password to the organs-demo-api-key.zip file

#### Instructions
- Obtain password for api key zip file
- Install api key json file from Firestore
    - Setup a firebase account and project
    - Setup a firestore database
    - Setup a firestore user with the proper permissions
    - Setup a firestore service account with the proper permissions
    - Download the service account json file
    - Unzip the file and copy the json file to the root of this repo and name it organs-demo-api-key.json
- Have Go 1.20+ installed
- Run: `make run` to build and start the server
  - Alternatively, cd into the `go` directory and run: `go run main.go`
- Crack open `http://localhost:8080/api/organs` in your browser

#### Project Status
This is primarily a Go-based project, with a Rust implementation currently in development as an alternative version.

#### TODO
- TODO: graphql server
- TODO: Svelte frontend
- TODO: Vue frontend 
- TODO: Flutter frontend
- TODO: React frontend
- TODO: Angular frontend

The output looks like this: 

![](output.png)