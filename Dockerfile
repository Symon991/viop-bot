# syntax=docker/dockerfile:1

FROM golang AS build
WORKDIR /app

COPY . ./

RUN go mod download
RUN go build -o /bot

FROM alpine:3.17
WORKDIR /

RUN apk add libc6-compat
RUN apk add ffmpeg
RUN apk add yt-dlp

COPY --from=build /bot /bot
COPY --from=build /app/config.json /

EXPOSE 8080

CMD ["/bot"]