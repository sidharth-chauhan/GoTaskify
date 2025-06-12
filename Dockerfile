FROM golang:alpine 

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o /app/main .

EXPOSE 8080

CMD ["/app/main"]
