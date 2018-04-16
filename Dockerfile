# build stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git \
  && go get -d -v github.com/cloudfoundry/go-socks5 \
  && apk del git
ADD . /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
EXPOSE 1080
ENTRYPOINT ./goapp
