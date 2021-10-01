FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./
COPY ./config ./config
COPY ./auth ./auth
COPY ./constant ./constant
COPY ./controller ./controller
COPY ./database ./database
COPY ./models ./models
COPY ./routes ./routes

RUN go build -o agent-allocation

CMD ./agent-allocation