FROM golang:1.14.2-alpine as builder
WORKDIR /go/src/hd-wallet
COPY . .
RUN apk add --update git make
RUN export GO111MODULE=on
RUN echo $GO111MODULE
RUN go mod vendor
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/hd-wallet/hd-wallet ./hd-wallet
# COPY --from=builder /go/src/hd-wallet/keys ./keys
CMD ["/hd-wallet"]
EXPOSE 6000