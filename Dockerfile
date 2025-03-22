FROM golang:1.22.4-alpine AS builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY internal internal

COPY cmd cmd

RUN mkdir resources

RUN touch resources/database.db

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/sword-health /app/cmd/main.go

FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=builder /app/resources/database.db /app/resources/database.db

COPY --from=builder /app/sword-health .

EXPOSE 8080 

ENTRYPOINT [ "/app/sword-health" ]