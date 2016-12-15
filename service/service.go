package translatorsvc

import (
	"fmt"

	"github.com/moul/translator/gen/pb"
	"golang.org/x/net/context"
)

type Service struct{}

func New() translatorpb.TranslatorServiceServer {
	return &Service{}
}

func (svc Service) Translate(ctx context.Context, input *translatorpb.TranslateRequest) (*translatorpb.TranslateResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}
