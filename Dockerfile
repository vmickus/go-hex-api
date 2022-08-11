############################
# STEP 1 build executable binary
############################
FROM golang:1.18 AS builder


WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api


############################
# STEP 2 build a small image
############################
FROM scratch AS runner

COPY --from=builder /app/api /api
COPY config.yaml /config.yaml

ENTRYPOINT ["/api"]
