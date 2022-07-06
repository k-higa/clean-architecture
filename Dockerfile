FROM golang:1.18.3 as dev


WORKDIR /go/src/app
#RUN #apt-updte && apt-get install
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 1326

CMD ["go", "run", "main.go"]
