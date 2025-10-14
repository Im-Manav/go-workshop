# Run the build.
FROM golang:1.25 AS build-stage

WORKDIR /src

# Cache and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy all the source code.
COPY . ./

# Build the application.
RUN go build -o app

# Build a small image for the application.
FROM gcr.io/distroless/base-debian12 AS runtime-stage

WORKDIR /

COPY --from=build-stage /src/app /app

# Expose the port the app runs on.
EXPOSE 8080
# Expose Prometheus metrics endpoint.
EXPOSE 9191

USER nonroot:nonroot

ENTRYPOINT ["/app"]
