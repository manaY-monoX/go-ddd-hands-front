# 使用するGoのバージョンとベースイメージを指定
FROM golang:1.23.2-bookworm
RUN apt-get update && apt-get install -y git
# アプリケーションのディレクトリを作成
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app