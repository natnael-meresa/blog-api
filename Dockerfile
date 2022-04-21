FROM golang:alpine
WORKDIR /
ADD . .
RUN go build -o bin/blog-api cmd/main.go


FROM alpine
WORKDIR /
COPY --from=builder /bin/blogapi .
RUN apk --no-cache add tzdata


EXPOSE 9000
CMD ["./bin/blog-api"]
