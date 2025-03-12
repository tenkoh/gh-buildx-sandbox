FROM golang:1.24.1 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w" -o /app/main

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/main /
EXPOSE 8080
CMD [ "/main" ]