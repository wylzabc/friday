FROM core.harbor.ebupt.com/library/builder-golang:1.12-alpine as builder

ADD . /friday

RUN cd /friday && make

# Pull first-demo into a second stage deploy alpine container
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /friday/build/bin/friday /friday/friday
EXPOSE 8080/tcp
CMD cd /friday && ./friday
