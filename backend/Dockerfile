FROM golang:1.21.6

ARG ATLAS_URI

ENV APP_HOME /go/src/opportunity
ENV ATLAS_URI ${ATLAS_URI}
ENV GIN_MODE=release

WORKDIR "$APP_HOME"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -tags=jsoniter -o opportunity 

EXPOSE 8080

CMD ["./opportunity"]
