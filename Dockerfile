FROM golang:1.19 as builder
WORKDIR /go/src/github.com/rbender-hbo/hello-go-rest
COPY ./ .
RUN make build

FROM gcr.io/distroless/base
WORKDIR /
COPY --chmod=555 --from=builder /go/src/github.com/rbender-hbo/hello-go-rest/build/start-server .
EXPOSE 8080
CMD ["./start-server"]
