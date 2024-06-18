FROM golang:1.22-alpine
RUN apk add build-base
COPY /backend /app/
WORKDIR /app
EXPOSE 8111
ENV GO111MODULE=on
RUN go mod download
RUN go mod tidy
RUN ls -la /app/server/
CMD go run /app/server/cmd/main.go
