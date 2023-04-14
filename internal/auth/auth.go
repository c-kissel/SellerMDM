package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/c-kissel/SellerMDM.git/specs"
	oidc "github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type authService struct {
	keycloak *gocloak.GoCloak
	realm    string

	oidcProvider    oidc.Provider
	oidcConfig      oidc.Config
	oauth2Config    oauth2.Config
	idTokenVerifier oidc.IDTokenVerifier
	token           oauth2.Token
}

func NewAuthMiddleware() *authService {

	oidcProvider := *createOidcProvider(context.Background())
	oidcConfig, oauth2Config := createConfig(oidcProvider)
	idTokenVerifier := *oidcProvider.Verifier(&oidcConfig)

	client := gocloak.NewClient(os.Getenv("KEYCLOAK_SERVER"))
	realm := os.Getenv("KEYCLOAK_REALM")

	return &authService{
		keycloak: client,
		realm:    realm,

		oidcProvider:    oidcProvider,
		oidcConfig:      oidcConfig,
		oauth2Config:    oauth2Config,
		idTokenVerifier: idTokenVerifier,
	}
}

func Authorization() func(next http.Handler) http.Handler {
	return NewAuthMiddleware().Authorize
}

type roleClaims struct {
	jwt.Claims
	Roles []string `json:"allowed-origins"`
}

// Authorization middleware
func (a *authService) Authorize(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		unauthorized := func(message string, data ...interface{}) {
			logrus.Error("Authorization... FAILED")
			logrus.Debugf(message, data...)
			w.WriteHeader(http.StatusUnauthorized)
		}

		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		logrus.Debug("In Authorize() middleware: performing authorization")
		ctx := r.Context()
		requiredScopes := ctx.Value(specs.KeycloakScopes)
		if requiredScopes != nil {

			logrus.Debugf("In Authorize() middleware: authorization required: %s", requiredScopes)

			reqToken := r.Header.Get("Authorization")
			if len(reqToken) == 0 {
				unauthorized("Authorization token is empty, check your API call.")
				return
			}
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				unauthorized("Incorrect authorization token: '%v'", reqToken)
				return
			}

			accessToken := splitToken[1]
			rptResult, err := a.keycloak.RetrospectToken(ctx, accessToken, a.oauth2Config.ClientID, a.oauth2Config.ClientSecret, os.Getenv("KEYCLOAK_REALM"))
			if err != nil {
				unauthorized("error introspecting token: %s", err.Error())
				return
			}

			if !*rptResult.Active {
				unauthorized("token is not active: %s", err.Error())
				return
			}

			_, mapClaims, err := a.keycloak.DecodeAccessToken(ctx, accessToken, a.realm)
			if err != nil {
				unauthorized("token is not valid: %s", err.Error())
				return
			}

			tokenRoles := make([]string, 0)

			realmAccessClaim, ok := (*mapClaims)["realm_access"].(map[string]interface{})
			if ok {
				rolesClaim, ok := realmAccessClaim["roles"].([]interface{})
				if ok {
					// rolesClaim is a []interface{} containing the role names
					for _, role := range rolesClaim {
						tokenRoles = append(tokenRoles, role.(string))
					}
				}
			}

			endpointScopes, ok := requiredScopes.([]string)
			if ok {
				for _, requiredScope := range endpointScopes {

					isFound := false
					for _, userRole := range tokenRoles {
						if userRole == requiredScope {
							isFound = true
						}
					}

					if !isFound {
						unauthorized("insufficient claims has '%v' want '%v'", tokenRoles, endpointScopes)
						return
					}
				}
			} else {
				unauthorized("failed to convert required scopes to []string: %v", requiredScopes)
				return
			}

		}
		logrus.Debug("Authorization... PASSED")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Returns OAuth config by name
func (srv *authService) GetOAuthConfig(name string) *oauth2.Config {
	return &srv.oauth2Config
}

// Extracts user data from OAuth data
func (srv *authService) ExtractUserData() error {
	return nil
}

func createOidcProvider(ctx context.Context) *oidc.Provider {
	url := fmt.Sprintf("%s/realms/%s", os.Getenv("KEYCLOAK_SERVER"), os.Getenv("KEYCLOAK_REALM"))
	provider, err := oidc.NewProvider(ctx, url)

	if err != nil {
		logrus.Fatal("Failed to fetch discovery document: ", err)
	}

	return provider
}

func createConfig(provider oidc.Provider) (oidc.Config, oauth2.Config) {
	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("KEYCLOAK_CLIENT_ID"),
	}

	config := oauth2.Config{
		ClientID:     oidcConfig.ClientID,
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		Endpoint:     provider.Endpoint(),
		RedirectURL:  os.Getenv("KEYCLOAK_REDIRECT"),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "orders"},
	}

	return *oidcConfig, config
}
