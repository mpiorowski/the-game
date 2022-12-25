FROM golang:1.19

WORKDIR /usr/src/app

COPY ./server/go.sum ./
COPY ./server/go.mod ./
RUN go mod download && go mod verify

COPY ./server ./
RUN go build -v -o /usr/local/bin ./...

CMD ["server"]

