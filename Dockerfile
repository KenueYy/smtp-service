FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./

# Копируем весь код после go.mod, чтобы go уже умел резолвить пути
COPY . .

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

# Просто сразу собираем; go принудительно скачает недостающие модули
RUN go build -v -o smtp-service ./cmd/server

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/smtp-service /app/smtp-service

CMD ["/app/smtp-service"]