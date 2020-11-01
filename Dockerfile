FROM golang:alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker-puller .


FROM alpine

RUN apk add --no-cache docker-cli

COPY scripts/docker-entrypoint.sh /entrypoint.sh
COPY --from=builder /app/docker-puller /docker-puller

ENTRYPOINT [ "/entrypoint.sh"]