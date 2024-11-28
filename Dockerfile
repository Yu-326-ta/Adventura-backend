# dev
FROM golang:1.22-bullseye AS dev
WORKDIR /work/adventura_backend
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./ ./
RUN make mod build-linux

# release
FROM alpine AS release
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=dev /work/adventura_backend/build/adventura_backend-linux-amd64 /usr/local/bin/adventura_backend
EXPOSE 8080
ENTRYPOINT ["adventura_backend"]
