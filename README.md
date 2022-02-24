# EFishery Test App

## How To Run

Clone this repository to your machine

    gir clone https://github.com/tirenn/test-efishery.git

Move to folder **_/test-efishery_**

    cd test-efishery

To run Auth-App move to folder **_/auth-app_**

    cd auth-app

To run Fetch-App move to folder **_/fetch-app_**

    cd fetch-app

Then copy **_.env.example_** and rename to **_.env_**

    cp .env.example .env

Edit **_.env_** file to your preference then run command

    source .env

Lastly run command

    go run main.go

## How To Run With Docker Compose

Move to folder **_/test-efishery_**, then edit file **_docker-compose.yml_**

Make sure you have **_docker-compose_** installed on your machine, then run command

    docker-compose up

For the documentation you can see here
[auth-app](https://github.com/tirenn/test-efishery/blob/main/auth-api.md)
[fetch-api](https://github.com/tirenn/test-efishery/blob/main/fetch-api.md)

## How To Run Test

Move to folder **_/auth-app_** or **_/fetch-app_** then run command

    go test ./domain/...

## C4 Diagram

### Context Diagram

![context diagram](https://github.com/tirenn/test-efishery/blob/main/Context%20Diagram.jpeg?raw=true)

### Container Diagram

![container diagram](https://github.com/tirenn/test-efishery/blob/main/Container%20Diagram.jpeg?raw=true)
