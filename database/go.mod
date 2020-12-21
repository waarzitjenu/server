module github.com/waarzitjenu/server/database

go 1.14

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/stretchr/testify v1.6.1
	go.etcd.io/bbolt v1.3.5 // indirect
)

replace (
	github.com/waarzitjenu/server/filesystem => ../filesystem
	github.com/waarzitjenu/server/types => ../types
)