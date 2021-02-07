FROM golang:1.15.2-alpine as builder

WORKDIR /

COPY http-server.go .

ARG ARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build http-server.go

FROM scratch
WORKDIR /
COPY --from=builder http-server .
ENTRYPOINT ["/http-server"]