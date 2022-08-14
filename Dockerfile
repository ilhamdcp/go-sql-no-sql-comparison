FROM golang:1.18

WORKDIR ~/go-sql-no-sql-comparison

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go run main.go

EXPOSE 8000