FROM golang:1.19-alpine

WORKDIR /app

COPY *.json ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-server

CMD [ "/go-server" ]