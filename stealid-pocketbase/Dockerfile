FROM docker.io/golang:latest

COPY . /pb
WORKDIR /pb

RUN go build -o pocketbase

EXPOSE 8080

CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]
