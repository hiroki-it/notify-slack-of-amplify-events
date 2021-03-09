#===================
# Global ARG
#===================
ARG GO_VERSION="1.15"
ARG LABEL="Hiroki <hasegawafeedshop@gmail.com>"

#===================
# Build Stage
#===================
FROM golang:${GO_VERSION} as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR ${GOPATH}/src

# インストールのキャッシュを活用するためにコピーしておく．
COPY go.mod go.sum ./

# ライブラリをインストールする．
RUN go get -u github.com/cosmtrek/air \
  && go mod download -x

COPY . .

# ビルドのアーティファクトを/go/binに配置する．
RUN go build -x -o ${GOPATH}/bin ./cmd

CMD ["sh"]
