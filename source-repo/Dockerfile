FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM scratch

WORKDIR /app

COPY --from=builder /app/myapp .

CMD ["./myapp"]
