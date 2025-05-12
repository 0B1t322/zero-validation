module github.com/0B1t322/zero-validation/grpc-example

go 1.23.0

replace github.com/0B1t322/zero-validation => ../..

require (
	github.com/0B1t322/zero-validation v0.0.2-rc
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.5
)

require (
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)
