package translatorsvc

import (
	"github.com/chai2010/gettext-go/gettext"
	"golang.org/x/net/context"

	"github.com/moul/translator/gen/pb"
)

type Service struct{}

func New() translatorpb.TranslatorServiceServer {
	gettext.BindTextdomain("translator", "locales", nil)
	gettext.Textdomain("translator")
	return &Service{}
}

func (svc Service) Translate(ctx context.Context, input *translatorpb.TranslateRequest) (*translatorpb.TranslateResponse, error) {
	gettext.SetLocale(input.Language)
	return &translatorpb.TranslateResponse{
		Message: gettext.PGettext("", input.Message),
	}, nil
}
