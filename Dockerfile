FROM golang:1.17.0 as build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
ADD . .
RUN go mod tidy && go build

FROM alpine
WORKDIR /simple-dns
COPY --from=build /build/simple-dns ./
RUN apk add --update curl && rm -rf /var/cache/apk/*
HEALTHCHECK --interval=5s --timeout=3s   CMD curl -fs http://localhost/healthz || exit 1
EXPOSE 80 53
ENTRYPOINT ["./simple-dns"]