# translator

[![GuardRails badge](https://badges.production.guardrails.io/moul/translator.svg)](https://www.guardrails.io)

:microphone: Translator micro-service

A translation micro-service using [Gettext](https://www.gnu.org/software/gettext/), [Protobuf](https://github.com/google/protobuf) and [Go-Kit](https://github.com/go-kit/kit), with more that 75% boilerplate code generated automatically using [protoc-gen-gotemplate](https://github.com/moul/protoc-gen-gotemplate).

## Code generation

```console
# Custom code
$ wc -l service/service.go pb/translator.proto cmd/translator/main.go
      23 service/service.go
      19 pb/translator.proto
      68 cmd/translator/main.go
      110 total
# Generated code
$ wc -l $(find gen -name "*.go")
      50 gen/endpoints/endpoints.go
     187 gen/pb/translator.pb.go
      59 gen/transports/grpc/grpc.go
      47 gen/transports/http/http.go
     343 total
```

## Usage

```console
$ curl localhost:8000/Translate -XPOST -d'{"message":"Hello world.","language":"fr_FR"}'
{"message":"Bonjour monde."}
$ curl localhost:8000/Translate -XPOST -d'{"message":"Hello world.","language":"en_EN"}'
{"message":"Hello world."}
$ curl localhost:8000/Translate -XPOST -d'{"message":"Hello world.","language":"foo"}'
{"message":"Hello world."}
```
