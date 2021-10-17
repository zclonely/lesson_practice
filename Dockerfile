FROM golang:1.16-alpine as build
LABEL user="zchao"
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /lesson_practice
COPY ./lesson_practice/go.mod .
COPY ./lesson_practice/go.sum .
RUN mkdir /lesson_practice/lesson_4
COPY ./lesson_practice/lesson_4/ /lesson_practice/lesson_4/
RUN go mod download
EXPOSE 80
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./lesson_4/main.go
CMD ["/lesson_practice/main"]
