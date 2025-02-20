FROM golang:1.18.1 AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.21.3 AS final
WORKDIR /app
COPY --from=build /build/main .
CMD [ "./main" ]
