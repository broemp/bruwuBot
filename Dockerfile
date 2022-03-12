FROM golang:1.18rc1-alpine3.15 AS build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o bot ./cmd/bot/main.go

# -----------------------------------

FROM alpine:3.15 AS final

WORKDIR /app
COPY --from=build /build/bot ./bot
ENTRYPOINT ["./bot"]
CMD ["-c", "config.json"]