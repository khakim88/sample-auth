package encode

import (
	"context"
	"encoding/json"
	"net/http"
)

var (
	CONTENTTYPEJSONANDCHARSETUTF8VALUE = "application/json; charset=utf-8"
	CONTENTTYPEKEY                     = "Content-Type"
)

type dataResponse struct {
	Data interface{} `json:"data"`
}

func EncodeResponseWithData(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(CONTENTTYPEKEY, CONTENTTYPEJSONANDCHARSETUTF8VALUE)

	resp := dataResponse{Data: response}
	return json.NewEncoder(w).Encode(resp)
}
