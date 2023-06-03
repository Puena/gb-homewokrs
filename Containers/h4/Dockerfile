FROM golang:1.20
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/container-app/main.go
CMD [ "./main" ]