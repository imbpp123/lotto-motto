FROM golang:1.19-alpine AS app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN go build -trimpath -o bin/cli /app/main.go

CMD ["bin/cli"]
