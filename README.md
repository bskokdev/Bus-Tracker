# Bus tracker service using MQTT broker

## Description

Simple HTTP web service which processes telemetries from MQTT server
[mqtt.hsl.fi](https://digitransit.fi/en/developers/apis/4-realtime-api/vehicle-positions/) written in Go<br/>
The service stores the telemetries from the MQTT server and enables [REST endpoints](#rest-endpoints) to browse through them via pagination.<br/>
It also enables endpoint to get the list of the nearest bus telemetries for given position.

See [how to run](#run) the service.

## Assumptions

Prior to building the service I had made the following assumptions.

- Web service should store all the telemetries from the subscription
- User is able to provide current location in form of latitude and longitude
- Web service would be used only in the same timezone

## Technologies used

- Go
- GORM - ORM to handle database tables in Go (Automigrations enabled)
- Docker

## How to run the service via docker-compose {#run}

The web service can be ran as a docker container specified by <b>docker-compose.yml file</b><br/>
To run the web service via the docker-compose you are going to need .env file with the following variables.

### Environmental variables

#### Database

- `DB_HOST` - host address for the database <br/>
  <b>Note: </b>has to be name of the database service in docker-compose if docker is used
- `DB_PORT` - port the database is going to run on
- `DB_USER` - database user
- `DB_PASSWORD` - password for the user
- `DB_NAME` - name of the database

#### HTTP server

- `HTTP_URL` - host address for the http server<br/>
  <b>Note: </b>has to be "0.0.0.0" for the docker environment runtime. Can be changed if docker not used
- `HTTP_PORT` - port that the http server will be exposed on

### Example .env configuration using docker-compose

```
# database connection
DB_HOST=telemetry-database
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=strongPw
DB_NAME=telemetries

# HTTP server
HTTP_URL=0.0.0.0
HTTP_PORT=3000
```

To build the container run following command: `docker-compose up -d --build` <br/>

## REST endpoints

### Get telemetries

`GET <host>:<port>/api/v1/telemetries`<br/>
`GET <host>:<port>/api/v1/telemetries?pageSize=20&page=1`<br/>

#### Query parameters<br/>

| Name     | Required | Default value | Description                    |
| -------- | -------- | ------------- | ------------------------------ |
| pageSize | false    | 20            | number of telemetries per page |
| page     | false    | 1             | number of the current page     |

### Get nearest buses

`GET <host>:<port>/api/v1/buses/nearest?lat=60.205183&lon=25.145534`<br/>

#### Query parameters<br/>

| Name | Required | Default value | Description                    |
| ---- | -------- | ------------- | ------------------------------ |
| lat  | true     | None          | latitude location of the user  |
| lon  | false    | None          | longitude location of the user |

### Testing

I have also added a new unit tests to test the most crutial parts of the application - calculating the distance between 2 latitudes and longitudes and page offset for the database queries.
`go test ./...`
