## Build
FROM golang:1.19-alpine as build

WORKDIR /app

COPY . .

RUN go build .

CMD [ "./sowhenthen" ]


## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=build ./app/sowhenthen /sowhenthen

ENV HOST=""
ENV PORT=""
ENV MONGO_URL=""

EXPOSE 8080

ENTRYPOINT [ "./sowhenthen" ]
