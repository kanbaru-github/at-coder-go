# AtCoderに合わせベースイメージとしてGo 1.20を使用
FROM golang:1.20

# 開発ユーザーを作成
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt-get update && apt-get install -y \
    curl \
    git \
    bash-completion \
    && groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME --shell /bin/bash \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# bash-completion を有効化
RUN echo "if [ -f /usr/share/bash-completion/bash_completion ]; then . /usr/share/bash-completion/bash_completion; fi" >> /home/vscode/.bashrc

# 作業ディレクトリを設定
WORKDIR /workspace

# デフォルトユーザーを設定
USER $USERNAME
