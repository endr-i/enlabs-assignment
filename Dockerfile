FROM golang:1.14

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .

RUN chmod +x ./wait-for-it.sh
RUN go mod tidy
RUN mkdir -p _dist
RUN go test && go build -o=_dist/app .

CMD ["/go/src/app/_dist/app"]
