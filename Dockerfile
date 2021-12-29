FROM golang:1.17-alpine

RUN go version
ENV GOPATH=/
WORKDIR /aquila_db

COPY ./ ./

# install psql
RUN apk update
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# install dependencies
RUN go mod download
# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go get github.com/githubnemo/CompileDaemon

# RUN go build -o aquila_db ./cmd/main.go
# CMD ["./main"] # --- to run vithout reloading

ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command=./main
# ENTRYPOINT CompileDaemon --build="go build -o main cmd/main.go" --command=./main