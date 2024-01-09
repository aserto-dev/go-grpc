module github.com/aserto-dev/go-grpc

go 1.20

// replace github.com/aserto-dev/mage-loot => ../mage-loot

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	github.com/magefile/mage v1.15.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240108191215-35c7eff3a6b1
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240108191215-35c7eff3a6b1
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240108191215-35c7eff3a6b1 // indirect
)
