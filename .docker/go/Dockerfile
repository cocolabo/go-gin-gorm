FROM golang:1.16.7-alpine3.14

WORKDIR /app
COPY ./api/go.mod .
COPY ./api/go.sum .

RUN apk update \
  && apk add --no-cache git

RUN go get -u github.com/cosmtrek/air

RUN go mod download

COPY ./api .

CMD ["air", "-c", ".air.toml"]