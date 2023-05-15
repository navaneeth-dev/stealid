FROM docker.io/golang:1.20

WORKDIR /pb

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/pocketbase

# for implant building
RUN go install mvdan.cc/garble@latest
RUN apt install mingw-w64

EXPOSE 8080

CMD ["pocketbase", "serve", "--http=0.0.0.0:8080"]
