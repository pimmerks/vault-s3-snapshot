FROM golang:1.18 AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN mkdir /build
WORKDIR /build

COPY . .

RUN go mod download
RUN go build \
        -a \
        -trimpath \
        -ldflags "-s -w -extldflags '-static'" \
        -tags 'osusergo netgo static_build' \
        -o /build/out/vault-s3-snapshot \
        ./main.go

FROM alpine

WORKDIR /
COPY --from=builder /build/out/vault-s3-snapshot /vault-s3-snapshot

ENTRYPOINT ["/vault-s3-snapshot"]
