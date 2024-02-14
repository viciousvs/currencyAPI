FROM golang:alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build cmd/app/main.go
EXPOSE 8080

CMD [ "/main" ]