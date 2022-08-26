FROM golang:latest AS builder

LABEL maintainer="Simon"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
######## Start a new stage from scratch #######
FROM alpine:latest AS release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# Expose port 8080 to the outside world
EXPOSE 3000
# Command to run the executable
CMD ["./main"]