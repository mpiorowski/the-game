FROM golang:1.19

WORKDIR /usr/src/app

COPY ./go.sum .
COPY ./go.mod .
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

