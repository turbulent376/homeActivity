module git.jetbrains.space/orbi/fcsd/timesheet

go 1.16

require (
	git.jetbrains.space/orbi/fcsd/kit v0.0.1-20211118-0920
	git.jetbrains.space/orbi/fcsd/proto v0.0.1-20211122-0940
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
)

replace git.jetbrains.space/orbi/fcsd/kit v0.0.1-20211118-0920 => git.jetbrains.space/orbi/fcsd/kit.git v0.0.1-20211118-0920

replace git.jetbrains.space/orbi/fcsd/proto v0.0.1-20211122-0940 => git.jetbrains.space/orbi/fcsd/proto.git v0.0.1-20211122-0940
