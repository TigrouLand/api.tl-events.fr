FROM alpine:3.16 AS build

WORKDIR /app/go/

RUN apk update
RUN apk upgrade
RUN apk add --update go

ADD . .
ENV GOPATH /app

RUN go get
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o eventsstats

FROM alpine:3.16

WORKDIR /app

COPY --from=build /app/go/eventsstats /app/eventsstats
RUN chmod +x ./eventsstats

CMD ["./eventsstats"]
