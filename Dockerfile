FROM golang:alpine
WORKDIR $GOPATH/src/server3
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o server3 .
EXPOSE 8083
ENTRYPOINT  ["./server3"]