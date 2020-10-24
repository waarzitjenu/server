FROM golang:1.14

WORKDIR /go/src/go-osmand-tracker

COPY ["go.mod", "go.sum", "main.go", "./"]
COPY ["internal/auxiliary", "./internal/auxiliary/"]
COPY ["internal/database", "./internal/database/"]
COPY ["internal/filesystem", "./internal/filesystem/"]
COPY ["internal/server", "./internal/server/"]
COPY ["internal/settings", "./internal/settings/"]
COPY ["internal/types", "./internal/types/"]

RUN ["go", "mod", "download"]

RUN ["go", "build", "-o", "main", "."]

CMD ["/go/src/go-osmand-tracker/main"]
