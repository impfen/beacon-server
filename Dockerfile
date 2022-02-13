
  
FROM golang:1.16 as builder 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install ./...

FROM scratch

CMD [ "/beacon-server ]

COPY --from=builder /beacon-server
USER 12345
