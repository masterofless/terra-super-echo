FROM golang:1.16-alpine AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /go-super-echo

FROM alpine:latest

WORKDIR /
COPY --from=build /go-super-echo /go-super-echo
# RUN addgroup notroot
# RUN adduser -G notroot -D notroot
# USER notroot:notroot
EXPOSE 8080

CMD [ "/go-super-echo" ]
