FROM golang:1.20 AS builder

WORKDIR /plugin

COPY ./go.mod ./
COPY ./go.sum ./


RUN go mod download
RUN go mod verify

COPY *.go ./

RUN go build -trimpath -buildmode=c-shared -o mysql-output-plugin.so

FROM fluent/fluent-bit

COPY --from=builder /plugin/mysql-output-plugin.so /fluent-bit/etc/
COPY ./testdata/fluent-bit.conf /fluent-bit/etc/
COPY ./testdata/plugins.conf /fluent-bit/etc/

ENTRYPOINT [ "/fluent-bit/bin/fluent-bit" ]
CMD [ "/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf" ]