FROM registry.freedomcore.ru/freedomcore/go:versioneer as builder
RUN mkdir -p /go/src/github.com/plexmediamanager/micro-database
COPY ./ /go/src/github.com/plexmediamanager/micro-database
WORKDIR /go/src/github.com/plexmediamanager/micro-database
RUN CGO_ENABLED=0 GOOS=linux  go build -ldflags="$(versioneer -package github.com/plexmediamanager/service)" -a -installsuffix cgo -o micro-database .

FROM registry.freedomcore.ru/freedomcore/go:alpine
COPY --from=builder /go/src/github.com/plexmediamanager/micro-database/micro-database /opt/micro-database
COPY --from=builder /go/src/github.com/plexmediamanager/micro-database/application.env /opt/application.env
RUN chmod +x /opt/micro-database
WORKDIR /opt
CMD ["./micro-database"]