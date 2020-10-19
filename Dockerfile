FROM golang:1.14

WORKDIR /go/src/go-osmand-tracker

COPY ["go.mod", "go.sum", "main.go", "./"]
COPY ["internal/auxillary", "./internal/auxillary/"]
COPY ["internal/database", "./internal/database/"]
COPY ["internal/server", "./internal/server/"]
COPY ["internal/types", "./internal/types/"]

RUN ["go", "mod", "download"]

RUN ["go", "build", "-o", "main", "."]

CMD ["/go/src/go-osmand-tracker/main"]
