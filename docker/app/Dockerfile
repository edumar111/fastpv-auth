FROM golang:1.11-alpine AS builder
RUN apk add --update git && apk add build-base  make gcc musl-dev linux-headers
RUN go get  github.com/edumar111/fastpv-auth
WORKDIR /go/src/github.com/edumar111/fastpv-auth

#RUN CGO_ENABLED=0 GOOS=linux go build   -o /go/bin/myapp *.go
RUN export GO111MODULE=on &&  go build  -o /go/bin/auth *.go


FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin /
COPY --from=builder /go/src/github.com/edumar111/fastpv-auth/settings /settings
##ENTRYPOINT ["/myapp"]
CMD ["/auth"]

# docker login
# docker build -t edumar111/fastpv-auth:v1.0.5 -f Dockerfile .
# docker push     edumar111/fastpv-auth:v1.0.5

# docker run  -itd  --name astpv-auth -p 8080:8080 --link fastpv-redis:fastpv-redis -e HOST_REDIS=fastpv-redis -e PORT_REDIS=6380 -e PASS_REDIS=fastpv123 -e GO_ENV=production edumar111/fastpv-auth:v1.0.3
