module github.com/aserto-dev/go-grpc

go 1.21

// replace github.com/aserto-dev/mage-loot => ../mage-loot

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/magefile/mage v1.15.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240506185236-b8a5c65736ae
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240506185236-b8a5c65736ae
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.34.1
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
)
