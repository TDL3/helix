FROM golang:1.16.2-alpine as build
WORKDIR /go/src/helix
ENV GO111MODULE=on 
ENV GOPROXY=https://goproxy.io,direct
ENV CGO_ENABLED=0
COPY . .
RUN go build -o server .

FROM alpine:latest
WORKDIR /etc/helix
ENV GIN_MODE=release
COPY --from=build /go/src/helix/config.yaml .
COPY --from=build /go/src/helix/resource .
COPY --from=build /go/src/helix/server .
EXPOSE 1080
ENTRYPOINT [ "/etc/helix/server" ]
