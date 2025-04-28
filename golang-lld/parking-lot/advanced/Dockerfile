FROM golang:1.24-alpine AS builder
RUN apk add --no-cache gcc musl-dev sqlite-dev
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 go build -o parking_app .

FROM alpine:3.18
RUN apk add --no-cache sqlite sqlite-libs
WORKDIR /app
COPY --from=builder /app/parking_app .
EXPOSE 8080
CMD ["./parking_app"]