FROM golang:1.12.5-alpine3.9 AS build-base
RUN apk add --no-cache git=2.20.1-r0
RUN apk --no-cache add ca-certificates
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger \
    && go get -u github.com/gobuffalo/packr/packr

FROM build-base AS build-env
RUN apk --no-cache add tzdata
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV D=/go/src/ahuigo.com/ahuigo/proj
COPY . $D/
WORKDIR $D/
RUN go mod download
RUN go mod vendor


FROM scratch
COPY --from=build-env /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Shanghai
COPY --from=build-env /bin/proj-server /proj-server
ENV GIN_MODE=release
EXPOSE 80
ENTRYPOINT [ "/proj-server" ]

