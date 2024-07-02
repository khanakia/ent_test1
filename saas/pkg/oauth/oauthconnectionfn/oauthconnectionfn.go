package oauthconnectionfn

import (
	"context"
	"saas/gen/ent"
	"saas/gen/ent/oauthconnection"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthRequestCache struct {
	RedirectURL       string                 `json:"redirectUrl"`
	OauthConnectionID string                 `json:"oauthConnectionId"`
	Metadata          map[string]interface{} `json:"metadata"`
}

type OauthResponseCache struct {
	Code              string        `json:"code"`              // retruned in callback url
	OauthConnectionID string        `json:"oauthConnectionId"` // retruned in callback url as state
	Token             *oauth2.Token `json:"token"`             // genertated token from code
}

type OauthRequestParams struct {
	// Provider    string `json:"provider"`
	RedirectUri string `json:"redirectUri"`
}

const (
	ProviderGoogle = "google"
)

func NewOauthRequester(oatconnId string, client *ent.Client) (*OauthRequester, error) {
	oauthConnection, err := client.OauthConnection.Query().Where(oauthconnection.ID(oatconnId), oauthconnection.Status(true)).First(context.Background())

	if err != nil {
		return nil, err
	}

	return &OauthRequester{
		oauthConnection: oauthConnection,
		redirectURL:     oauthConnection.RedirectURL, // default redirectUrl
		state:           oauthConnection.ID,          // default state
	}, nil
}

type OauthRequester struct {
	oauthConnection *ent.OauthConnection
	redirectURL     string
	state           string
}

func (cls OauthRequester) GetAuthUrl() string {
	oauthConfGl := cls.OauthConfig()
	url := oauthConfGl.AuthCodeURL(cls.state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return url
}
func (cls OauthRequester) TokenFromCode(code string) (*oauth2.Token, error) {
	oauthConfGl := cls.OauthConfig()
	token, err := oauthConfGl.Exchange(context.Background(), code)
	return token, err
	// logger.Info("TOKEN>> AccessToken>> " + token.AccessToken)
	// logger.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
	// logger.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)
}

func (cls *OauthRequester) GetOauthConnection() *ent.OauthConnection {
	return cls.oauthConnection
}

func (cls *OauthRequester) SetState(state string) {
	cls.state = state
}

func (cls *OauthRequester) SetRedirectURL(url string) {
	cls.redirectURL = url
}

func (cls *OauthRequester) GetRedirectURL() string {
	return cls.redirectURL
}

// we assume there will never be error as we will validate the provider while inserting into the database
func (cls OauthRequester) Endpoint() oauth2.Endpoint {
	switch cls.oauthConnection.Provider {
	case ProviderGoogle:
		return google.Endpoint
	}

	return oauth2.Endpoint{}
}

func (cls OauthRequester) OauthConfig() *oauth2.Config {
	scopes := strings.Split(cls.oauthConnection.Scopes, ",")
	config := &oauth2.Config{
		ClientID:     cls.oauthConnection.ClientID,
		ClientSecret: cls.oauthConnection.ClientSecret,
		RedirectURL:  cls.redirectURL,
		Scopes:       scopes,
		Endpoint:     cls.Endpoint(),
	}
	return config
}
