module github.com/turbulent376/homeactivity/activity

go 1.16

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/turbulent376/kit v0.0.2
	github.com/turbulent376/proto v0.0.3
	github.com/vektra/mockery/v2 v2.9.4 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
)

//replace github.com/turbulent376/homeactivity/kit v0.0.1-20211118-0920 => github.com/turbulent376/homeactivity/kit.git v0.0.1-20211118-0920

///replace github.com/turbulent376/homeactivity/proto v0.0.1-20211122-0940 => github.com/turbulent376/homeactivity/proto.git v0.0.1-20211122-0940
