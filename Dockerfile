# Builder image
FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY db.go .
COPY app.go .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o saved

# Production image
FROM golang:1.12-stretch
COPY --from=builder /app/saved /app/

VOLUME [ "/uploads" ]

EXPOSE 4444
ENTRYPOINT [ "/app/saved" ]
