FROM golang:1.20 AS builder

# cmetrics needs to be installed in the builder image
ARG CMETRICS_VERSION=0.5.8
ENV CMETRICS_VERSION=${CMETRICS_VERSION}
ARG CMETRICS_RELEASE=v0.5.8
ENV CMETRICS_RELEASE=${CMETRICS_RELEASE}

ARG PACKAGEARCH=amd64
ENV PACKAGEARCH=${PACKAGEARCH}

WORKDIR /plugin

COPY ./go.mod ./
COPY ./go.sum ./


RUN go mod download
RUN go mod verify

COPY *.go ./

ADD https://github.com/fluent/cmetrics/releases/download/${CMETRICS_RELEASE}/cmetrics_${CMETRICS_VERSION}_${PACKAGEARCH}-headers.deb external/
ADD https://github.com/fluent/cmetrics/releases/download/${CMETRICS_RELEASE}/cmetrics_${CMETRICS_VERSION}_${PACKAGEARCH}.deb external/
RUN dpkg -i external/*.deb

RUN go build -trimpath -buildmode=c-shared -o mysql-output-plugin.so

FROM fluent/fluent-bit

COPY --from=builder /plugin/mysql-output-plugin.so /fluent-bit/etc/
COPY ./testdata/fluent-bit.conf /fluent-bit/etc/
COPY ./testdata/plugins.conf /fluent-bit/etc/

ENTRYPOINT [ "/fluent-bit/bin/fluent-bit" ]
CMD [ "/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf" ]