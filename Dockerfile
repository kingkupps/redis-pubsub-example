# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /redis-pubsub-example

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /redis-pubsub-example /redis-pubsub-example

EXPOSE 1932

USER nonroot:nonroot

ENTRYPOINT ["/redis-pubsub-example"]