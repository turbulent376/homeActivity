module git.jetbrains.space/orbi/fcsd/api

go 1.16

require (
	git.jetbrains.space/orbi/fcsd/kit v0.0.1-20211202-1930
	git.jetbrains.space/orbi/fcsd/proto v0.0.1-20211122-0940
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gorilla/mux v1.8.0
	github.com/swaggo/http-swagger v1.1.1
	github.com/swaggo/swag v1.7.1
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
)

replace git.jetbrains.space/orbi/fcsd/kit v0.0.1-20211202-1930 => git.jetbrains.space/orbi/fcsd/kit.git v0.0.1-20211202-1930

replace git.jetbrains.space/orbi/fcsd/proto v0.0.1-20211122-0940 => git.jetbrains.space/orbi/fcsd/proto.git v0.0.1-20211122-0940
