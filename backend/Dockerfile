FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN go build -o server ./cmd/server/main.go

CMD ["./server"]
