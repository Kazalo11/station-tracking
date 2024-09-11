#build stage
FROM golang:1.22.6-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/app


#run stage
FROM alpine:latest
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]