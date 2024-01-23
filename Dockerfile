FROM golang
# Work directory
WORKDIR /progrma

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Starting our application
CMD ["go", "run", "cmd/main.go"]

# Exposing server port
EXPOSE 8080