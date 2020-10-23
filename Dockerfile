FROM golang:alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpinelinux/docker-cli

RUN mkdir /app
WORKDIR /app/

COPY --from=builder /app/main .

CMD ["/app/main"]