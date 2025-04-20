FROM golang:alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o notification ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=build /app/notification .

CMD ["./notification"]
