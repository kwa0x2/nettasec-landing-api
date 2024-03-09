FROM golang:1.21.5

WORKDIR /go/src/nettasec.co.uk/landing/backend

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

CMD [ "./main" ]
