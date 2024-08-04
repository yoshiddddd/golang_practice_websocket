# ベースイメージを指定
FROM golang:1.18.3

# タイムゾーンの設定
ENV TZ=Asia/Tokyo

# Goの環境変数を設定
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app

# 作業ディレクトリを設定
WORKDIR ${ROOTPATH}

# モジュールファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションのソースコードをコピー
COPY src ./src

# ビルド用ディレクトリに移動
WORKDIR ${ROOTPATH}/src

# アプリケーションをビルド
RUN go build -o /go/bin/app main.go

# コンテナがリッスンするポートを指定
EXPOSE 8080

# コンテナの起動時に実行されるコマンド
CMD ["/go/bin/app"]
