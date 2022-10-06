FROM golang:1.19 AS app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN go build -trimpath -o bin/display_eurojackpot cmd/eurojackpot/main.go
RUN chmod +x bin/display_eurojackpot

CMD ["bin/display_eurojackpot"]
