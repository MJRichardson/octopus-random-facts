# Build
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o /octopus-facts-service cmd/octopus-facts-service/main.go

# Runtime
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /octopus-facts-service /octopus-facts-service
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/octopus-facts-service"]