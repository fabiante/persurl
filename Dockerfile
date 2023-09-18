# Build the application from source
FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /persurl cli/main.go

FROM alpine AS runtime

WORKDIR /

COPY --from=build /persurl /persurl

EXPOSE 8060

ENV GIN_MODE=release

ENTRYPOINT ["/persurl"]
CMD ["run"]
