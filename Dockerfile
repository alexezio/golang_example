FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o app

FROM alpine

COPY --from=builder /app/app ./

CMD ["./app"]

