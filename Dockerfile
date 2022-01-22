FROM golang:1.16.13

WORKDIR /go/src/waarzitjenu/server

COPY ["go.mod", "go.sum", "main.go", "./"]
COPY ["internal/auth", "./internal/auth"]
COPY ["internal/database", "./internal/database"]
COPY ["internal/engine", "./internal/engine"]
COPY ["internal/filesystem", "./internal/filesystem"]
COPY ["internal/settings", "./internal/settings"]
COPY ["internal/types", "./internal/types"]


RUN ["go", "mod", "download"]
RUN ["go", "mod", "verify"]
RUN ["go", "vet"]

RUN ["go", "build", "-o", "main", "."]

CMD ["/go/src/waarzitjenu/server/main"]
