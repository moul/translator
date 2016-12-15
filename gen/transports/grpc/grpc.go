package translator_grpctransport



import (
        "fmt"

	context "golang.org/x/net/context"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/translator/gen/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/translator/gen/endpoints"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.TranslatorServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{
		
                
                
		translate: grpctransport.NewServer(
			ctx,
			endpoints.TranslateEndpoint,
			decodeTranslateRequest,
			encodeTranslateResponse,
			options...,
		),
                
		
                
	}
}

type grpcServer struct {
	
	translate grpctransport.Handler
	
}


func (s *grpcServer) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	_, rep, err := s.translate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TranslateResponse), nil
}

func decodeTranslateRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeTranslateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.TranslateResponse)
	return resp, nil
}

