FROM golang as builder

WORKDIR /app

# Copy the go mod and sum
COPY go.mod go.sum ./
RUN go mod download

# Copy and build the source
COPY ./ ./
RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o webapi .

# ------------------------------------------------------------------------------------

FROM alpine:3

# Install system dependencies
RUN apk --update add curl ca-certificates libstdc++ libc6-compat binutils

# Copy built files from the build image
WORKDIR /app
COPY --from=builder /app/webapi /app/webapi
COPY --from=builder /app/public /app/public

# Declare environment variables
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

# App ports
EXPOSE 3001

# Set the app entrypoint
ENTRYPOINT ["/app/webapi"]
