module github.com/0B1t322/zero-validation/grpc-example

go 1.24.1

replace github.com/0B1t322/zero-validaton => ../..

require (
	github.com/0B1t322/zero-validaton v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.5
)

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
