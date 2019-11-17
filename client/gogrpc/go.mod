module github.com/netflix/conductor/client/gogrpc

require (
	github.com/stretchr/testify v1.2.1
	google.golang.org/grpc v1.15.0
)

go 1.13

replace github.com/netflix/conductor/client/gogrpc => github.com/mactaggart/conductor/client/gogrpc grpc-mods