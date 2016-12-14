pb/translator.pb.go: translator.proto
	protoc --gogo_out=plugins=grpc:pb ./translator.proto
