# angular2-jwt-starter

## Introduction
This project is a simple [Angular 2](https://angular.io/) application demonstrating the use of the [angular2-jwt](https://github.com/auth0/angular2-jwt) helper library, by [Auth0](https://auth0.com/).  It contains a login component, a home component, a user service, and an authentication service.  The application communicates with a simple [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) server written in [Go](https://golang.org/).

You need both [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed, to be able to run this project.

## Installation
Type `npm install` in the project directory, and wait for all dependencies to be downloaded, before you proceed.

## How to start
First, you need to start the Docker containers, by typing `npm run start-backend`.  This will start two containers, one for the REST server, and one for the Redis server.  If you had any of these containers already running, they will be stopped first.

After starting the containers, start the Angular 2 application, by typing `npm run start-frontend`.  Then point your browser to `http://127.0.0.1:4200`.

## Developing
During development, you may need to restart the containers individually.  If you modify the REST server code, you need to rebuild its Docker image for the REST server, and restart it. This can be done by typing `npm run restart-rest`.  If you need to restart Redis (which will get rid of all the stored data, btw), type `npm run restart-redis`.

## Stopping the containers
The Docker containers will run until you stop them.  This can be done by typing `npm run stop-backend`.
