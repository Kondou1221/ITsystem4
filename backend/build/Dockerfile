# とりあえずgolangのバージョンは最新を使用
FROM golang:1.21.5
# 作業用ディレクトリの作成
WORKDIR /go/src
# アップデートとgitのインストールと依存関係をインストール
RUN apt-get update \
&& apt-get install -y git \
&& go mod init back-go \
&& go get github.com/gin-gonic/gin \
&& go get github.com/jinzhu/gorm \
&& go get github.com/jinzhu/gorm/dialects/postgres