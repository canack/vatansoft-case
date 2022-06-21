FROM golang:1.18 as goenv

COPY . /todo-api
WORKDIR /todo-api

RUN go mod tidy
RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

FROM gcr.io/distroless/static

COPY --from=goenv /todo-api/todo-api /

CMD ["/todo-api"]