module github.com/aserto-dev/go-grpc

go 1.19

// replace github.com/aserto-dev/mage-loot => ../mage-loot

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	github.com/magefile/mage v1.15.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230706204954-ccb25ca9f130
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230706204954-ccb25ca9f130
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230711160842-782d3b101e98 // indirect
)
