# -----------------------------------------------------------------------------
# step 1: build go app
FROM golang:1.13.6-alpine3.11 as build-step

RUN mkdir /go-app
WORKDIR /go-app
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/go-app

# -----------------------------------------------------------------------------
# step 2: exec
FROM scratch
COPY --from=build-step /go/bin/go-app /go/bin/go-app
EXPOSE 8080
ENTRYPOINT ["/go/bin/go-app"]