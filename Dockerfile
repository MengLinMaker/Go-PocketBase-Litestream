FROM bitnami/golang:1.21.6 as build

# Install Litestream for SQLite3 backup
ARG LITESTREAM_RELEASES="https://github.com/benbjohnson/litestream/releases" \
    LITESTREAM_VERSION="v0.3.13"
RUN CPU_ARCHITECTURE="$(arch)" && \
    echo "CPU Architecture:" $CPU_ARCHITECTURE && \
    case $CPU_ARCHITECTURE in \
        aarch64) curl -sL -o /litestream.deb "$LITESTREAM_RELEASES/download/$LITESTREAM_VERSION/litestream-$LITESTREAM_VERSION-linux-arm64.deb" ;; \
        *) curl -sL -o /litestream.deb "$LITESTREAM_RELEASES/download/$LITESTREAM_VERSION/litestream-$LITESTREAM_VERSION-linux-amd64.deb" ;; \
    esac && \
    apt install /litestream.deb

# Build Golang app with dependencies
COPY src/go.mod src/go.sum /
RUN export GOPROXY=https://proxy.golang.org && \
    go mod download -x
COPY src /
RUN export CGO_ENABLED=1 && \
    go build -x -o /server.bin /

# Minimise final stage docker image for production
# Smallest glibc alternative to Apline
FROM busybox:1.36.1 AS final

# Specify backup location secret
ENV LITESTREAM_ACCESS_KEY_ID \
    LITESTREAM_SECRET_ACCESS_KEY \
    S3_ENDPOINT \
    S3_DATA_BUCKET \
    S3_LOGS_BUCKET \
    STAGE

COPY --from=build /lib/**/libdl.so.2 /lib
COPY --from=build /usr/bin/litestream /bin
COPY --from=build /server.bin /
COPY etc /etc

EXPOSE 8080
CMD ["ash", "/etc/server.sh"]
