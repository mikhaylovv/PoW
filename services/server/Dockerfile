FROM golang:1.20 as builder
WORKDIR /app
ADD . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/server services/server/cmd/main.go

FROM alpine:latest
LABEL ms="client"
COPY --from=builder /go/bin/server ./server
EXPOSE 8080
CMD ["./server"]
