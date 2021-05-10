FROM golang:alpine
RUN apk update && \
    apk upgrade && \
    apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gin-gonic/gin
RUN go get cloud.google.com/go/firestore
RUN go get cloud.google.com/go/storage
RUN go get firebase.google.com/go
RUN go get github.com/joho/godotenv
RUN go get github.com/kr/pretty
RUN go get golang.org/x/crypto
RUN go get golang.org/x/oauth2
RUN go get google.golang.org/api
EXPOSE 3000
CMD go run main.go . -DFORGROUND