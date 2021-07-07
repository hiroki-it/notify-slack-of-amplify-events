#===================
# Global ARG
#===================
ARG GO_VERSION="1.16.5"
ARG LABEL="Hiroki <hasegawafeedshop@gmail.com>"

#===================
# Build Stage
#===================
FROM golang:${GO_VERSION} as build

# ビルドするバイナリファイルをスタティックリンク化する．
ENV CGO_ENABLED=0

WORKDIR /go/src

# インストールのキャッシュを活用するためにコピーしておく．
COPY go.mod go.sum ./

# ライブラリをインストールする．
RUN go get github.com/cosmtrek/air@v1.15.1

COPY . .

# ビルドのアーティファクトを/go/binに配置する．
RUN go build -x -o /go/bin ./cmd

#===================
# Base Stage
#===================
FROM golang:${GO_VERSION}-alpine as base

WORKDIR /go/src

# /go/binにパスを通す．
ENV PATH $PATH:/go/bin

COPY --from=build /go ../

#===================
# Develop Stage
#===================
FROM base as dev
LABEL maintainer=${LABEL}

CMD ["/go/bin/cmd"]

#===================
# Staging Stage
#===================
FROM base as stg
LABEL maintainer=${LABEL}

CMD ["/go/bin/cmd"]

#===================
# Production Stage
#===================
FROM base as prd
LABEL maintainer=${LABEL}

CMD ["/go/bin/cmd"]
