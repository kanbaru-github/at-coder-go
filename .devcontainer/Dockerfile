# ベースイメージとして golang:1.20 を使用
FROM golang:1.20

# 必要なツールをインストール
RUN apt-get update && apt-get install -y \
    curl \
    git \
    && apt-get clean

# 作業ディレクトリを設定
WORKDIR /workspace
