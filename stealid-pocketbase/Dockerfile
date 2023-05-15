FROM docker.io/golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./stealid-pocketbase/
RUN cd ./stealid-pocketbase/ && go mod download && go mod verify

COPY . ./stealid-pocketbase
RUN cd ./stealid-pocketbase/ && go build -v -o /usr/local/bin/pocketbase

# for implant building
RUN go install mvdan.cc/garble@latest
RUN apt-get update
RUN apt-get install -y mingw-w64

EXPOSE 8080

CMD ["pocketbase", "serve", "--http=0.0.0.0:8080"]
