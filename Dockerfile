FROM golang:1.14-alpine as builder
WORKDIR /go/src/hd-wallet
COPY . .
RUN apk add --update git make
# RUN go get ./...
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/hd-wallet/hd-wallet ./hd-wallet
# COPY --from=builder /go/src/hd-wallet/keys ./keys
CMD ["/hd-wallet"]
EXPOSE 6000