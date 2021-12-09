module git.jetbrains.space/orbi/fcsd/auth

go 1.16

require (
	firebase.google.com/go v3.13.0+incompatible
	git.jetbrains.space/orbi/fcsd/kit v0.0.1-20211027-0915
	git.jetbrains.space/orbi/fcsd/proto v0.0.1-20211122-0940
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/lib/pq v1.10.3 // indirect
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
)

replace git.jetbrains.space/orbi/fcsd/kit v0.0.1-20211118-0920 => git.jetbrains.space/orbi/fcsd/kit.git v0.0.1-20211118-0920

replace git.jetbrains.space/orbi/fcsd/proto v0.0.1-20211122-0940 => git.jetbrains.space/orbi/fcsd/proto.git v0.0.1-20211122-0940
