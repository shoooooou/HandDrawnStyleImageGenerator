# ベースイメージとしてGoを使用
FROM golang:1.22-alpine

# 作業ディレクトリを作成
WORKDIR /app

# ホストのGoモジュールファイルをコンテナにコピー
COPY go.mod go.sum ./

# Goモジュールをダウンロード
RUN go mod download

# ホストのソースコードをコンテナにコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# コンテナ起動時に実行されるコマンド
CMD ["./main"]