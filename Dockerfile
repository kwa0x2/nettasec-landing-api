FROM golang:1.21.5

WORKDIR /go/src/nettasec.co.uk/landing/backend

COPY . .

RUN go get -v

RUN go build -o main .

CMD [ "./main" ]