# ビルド用ステージ
FROM golang:1.24-alpine AS builder
WORKDIR /app

# 依存解決
COPY go.mod go.sum ./
RUN go mod download

# ソースコピー＆ビルド
COPY . .
RUN go build -o server main.go

# 実行用ステージ
FROM alpine:latest
WORKDIR /app

# バイナリだけコピー
COPY --from=builder /app/server .

# listen ポート
EXPOSE 8080

# デフォルトコマンド
ENTRYPOINT ["./server"]
