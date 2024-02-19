FROM golang:1.21 AS builder

ENV GO111MODULE=on 
ENV GOPROXY=http://goproxy.cn,direct

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main .


FROM ubuntu:20.04 AS prod
WORKDIR /app
ARG PROJECT_NAME=uniquehr

COPY --from=builder /app/main ./${PROJECT_NAME}
COPY --from=builder /app/config.local.yml ./config.local.yml
COPY --from=builder /app/docs/ ./docs/

EXPOSE 3333

RUN apt-get -qq update &&\
    apt-get -qq install -y --no-install-recommends ca-certificates &&\
    echo "./${PROJECT_NAME} server" > ./run.sh &&\
    chmod u+x ./run.sh

CMD ./run.sh
