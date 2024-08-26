# STAGE 1
FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /go/brantas
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
RUN apk add --no-cache make
COPY . .
RUN go build -o main ./main.go

# STAGE 2
FROM alpine:3.19.0
# This is for add appuser in apine linux
# RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# USER appuser
WORKDIR /root/
COPY --from=builder /go/brantas/main .
COPY --from=builder /go/brantas/.env .
EXPOSE 8910
CMD ["nohup", "./main"]