FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN go build -o blog-api cmd/main.go


FROM alpine
WORKDIR /
COPY --from=builder blog-api blog-api
RUN apk --no-cache add tzdata


CMD ["./blog-api"]
