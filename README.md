# angular2-jwt-starter

This project is a simple [Angular 2](https://angular.io/) application demonstrating the use of the [angular2-jwt](https://github.com/auth0/angular2-jwt) helper library,
by [Auth0](https://auth0.com/).  It contains a login component, a home component, a user service, and an authentication service.  The application communicates with a simple
[REST](https://en.wikipedia.org/wiki/Representational_state_transfer) server written in [Go](https://golang.org/).

To run this project, you need both [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed.  The `app.yml` configuration file
builds and starts two containers, the REST server, and a Redis server.

Starting an stopping these containers can be done with `npm run start-backend` and `npm run stop-backend`, respectively.

After starting the containers, you can run the frontend Angular 2 application, with `npm run start`.  Point your browser to `http://127.0.0.1:4200` to start the application.
  
