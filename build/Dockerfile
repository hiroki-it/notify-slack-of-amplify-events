#===================
# Global ARG
#===================
ARG GO_VERSION="1.16.0"
ARG LABEL="Hiroki <hasegawafeedshop@gmail.com>"

#===================
# Build Stage
#===================
FROM golang:${GO_VERSION} as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src

# インストールのキャッシュを活用するためにコピーしておく．
COPY go.mod go.sum ./

# ライブラリをインストールする．
RUN go get -u github.com/cosmtrek/air \
  && go mod download

COPY . .

# ビルドのアーティファクトを$HOME/go/binに配置する．
RUN go build -o ${HOME}/go/bin ./cmd

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/notify_slack_of_amplify_events/cmd/app /app

CMD ["/app"]
