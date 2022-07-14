FROM golang:1.11-alpine AS builder
RUN apk add --update git && apk add build-base  make gcc musl-dev linux-headers
WORKDIR /go/src/github.com/edumar111
RUN git clone  https://github.com/edumar111/fastpv-auth.git
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
# docker build -t edumar111/fastpv-auth:1.0.0 -f Dockerfile .
# docker push     edumar111/fastpv-auth:1.0.0

# docker run  -it  --name astpv-auth -p 8080:8080 --network fastpv  -e HOST_MYSQL=mysql-service-svc -e  HOST_REDIS=redis-service-svc -e GO_ENV=production edumar111/fastpv-auth:1.0.0
