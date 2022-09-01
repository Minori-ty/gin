FROM golang:alpine as builder

WORKDIR /WWW/WWWROOT/
COPY ./ /WWW/WWWROOT/

RUN go env -w GOPROXY=http://goproxy.cn
RUN go build

FROM alpine

WORKDIR /WWW/WWWROOT/
COPY --from=builder /WWW/WWWROOT/gin /WWW/WWWROOT/

EXPOSE 4000

CMD /WWW/WWWROOT/gin