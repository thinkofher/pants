FROM golang:1.13-alpine

WORKDIR /src/pants/app
COPY . .

RUN go get github.com/pilu/fresh
RUN go get -d -v ./...

CMD ["fresh"]
