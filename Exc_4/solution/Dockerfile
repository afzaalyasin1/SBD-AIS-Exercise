  # todo
#  The Builder
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copy module files
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Run the build script

RUN sed -i 's/\r$//' ./scripts/build-application.sh && \
    chmod +x ./scripts/build-application.sh && \
    sh ./scripts/build-application.sh
# The Final Image 
FROM alpine:latest
WORKDIR /app

# Copy *only* the compiled binary
COPY --from=builder /app/ordersystem .

EXPOSE 3000
CMD ["./ordersystem"]