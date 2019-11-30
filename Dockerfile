# Builder layer.
FROM golang:1.13.4-buster as build

ENV GOPATH=/go

WORKDIR $GOPATH/src/github.com/tngranados/tech-challenge-time/
ADD . .

RUN make deps
RUN GOOS=linux go build -o app

# Runner layer.
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /go/src/github.com/tngranados/tech-challenge-time/app ./app
COPY --from=build /go/src/github.com/tngranados/tech-challenge-time/.env ./.env

CMD ["./app"]
