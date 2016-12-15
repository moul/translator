.PHONY: build
build: translator

translator: pb/translator.pb.go gen/.generated

pb/translator.pb.go: pb/translator.proto
	cd pb; protoc --gogo_out=plugins=grpc:. ./translator.proto

gen/.generated:	pb/translator.proto
	@rm -rf gen
	@mkdir -p gen
	cd pb; protoc --gotemplate_out=template_dir=../vendor/github.com/moul/protoc-gen-gotemplate/examples/go-kit/templates/{{.File.Package}}/gen:../gen ./translator.proto
	@touch gen/.generated
