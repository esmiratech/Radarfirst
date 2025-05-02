# Grpc

This is a basic grpc service built with golang that contains an endpoint `/helloworld.Greeter/SayHello` which writes to a postgres database and returns a response. 

It also has an endpoint for a healthcheck: `grpc.health.v1.Health/Check` which pings the underlying postgres connection verifying connectivity.

 # Run locally

* Have go installed (go.dev/download)
* `go run . --port=:<port to listen on>` or `go build . && ./helloworld --port:<port to listen on>`

# Task

Update the dockerfile to serve the grpc service on port :5555

Needed Environment variables:

 * `PGHOST` - host name
 * `PGPORT` - port of postrgres server
 * `PGDATABASE` - database to connect to
 * `PGUSER` - postgres user to use
 * `PGPASSWORD` - password to use for the user

Flags needed to run the grpc server:

* `--port` - port for the gRPC service to listen on
