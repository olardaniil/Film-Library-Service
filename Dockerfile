FROM golang:latest

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o film-library-service ./cmd/app/main.go

CMD ["./film-library-service"]





