FROM golang:1.11.2-alpine3.8 as BUILD

ENV GOOS=linux
ENV GOARCH=386

ADD server.go /home/server.go
WORKDIR /home
RUN go build server.go

VOLUME ["/static"]

FROM busybox
COPY --from=BUILD /home/server /home/
ENTRYPOINT ["/home/server"]
CMD ["/static"]