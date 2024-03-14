FROM golang:1.22.1-alpine AS builder
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
COPY ./public ./

ENV AUTH_GITHUB_ID=""
ENV AUTH_GITHUB_SECRET=""
ENV AUTH_GOOGLE_ID=""
ENV AUTH_GOOGLE_SECRET=""
ENV AUTH_SECRET=""
ENV DATABASE_URI=""
ENV EMAIL_FROM=""
ENV EMAIL_SERVER=""
ENV GEOLITE2_DB=""
ENV PROVAR_API_URL=""
ENV PROVAR_APP_URL=""
ENV PORT=3001

EXPOSE 3001

ENTRYPOINT ["/app/webapi"]
