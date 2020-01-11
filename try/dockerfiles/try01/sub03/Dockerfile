FROM golang:1.13.6-alpine3.11

COPY . .
RUN ls -l
RUN go build -o try01sub03
RUN ls -l
ENTRYPOINT ["try01sub03"]
