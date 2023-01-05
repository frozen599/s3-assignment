FROM golang:1.19-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o apiserver ./api/cmd/friends/

FROM scratch

COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]