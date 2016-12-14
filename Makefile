pb/translator.pb.go: pb/translator.proto
	cd pb; protoc --gogo_out=plugins=grpc:. ./translator.proto
