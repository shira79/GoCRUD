FROM golang:latest

RUN mkdir /app
WORKDIR /app

# ModuleモードをON
ENV GO111MODULE=on
EXPOSE 8080

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.27.3
CMD ["air"]