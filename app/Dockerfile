FROM golang:1.13-alpine

RUN apk add --no-cache make git

WORKDIR /src/pants/app
COPY . .

RUN go get -d -v ./...

CMD ["make", "fresh"]
