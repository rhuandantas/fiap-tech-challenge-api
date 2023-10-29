FROM golang:1.20-alpine3.17 as base
LABEL authors="rhuan"
RUN apk update
WORKDIR /app/tech_challenge
COPY go.mod go.sum ./
# separate in a sh file
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/swag
RUN go get -u github.com/google/wire/cmd/wire
RUN wire ./...
COPY wire_gen.go ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api

FROM alpine:3.17 as binary
COPY --from=base /app/tech_challenge/api .
EXPOSE 3000
CMD ["./api"]