#===================
# Global ARG
#===================
ARG GO_VERSION="1.15"
ARG ALPINE_VERSION="3.13"
ARG LABEL="Hiroki <hasegawafeedshop@gmail.com>"

#===================
# Build Stage
#===================
FROM golang:${GO_VERSION} as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/src

# インストールのキャッシュを活用するためにコピーしておく．
COPY go.mod go.sum ./

# ライブラリをインストールする．
RUN go get github.com/cosmtrek/air@v1.15.1 \
  && go mod download -x

COPY . .

# ビルドのアーティファクトを/go/binに配置する．
RUN go build -x -o /go/bin ./cmd

FROM alpine:${ALPINE_VERSION}

WORKDIR /go

COPY --from=build /go/bin ./bin

# /go/binにパスを通す．
ENV PATH $PATH:/go/bin

CMD ["/bin/sh"]
