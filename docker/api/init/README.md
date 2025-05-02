# API

This python API contains one endpoint `/hello` which accepts `POST`. API communicates with a gRPC service (listening on a specified port).

## Example payload

```
{ "name": "bob" }
```

# Build app locally

* Have python 3.10+ installed
* `pip install -q -r requirements.txt`
* `python app.py` 

# How proto files were generated

* `python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. helloworld.proto`

# Task

Create a containerized API which listens on http://localhost:8080 from the host OS.

Environment variables needed:

* `GRPC_ENDPOINT` - name and port of the gRPC service.
