# backend/Dockerfile

FROM golang:1.24

WORKDIR /app

# Install gcc for CGo (go-sqlite3)
RUN apt-get update && apt-get install -y gcc

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app . && ls -lh app && echo "✅ Build complete"

EXPOSE 9400

CMD ["./app"]
