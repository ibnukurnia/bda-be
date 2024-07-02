FROM golang:1.21.6-alpine3.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOARCH="amd64" \
    GOOS=linux

RUN apk --no-cache add build-base

WORKDIR /build

COPY . .
RUN go build --ldflags "-extldflags -static" -o main .
RUN go run cmd/keygen/main.go

RUN cat .env | grep APP_KEY

FROM alpine:3.19 AS runner

RUN apk update && \
    apk upgrade && \
    apk --no-cache add tzdata openssl

WORKDIR /www

COPY --from=builder /build/main /www/
COPY --from=builder /build/.env /www/

# ENV GIN_MODE=release

EXPOSE 5000

CMD ["./main"]
