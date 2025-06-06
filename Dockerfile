FROM golang:1.24-alpine
LABEL authors="cnossa"
LABEL description="A simple supermarket API written in Go"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/supermarket-api

EXPOSE 8080
CMD ["/bin/supermarket-api"]