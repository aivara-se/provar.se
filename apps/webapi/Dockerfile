FROM golang:1.21-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./webapi

FROM alpine:3
RUN apk add ca-certificates
WORKDIR /app
COPY --from=builder /app/webapi ./
COPY ./GeoLite2-City.mmdb ./
ENV MONGO_URI=""
ENV PORT=3001
EXPOSE 3001
ENTRYPOINT ["/app/webapi"]
