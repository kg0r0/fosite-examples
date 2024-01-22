package authorizationserver

import (
	"net/http"
)

func RevokeEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	err := oauth2.NewRevocationRequest(ctx, req)

	oauth2.WriteRevocationResponse(ctx, rw, err)
}
