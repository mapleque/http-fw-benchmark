FROM golang:1.18.7
USER root
WORKDIR /opt
COPY go.mod /opt/go.mod
COPY main.go /opt/main.go
RUN go build -o /opt/service /opt/main.go
ENTRYPOINT /opt/service
