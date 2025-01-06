FROM alpine:3

# Install system dependencies
RUN apk add ca-certificates

# Copy the binary and resources
WORKDIR /app
COPY ./webapi ./
COPY ./public ./

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
