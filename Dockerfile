FROM golang:1.12-alpine as builder
WORKDIR /go/src/hd-wallet
COPY . .
RUN apk add --update git make
RUN go get -u github.com/Masterminds/glide
RUN glide install
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/hd-wallet/hd-wallet ./hd-wallet
# COPY --from=builder /go/src/hd-wallet/keys ./keys
CMD ["/hd-wallet"]
EXPOSE 6000