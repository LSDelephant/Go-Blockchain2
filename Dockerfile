FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o /go-blockchain ./cmd/node
EXPOSE 8000
CMD ["/go-blockchain"]
