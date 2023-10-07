# builder
FROM golang:1.21-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CG_ENABLED=0 go build -o crudTemplate ./cmd/api
RUN chmod +x /app/crudTemplate

# tiny exec img
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/crudTemplate /app
EXPOSE 80
CMD ["/app/crudTemplate"]