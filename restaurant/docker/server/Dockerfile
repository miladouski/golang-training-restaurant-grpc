FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build ./restaurant/order_server/cmd/order_server/server.go
CMD ["./server"]