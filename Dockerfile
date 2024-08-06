FROM golang:latest AS env

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
WORKDIR /app
RUN go mod download

FROM env AS build

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./profile /app/cmd/profile

FROM scratch
COPY --from=build /app/profile /app/profile

EXPOSE 4848
CMD ["/app/profile"]
