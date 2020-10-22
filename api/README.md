# Interacting with the API

Interacting with the `go-osmand-tracker` API endpoint is very easy. Simplicity is king.

## Getting started: Basic interaction examples

Once the server is running and receiving location updates from the OsmAnd app, assuming default settings, the API should be available at `http://localhost:8080`.

**Submitting location updates**

Manual location updates (without using the OsmAnd app) can be made using any tool that can generate simple HTTP GET requests, for example, Postman, a web browser or `curl`.

An example of an update request in `curl` is shown below:

```sh
curl -I 'http://localhost:8080/submit?lat=48.858370&lon=2.294481&timestamp=1600000000000&hdop=1&altitude=10&speed=12.3456'
```

A successful location update does not return any data, but the HTTP response code should be `204 No Content`.

**Retrieving the last known location**

The last known location can be retrieved using the `/retrieve` endpoint. An example of retrieving the last known location with `curl` is shown below:

```sh
curl -s 'http://localhost:8080/retrieve'
```

Instead of using `curl`, you could also open the url in a browser. Modern browsers like Mozilla Firefox will show an output in colour.

The response will look similar to this:

```json
{
  "ID": 6538,
  "Timestamp": 1602026742,
  "Data": {
    "latitude": 48.85837,
    "longitude": 2.294481,
    "timestamp": 1600000000000,
    "hdop": 1,
    "altitude": 10,
    "speed": 12.3456
  }
}
```

**Retrieving multiple locations (location history)**

By adding the `count` parameter tot the `/retrieve` endpoint, you could retrieve a location history.

An example for retrieving the last three known locations:

```sh
curl -s 'http://localhost:8080/retrieve?count=3'
```

The response would look similar to something like this:

```json
[
  {
    "ID": 102,
    "Timestamp": 1602512754,
    "Data": {
      "latitude": -4.916161356915467,
      "longitude": -160.2222180986629,
      "timestamp": 1602512754000,
      "hdop": 1.1547510249347204,
      "altitude": 134.20841187793522,
      "speed": 64.77147256992421
    }
  },
  {
    "ID": 101,
    "Timestamp": 1602512754,
    "Data": {
      "latitude": 33.97124421257169,
      "longitude": 160.0357362345378,
      "timestamp": 1602512754,
      "hdop": 13.786868053846266,
      "altitude": 1681.498893225688,
      "speed": 53.014155880294894
    }
  },
  {
    "ID": 100,
    "Timestamp": 1602512753,
    "Data": {
      "latitude": -60.01392249824546,
      "longitude": 98.42609160095583,
      "timestamp": 1602512753,
      "hdop": 14.713819682956101,
      "altitude": 885.1180921964191,
      "speed": 51.32874621010568
    }
  }
]
```

## Using Postman

There also is a [Postman collection][] available. This Postman collection can be used to submit and retrieve location updates in an easy way.

## OpenAPI 3.0 Specification

The OpenAPI 3.0 specification (formerly called 'Swagger') is available at [/api/swagger.yaml](./swagger.yaml). This file could be easily imported into the Swagger Editor, for example at https://editor.swagger.io/. Please keep in mind that the OpenAPI specification is leading, but the actual API offered by the application might be behind. This is a natural result of software development, i.e. the requirements and/or specifications change over time and the software needs to be (partially) (re)written to conform to those changes.

### Running the Swagger editor locally

It is also possible to run the Swagger editor locally. Assuming a working installation of Docker, this is what you need:

```sh
docker pull swaggerapi/swagger-editor
docker run -d -p 80:8080 swaggerapi/swagger-editor
```

After this, the Swagger Editor should be available at https://localhost/.

[Postman collection]: https://documenter.getpostman.com/view/5679145/TVRhbpN8 "The Postman collection for go-osmand-tracker"
