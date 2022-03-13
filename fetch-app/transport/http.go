package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/khakim88/sample-auth/fetch-app/common"
	"github.com/khakim88/sample-auth/fetch-app/endpoint"

	"github.com/khakim88/sample-auth/fetch-app/transport/decode"
	"github.com/khakim88/sample-auth/fetch-app/transport/encode"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(r *mux.Router) http.Handler {
	//start database conn
	// repo, err := postgres.NewPostgresConn(repository.DBConfiguration{
	// 	DBHost:            config.Get(constant.DBHostKey, "localhost"),
	// 	DBPort:            config.Get(constant.DBPortKey, "5432"),
	// 	DBUser:            config.Get(constant.DBUserKey, "postgres"),
	// 	DBPassword:        config.Get(constant.DBPasswordKey, "P@ssw0rd123"),
	// 	DBName:            config.Get(constant.DBNameKey, "promo"),
	// 	MaxConnection:     10,
	// 	MaxIdleConnection: 10,
	// })

	// if err != nil {
	// 	logger.Error(err)
	// 	os.Exit(1)
	// }

	//svc := service.NewPromoService(repo)

	//endpoint := endpoint.MakeEndpoints(&svc)

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(common.EncodeLegacyError),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	}

	getResourceEndpoint := endpoint.MakeGetResource
	getResourceHandler := kithttp.NewServer(
		getResourceEndpoint,
		decode.DecodeGetResourceRequest,
		encode.EncodeResponseWithData,
		opts...,
	)

	r.Handle("/resource/", getResourceHandler).Methods("POST")

	return r
}

//encode Response
func EncodeResponseHttp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
