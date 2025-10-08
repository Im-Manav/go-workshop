FROM golang:1.23 AS build-stage

# Build the Go application.
# This uses a larger image with the Go toolchain to build the application.

WORKDIR /app

# Creating a specific layer for the `go.mod` and `go.sum` files allows Docker to cache dependencies, so they're only downloaded again if they change.
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code into the container, this will change more frequently than the dependencies.
COPY . ./

# Go can interoperate with C libraries, so sometimes, you'll see bindings for C libraries.
# If you're not using them, you can set CGO_ENABLED=0 to disable cgo which can reduce the size of the final image and make your application more portable, since you don't have to have libc installed on the target system.
# However, this example uses sqlite3, which is a C library, so we need to enable cgo.

# The `-s` and `-w` flags are used to strip the debug information from the binary which produces a smaller binary - you still get stack traces, but they're not as detailed.

RUN go build -ldflags="-s -w" -o /app/api .

FROM ubuntu:jammy AS runtime-stage

# Now, we switch into the runtime stage, which is based on a smaller image.
# At runtime, we don't need the Go compiler etc.

# Using a smaller image reduces the size of the final image, but also reduces the attack surface.

# Copy the built application from the build stage into the runtime stage.
COPY --from=build-stage /app/api /usr/local/bin/api

ENTRYPOINT [ "/usr/local/bin/api" ]
