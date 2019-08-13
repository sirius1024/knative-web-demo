FROM golang:alpine AS builder
RUN echo $GOPATH
RUN apk add build-base --no-cache --update-cache --allow-untrusted
ENV GO111MODULE=on
RUN adduser -D -g 'root' appuser
WORKDIR /go/src/github.com/sirius1024/knative-web-demo
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o /bin/knative-web-demo

FROM alpine
RUN apk add tzdata --no-cache --update-cache --allow-untrusted && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
WORKDIR /data/app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/knative-web-demo .
RUN chown -R appuser:root /data/app
USER appuser
ENTRYPOINT [ "./knative-web-demo" ]
