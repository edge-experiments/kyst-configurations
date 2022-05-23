FROM golang:1.17
WORKDIR /wrap4kyst
COPY . .
RUN go build -o wrap4kyst main.go

FROM alpine:3.8
COPY --from=0 /wrap4kyst/wrap4kyst /wrap4kyst
