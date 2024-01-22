package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/kg0r0/fosite-examples/authorizationserver"
	client "github.com/kg0r0/fosite-examples/client"
	"github.com/kg0r0/fosite-examples/resourceserver"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var clientConf = oauth2.Config{
	ClientID:     "my-client",
	ClientSecret: "foobar",
	RedirectURL:  "http://localhost:3846/callback",
	Scopes:       []string{"photos", "openid", "offline"},
	Endpoint: oauth2.Endpoint{
		TokenURL: "http://localhost:3846/oauth2/token",
		AuthURL:  "http://localhost:3846/oauth2/auth",
	},
}

var appClientConf = clientcredentials.Config{
	ClientID:     "my-client",
	ClientSecret: "foobar",
	Scopes:       []string{"fosite"},
	TokenURL:     "http://localhost:3846/oauth2/token",
}

var appClientConfRotated = clientcredentials.Config{
	ClientID:     "my-client",
	ClientSecret: "foobaz",
	Scopes:       []string{"fosite"},
	TokenURL:     "http://localhost:3846/oauth2/token",
}

func main() {
	// authorization server
	http.HandleFunc("/oauth2/auth", authorizationserver.AuthEndpoint)
	http.HandleFunc("/oauth2/token", authorizationserver.TokenEndpoint)

	// revoke tokens
	http.HandleFunc("/oauth2/revoke", authorizationserver.RevokeEndpoint)
	http.HandleFunc("/oauth2/introspect", authorizationserver.IntrospectionEndpoint)

	// client
	http.HandleFunc("/", client.IndexHandler(clientConf))

	http.HandleFunc("/client", client.ClientEndpoint(appClientConf))
	http.HandleFunc("/client-new", client.ClientEndpoint(appClientConfRotated))
	http.HandleFunc("/owner", client.OwnerHandler(clientConf))
	http.HandleFunc("/callback", client.CallbackHandler(clientConf))

	// resource server
	http.HandleFunc("/protected", resourceserver.ProtectedEndpoint(appClientConf))

	port := "3846"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("Please open your webbrowser at http://localhost:" + port)
	_ = exec.Command("open", "http://localhost:"+port).Run()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
