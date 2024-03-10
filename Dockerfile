FROM golang:1.21.5

WORKDIR /go/src/nettasec.co.uk/landing/backend

COPY myndfin.crt /etc/ssl/certs/myndfin.crt
COPY myndfin.key /etc/ssl/private/myndfin.key

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

CMD [ "./main" ]
