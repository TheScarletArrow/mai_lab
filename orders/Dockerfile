FROM golang:alpine AS b
WORKDIR /a
ADD . /a
RUN cd /a && go build -o orders ./cmd/main.go

FROM alpine
WORKDIR /a
COPY --from=b /a/orders /a

EXPOSE 8080
ENTRYPOINT ./orders