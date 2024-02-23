#Based on https://stackoverflow.com/questions/52266332/manually-fetch-dependencies-from-go-mod
FROM golang:latest as builder
WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go app
COPY cmd/main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rinha-app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/rinha-app .

CMD [ "./rinha-app" ]