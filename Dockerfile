FROM golang:latest AS builder
ENV GO111MODULE=on 
ENV GOPROXY=http://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main .

FROM scratch AS prod
WORKDIR /app
COPY --from=builder /app/main ./main
COPY --from=builder /app/config.yaml ./config.yaml
EXPOSE 3333
CMD ["/app/main"]
