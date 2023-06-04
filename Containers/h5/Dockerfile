FROM golang:1.20
WORKDIR /app
COPY . .
EXPOSE 8080
RUN go build -o main ./cmd/container-app/main.go
CMD [ "./main" ]