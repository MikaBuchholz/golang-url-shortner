# Url Shortner Service

This service offers shortning of links and getting back the url of said shortened url

## Dependencies

Prepare

Install
    Docker / Docker-Desktop
    Make (if not present)
    sqlx-cli
    Go

# All-In-One setup

This will set everything up and run the server / postgres DB

    make bootstrap

# Individual commands

Build the server image

    make build-server-image

Run the server image

    make run-server-image

Verify the server and postgres DB are working properly

    make mock-post

Expected return is

    {"id":"d4c9d9","view_count":0}

- Anything else is wrong

## How to use

### API ENDPOINTS

- server is running on <http://localhost:8080> by default

### Create new Shortened URL

Example POST request

    http://localhost:8080/api/v1/url/new

**Must** provide in request body:

    {"payload": "your_url"}

**Return on success**

    {"id":"some_id","view_count": your_view_count}

- Error reponse on failure

Example GET request

    http://localhost:8080/api/v1/url/{your_id}

- {your_id} is to be replaced with any id that was returned by the **POST** request

- No request body needed

**Return on success**

    {"true_url":"your_original_url","view_count": your_view_count}

- Error reponse on failure

## TODO

- ~~Update view count on **GET** request for any shortened url~~

- Deploy to google cloud / any service

- Build frontend to use this service

- Add Swagger docs
