FROM golang:1.23rc2-alpine3.19 as build_fase

# Install required packages
WORKDIR /app

COPY main.go ./main.go
COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY handlers ./handlers
COPY factory ./factory
COPY models ./models

RUN go mod download
RUN go build -o entrypoint main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=build_fase /app/entrypoint ./entrypoint

CMD [ "./entrypoint" ]