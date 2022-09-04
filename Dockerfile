
#build stage

FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app .
EXPOSE 9090

CMD ["/app/main"]