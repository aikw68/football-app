FROM golang:1.17.2
# ワーキングディレクトリの設定
WORKDIR /football
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /football
WORKDIR cmd
CMD ["go", "run", "main.go","handler.go"]
EXPOSE 80