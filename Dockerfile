FROM golang:1.19-buster

EXPOSE 8000

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o gsl ./cmd/main.go

CMD ["./gsl"]