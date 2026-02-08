FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# distroless for security + speed + size
FROM gcr.io/distroless/static:nonroot

WORKDIR /

USER 65532:65532
ENTRYPOINT [ "/mirage" ]


