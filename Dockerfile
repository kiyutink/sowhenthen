## Build
FROM --platform=$BUILDPLATFORM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN GOOS=linux GOARCH=amd64 go build .


## Deploy
FROM alpine:latest

WORKDIR /app

COPY --from=build ./app/sowhenthen ./sowhenthen
COPY --from=build ./app/templates ./templates

ENV HOST=""
ENV PORT=""
ENV MONGO_URL=""

EXPOSE 8080

ENTRYPOINT [ "./sowhenthen" ]
