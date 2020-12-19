FROM golang:1.14

WORKDIR /go/src/waarzitjenu/server

COPY ["go.mod", "go.sum", "main.go", "./"]
COPY ["auth", "./auth"]
COPY ["database", "./database"]
COPY ["engine", "./engine"]
COPY ["filesystem", "./filesystem"]
COPY ["settings", "./settings"]
COPY ["types", "./types"]


RUN ["go", "mod", "download"]

RUN ["go", "build", "-o", "main", "."]

CMD ["/go/src/waarzitjenu/server/main"]