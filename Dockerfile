FROM golang:alpine AS builder

LABEL maintainer="Nanang Suryadi <nanang.suryadi@telkom.co.id>"

RUN set -eux && apk --update --no-cache add ca-certificates upx

WORKDIR /project

COPY cmd cmd
COPY pkg pkg

# Copy and download dependencies.
COPY go.mod go.sum ./
RUN go mod tidy

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags dynamic -ldflags="-s -w" -o ./build/app ./cmd/bin/main.go

COPY ./config.yaml ./config.yaml

# Compress binary
RUN set -eux && upx "./build/app" && upx -t "./build/app" && chmod +x "./build/app"

# Clear cache and app files
RUN set -eux && apk del ca-certificates upx && rm -rf /var/cache/apk/*

FROM alpine:3.16
RUN date
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN date

RUN apk add --no-cache

RUN mkdir /.data

COPY --from=builder ["/project/build/app", "/app"]
COPY --from=builder ["/project/config.yaml", "/config.yaml"]

EXPOSE 8007

CMD [ "./app" ]