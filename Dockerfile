FROM golang  AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
EXPOSE 8080  
WORKDIR /app

COPY . /app
RUN go get -d ./...
RUN go build -o server .

FROM scratch
COPY --from=builder /app .
CMD ["./server"]