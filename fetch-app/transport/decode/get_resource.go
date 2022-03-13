package decode

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/khakim88/sample-auth/fetch-app/common"
	"github.com/khakim88/sample-auth/fetch-app/model"
	"github.com/rs/zerolog/log"
)

var (
	errParseBody = common.NewLegacyError("error parse body payload")
)

func DecodeGetResourceRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().
			Str("method", "payload").
			Msg(err.Error())
		return nil, common.NewLegacyError(err.Error())
	}

	validationRequest := new(model.ValidatePromotionRequest)
	err = json.Unmarshal(body, &validationRequest)
	if err != nil {
		log.Error().
			Str("method", "payload").
			Msg(errParseBody.Error())
		return nil, errParseBody
	}
	return *validationRequest, nil
}
