DOCKER_IMAGE ?=	moul/translator
SOURCES := cmd/translator/main.go service/service.go
MO_FILES := $(shell find locales -name "*.po" | sed 's/\.po/\.mo/g')

.PHONY: build
build: translator

translator: gen/pb/translator.pb.go gen/.generated $(SOURCES) $(MO_FILES)
	go build -o translator ./cmd/translator

gen/pb/translator.pb.go: pb/translator.proto
	@mkdir -p gen/pb
	cd pb; protoc --gogo_out=plugins=grpc:../gen/pb ./translator.proto

gen/.generated:	pb/translator.proto
	@mkdir -p gen
	cd pb; protoc --gotemplate_out=destination_dir=../gen,template_dir=../vendor/github.com/moul/protoc-gen-gotemplate/examples/go-kit/templates/{{.File.Package}}/gen:../gen ./translator.proto
	@touch gen/.generated

%.mo: %.po
	msgfmt -o $@ $<

.PHONY: install
install:
	go install ./cmd/translator

.PHONY: docker.build
docker.build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker.run
docker.run:
	docker run -p 8000:8000 -p 9000:9000 $(DOCKER_IMAGE)
