FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o alert-hub-go ./main.go

FROM debian:bookworm
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /root/
COPY --from=builder /app/alert-hub-go .
EXPOSE 8080
CMD ["./alert-hub-go"]
