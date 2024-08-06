# ベースイメージを指定
FROM golang:1.18.3

# タイムゾーンの設定
ENV TZ=Asia/Tokyo

# Goの環境変数を設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app

WORKDIR ${ROOTPATH}

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src

WORKDIR ${ROOTPATH}/src

RUN go build -o /go/bin/app main.go

EXPOSE 8080

CMD ["/go/bin/app"]
