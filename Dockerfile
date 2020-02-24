FROM golang:1.13

COPY . $GOPATH/src/github.com/awcodify/j-man/
WORKDIR $GOPATH/src/github.com/awcodify/j-man/

COPY config.yaml.example config.development.yaml
RUN go get -d -v ./...
RUN go build -o jman cmd/jmanager/main.go

ENTRYPOINT ["./jman"]
