FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN go build -o bin/blog-api cmd/main.go


FROM alpine
COPY --from=build /app/bin/blog-api /app/blog-api
WORKDIR /app
RUN apk --no-cache add tzdata


CMD ["./blog-api"]
