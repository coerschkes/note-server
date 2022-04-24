FROM golang:1.18

RUN mkdir server 
WORKDIR /server

EXPOSE 8080

COPY go.mod go.sum ./
RUN go mod download 
RUN go mod verify
COPY ./src/ /server/src/
COPY ./resources/ /server/resources/
RUN go build -v -o ./src/note-server ./src
WORKDIR /server/src

CMD ["./note-server"]