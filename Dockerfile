FROM golang:1.23.2

# Set destination for COPY
WORKDIR /

COPY go.mod ./

RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Optional:

EXPOSE 8080

# Run
CMD ["/main"]