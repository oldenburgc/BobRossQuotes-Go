FROM golang:1.24-alpine AS build-stage

#set destination for copy
WORKDIR /app

#Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code into the container
COPY *.go ./
COPY templates/ ./templates/
COPY data/ ./data/

RUN go build -o /bobross-go

#Deploy application binary into a lean image
FROM amazonlinux:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /bobross-go /bobross-go
COPY --from=build-stage /app/templates/ /templates/
COPY --from=build-stage /app/data/ /data/

EXPOSE 8080

#USER ec2-user:ec2-user

ENTRYPOINT ["/bobross-go"]

HEALTHCHECK --interval=5s --timeout=3s \
  CMD curl --fail http://localhost:8080/ || exit 1