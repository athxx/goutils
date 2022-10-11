FROM golang:1.17 as builder
ENV GOPROXY=https://goproxy.cn,direct GO111MODULE=on
# Copy the Go Modules manifests
WORKDIR /soulma/svr
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download
# Copy the go source
# COPY main.go main.go
COPY app app
COPY cmd cmd
# Build
#sed -e 's|http://deb.debian.org|https://mirrors.tuna.tsinghua.edu.cn|g' -e 's|http://security.debian.org|https://mirrors.tuna.tsinghua.edu.cn|g' -i.bak /etc/apt/sources.list && \apt update && apt install libturbojpeg0 libpng16-16 && \
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /ox cmd/ox.go

#this base image cannot run bin file zyos
FROM scratch
COPY --from=builder /ox /ox
COPY etc/.env.dev /.env
WORKDIR /
ENTRYPOINT [ "/ox" ]