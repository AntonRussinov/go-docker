#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
COPY ./config.json /go/bin/app/
RUN go get -d -v ./...
RUN go mod tidy
RUN go build -o /go/bin/app/ -v ./...


#final stage
FROM alpine:latest
EXPOSE 80
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app/test_exercise"]
LABEL Name=testexercise Version=0.0.1