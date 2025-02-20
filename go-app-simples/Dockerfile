FROM golang:1.18.1
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
CMD [ "./main" ]
