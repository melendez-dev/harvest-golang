FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

# install air for live reloading
RUN go install github.com/air-verse/air@latest

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 3002

CMD ["air"]