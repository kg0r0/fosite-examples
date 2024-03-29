package authorizationserver

import (
	"log"
	"net/http"
)

func IntrospectionEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	mySessionData := newSession("")
	ir, err := oauth2.NewIntrospectionRequest(ctx, req, mySessionData)
	if err != nil {
		log.Printf("Error occurred in NewIntrospectionRequest: %+v", err)
		oauth2.WriteIntrospectionError(ctx, rw, err)
		return
	}

	oauth2.WriteIntrospectionResponse(ctx, rw, ir)
}
