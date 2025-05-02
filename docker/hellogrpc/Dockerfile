# step1
# ---------------------

FROM golang AS build

RUN apt-get update
RUN apt-get install -y protobuf-compiler

WORKDIR /app

COPY ./init/go.mod ./init/go.sum ./

RUN go mod download
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install github.com/grpc-ecosystem/grpc-health-probe@latest

COPY ./init/*.go ./
COPY ./init/proto/*.proto ./proto/

RUN protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./proto/*.proto 

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o ./helloworld
RUN chmod +x ./helloworld


# step2
# ---------------------

FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=build /app/helloworld .

USER nonroot:nonroot

CMD ["./helloworld", "-port", ":5555"]