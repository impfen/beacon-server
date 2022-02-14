
  
FROM golang:1.16 as builder 

COPY . /beacon-server 

WORKDIR /beacon-server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM scratch

ENTRYPOINT [ "/beacon-server" ]

COPY --from=builder /beacon-server/beacon-server /beacon-server

EXPOSE 8090

USER 12345
