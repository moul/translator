package translator_endpoints



import (
	"fmt"
	context "golang.org/x/net/context"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/translator/gen/pb"
	"github.com/go-kit/kit/endpoint"
)

var _ = fmt.Errorf

type Endpoints struct {
	
	TranslateEndpoint endpoint.Endpoint
	
}


func (e *Endpoints)Translate(ctx context.Context, in *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	out, err := e.TranslateEndpoint(ctx, in)
	if err != nil {
		return &pb.TranslateResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.TranslateResponse), err
}



func MakeTranslateEndpoint(svc pb.TranslatorServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.TranslateRequest)
		rep, err := svc.Translate(ctx, req)
		if err != nil {
			return &pb.TranslateResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}


func MakeEndpoints(svc pb.TranslatorServiceServer) Endpoints {
	return Endpoints{
		
		TranslateEndpoint: MakeTranslateEndpoint(svc),
		
	}
}
