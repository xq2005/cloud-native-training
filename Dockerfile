FROM docker.io/golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/multiply-example .
FROM scratch AS bin
COPY --from=build /out/multiply-example /

CMD ["/multiply-example"]
