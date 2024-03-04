FROM golang:1.22-alpine

WORKDIR /api

RUN apk update && apk add make

RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3000

CMD ["air"]
# CMD ["go", "run", "./cmd/api/main.go"]
