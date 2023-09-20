# This layer builds the webapp
FROM node as build-webapp

WORKDIR /app

COPY webapp/package.json webapp/package-lock.json ./

WORKDIR webapp

RUN npm ci

COPY webapp/ ./

RUN npm run build

# This layer builds the go application
FROM golang:latest AS build

WORKDIR /app

# Copy go module files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy prebuild webapp
COPY --from=build-webapp /app/webapp/dist ./webapp/dist

# Copy go source code
COPY . ./

# Build go application
RUN CGO_ENABLED=0 go build -o /persurl cli/main.go

FROM alpine AS runtime

WORKDIR /

COPY --from=build /persurl /persurl
COPY docker/entrypoint.sh /entrypoint.sh

EXPOSE 8060

ENV GIN_MODE=release

ENTRYPOINT ["sh", "/entrypoint.sh"]
