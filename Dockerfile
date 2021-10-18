FROM golang:1.16-alpine as build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN  go mod download

COPY . ./

RUN go build -o app
RUN chmod +x /app

EXPOSE 8080

ENTRYPOINT [ "/app/app" ]