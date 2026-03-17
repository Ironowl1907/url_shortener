FROM golang:1.25

WORKDIR /app

# Copy dependency files first for better layer caching
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY backend/ .

RUN go build -o server ./cmd/server/main.go

CMD ["./server"]
