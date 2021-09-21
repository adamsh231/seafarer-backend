FROM golang:1.16

ENV GO111MODULE=on

ENV PKG_NAME=seafarer-backend
ENV PKG_PATH=$GOPATH/src/$PKG_NAME

RUN apt update -y && apt upgrade -y
RUN apt install git -y
RUN git config --global url."https://adamsh231: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@github.com".insteadOf "https://github.com"

RUN apt install wkhtmltopdf -y

WORKDIR $PKG_PATH/
COPY . $PKG_PATH/

RUN go mod vendor
WORKDIR $PKG_PATH/server/http

RUN go build main.go
EXPOSE 4000
CMD ["sh", "-c", "./main"]
