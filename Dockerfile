FROM golang:1.18

RUN mkdir server 
WORKDIR /server

EXPOSE 8080

COPY ./src/ /server/src/
COPY go.mod go.sum ./
RUN go mod download 
RUN go mod verify
RUN go build -v -o note-server ./src
RUN ls -la ./src

CMD ["./note-server"]