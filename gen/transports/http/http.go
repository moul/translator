package translator_httptransport



import (
       "log"
	"net/http"
	"encoding/json"
	context "golang.org/x/net/context"

	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/translator/gen/pb"
        gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/translator/gen/endpoints"
)


func MakeTranslateHandler(ctx context.Context, svc pb.TranslatorServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeTranslateRequest,
		encodeTranslateResponse,
                []httptransport.ServerOption{}...,
	)
}

func decodeTranslateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.TranslateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeTranslateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


func RegisterHandlers(ctx context.Context, svc pb.TranslatorServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
        log.Println("new HTTP endpoint: \"/Translate\" (service=Translator)")
	mux.Handle("/Translate", MakeTranslateHandler(ctx, svc, endpoints.TranslateEndpoint))
	
	return nil
}
