FROM golang:1.13-buster

WORKDIR /src/pants/app
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ["./app"]
