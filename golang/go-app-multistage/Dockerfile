FROM golang:1.18.1 AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.21.3 AS app
WORKDIR /app
RUN addgroup -S appgroup && adduser -S -D -H app -G appgroup
COPY --from=build --chmod=755 --chown=app:appgroup /build/main .
USER app
CMD [ "./main" ]
