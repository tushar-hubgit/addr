# ADDR REST API

ADDR API is a simple public IP address API, easy enough to integrate into any application in seconds.

## Run the app

    go run main.go

# REST API

The REST API to the example app is described below.

## Get Public IP (Plain Text)

### Request

`GET /`

    curl -i http://localhost:8080

### Response

    HTTP/1.1 200 OK
    Date: Sat, 29 Oct 2022 13:31:17 GMT
    Status: 200 OK
    Connection: close
    Content-Type: text/plain; charset=utf-8
    Content-Length: 9

    127.0.0.1

## Get Public IP (JSON)

### Request

`GET /?format=json`

    curl -i -H 'Accept: application/json' http://localhost:8080?format=json

### Response

    HTTP/1.1 200 OK
    Date: Sat, 29 Oct 2022 13:31:17 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 19

    {"ip": "127.0.0.1"}