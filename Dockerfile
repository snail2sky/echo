FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /data0/apps

RUN git clone https://github.com/snail2sky/echo.git 

WORKDIR /data0/apps/echo

RUN go build

FROM alpine:latest

WORKDIR /data0/apps/echo

COPY --from=builder /data0/apps/echo/echo .

ENTRYPOINT [ "./echo" ]