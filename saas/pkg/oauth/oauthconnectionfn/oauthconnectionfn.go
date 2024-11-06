package oauthconnectionfn

import (
	"context"
	"errors"
	"lace/util"
	"saas/gen/ent"
	"saas/gen/ent/oauthconnection"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

type OauthRequestCache struct {
	AppID             string                 `json:"appId"`
	RedirectURL       string                 `json:"redirectUrl"`
	OauthConnectionID string                 `json:"oauthConnectionId"`
	Metadata          map[string]interface{} `json:"metadata"`
	// Provider          string                 `json:"provider"` // google | github
}

type OauthResponseCache struct {
	Code              string        `json:"code"`              // retruned in callback url
	OauthConnectionID string        `json:"oauthConnectionId"` // retruned in callback url as state
	Token             *oauth2.Token `json:"token"`             // genertated token from code
	Provider          string        `json:"provider"`          // google | github
}

type OauthRequestParams struct {
	// Provider    string `json:"provider"`
	RedirectUri string `json:"redirectUri"`
}

const (
	ProviderGoogle = "google"
	// ProviderTwitter   = "twitter"
	ProviderGithub   = "github"
	ProviderLinkedin = "linkedin"
	// ProviderMicrosoft = "microsoft"
	ProviderFacebook = "facebook"
)

func NewOauthRequester(appID, oatconnId string, client *ent.Client) (*OauthRequester, error) {
	oauthConnection, err := client.OauthConnection.Query().Where(oauthconnection.AppID(appID), oauthconnection.ID(oatconnId), oauthconnection.Status(true)).First(context.Background())

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
	case ProviderGithub:
		return github.Endpoint
	case ProviderLinkedin:
		return linkedin.Endpoint
	// case ProviderMicrosoft:
	// 	return microsoft.Endpoint
	case ProviderFacebook:
		return facebook.Endpoint
		// case ProviderTwitter:
		// 	return oauth2.Endpoint{
		// 		AuthURL:  "https://twitter.com/i/oauth2/authorize",
		// 		TokenURL: "https://api.twitter.com/oauth2/token",
		// 		// DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		// 		AuthStyle: oauth2.AuthStyleInParams,
		// 	}
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

func OauthTokenFromResponseCache(responseCache OauthResponseCache) (*oauth2.Token, error) {
	if responseCache.Token == nil {
		return nil, errors.New("token is missing in response cache")
	}

	// some oauth2 does not have referesh token e.g. github
	// if len(responseCache.Token.RefreshToken) == 0 {
	// 	return nil, errors.New("refresh token is missing in response cache")
	// }

	var oauth2Token *oauth2.Token

	err := util.InterfaceToStruct(responseCache.Token, &oauth2Token)
	if err != nil {
		return nil, err
	}

	return oauth2Token, nil
}

func OauthToken(accessToken, refreshToken string, expiry time.Time, tokenType string) *oauth2.Token {
	// if time is zero or not passed then expire the token so it can be regnerated otherwise it will
	// keep returning the same token
	if expiry.IsZero() {
		expiry = time.Now().Add(time.Duration(-24) * time.Hour)
	}

	oauth2Token := &oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    tokenType,
		RefreshToken: refreshToken,
		Expiry:       expiry,
	}
	return oauth2Token
}
