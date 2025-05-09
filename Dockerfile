FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/air-verse/air@latest

COPY . .

COPY .env .

RUN go build -o main ./cmd

EXPOSE 8000

CMD ["air"]