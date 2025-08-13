FROM golang:1.13
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN go build main.go
RUN mkdir output outputzip
CMD ["./main"]
