FROM golang:1.18.1
WORKDIR /app
RUN groupadd appgroup && useradd -r app -g appgroup
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
USER app
CMD [ "./main" ]
