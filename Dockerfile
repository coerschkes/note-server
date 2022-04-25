FROM golang:1.18

RUN mkdir server 
WORKDIR /server

EXPOSE 8080

COPY go.mod go.sum ./
RUN go mod download 
RUN go mod verify
COPY ./src/ /server/src/
RUN go build -v -o ./src/main/note-server ./src/main
WORKDIR /server/src/main/

CMD ["./note-server"]