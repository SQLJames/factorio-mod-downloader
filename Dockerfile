FROM cgr.dev/chainguard/go:latest as build

WORKDIR /app

# Copy all files except the ones listed in .dockerignore
COPY . .

# Download dependencies
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o factorio-mod-downloader ./cmd/cli
RUN chmod +x ./cmd/cli
# Start a new stage from scratch
FROM cgr.dev/chainguard/static:latest

# Copy the binary from the build stage
COPY --from=build /app/factorio-mod-downloader /usr/local/bin/factorio-mod-downloader

# Default command with no arguments
CMD ["/usr/local/bin/factorio-mod-downloader"]

# Command to override default behavior and pass arguments
ENTRYPOINT ["/usr/local/bin/factorio-mod-downloader"]