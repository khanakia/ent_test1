package oauthuserinfo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"saas/pkg/oauth/facebookapi"
	"saas/pkg/oauth/githubapi"
	"saas/pkg/oauth/googleauthfn"
	"saas/pkg/oauth/oauthconnectionfn"

	"golang.org/x/oauth2"
)

// UserInfo represents the OpenID Connect userinfo claims.
type UserInfo struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type userInfoRaw struct {
	Subject string `json:"sub"`
	Profile string `json:"profile"`
	Email   string `json:"email"`
	// Handle providers that return email_verified as a string
	// https://forums.aws.amazon.com/thread.jspa?messageID=949441&#949441 and
	// https://discuss.elastic.co/t/openid-error-after-authenticating-against-aws-cognito/206018/11
	EmailVerified stringAsBool `json:"email_verified"`
}

func UserInfoGet(provider string, token *oauth2.Token) (*UserInfo, error) {
	switch provider {
	case oauthconnectionfn.ProviderGoogle:
		userinfo, err := googleauthfn.GetUserInfo(token.AccessToken)
		if err != nil {
			return nil, err
		}
		return &UserInfo{
			Email:     userinfo.Email,
			FirstName: userinfo.GivenName,
			LastName:  userinfo.FamilyName,
		}, nil

	case oauthconnectionfn.ProviderGithub:
		userinfo, err := githubapi.GetUserInfo(token.AccessToken)
		if err != nil {
			return nil, err
		}
		return &UserInfo{
			Email:     userinfo.Email,
			FirstName: userinfo.FirstName,
			LastName:  userinfo.LastName,
		}, nil

	case oauthconnectionfn.ProviderFacebook:
		userinfo, err := facebookapi.GetUserInfo(token.AccessToken)
		if err != nil {
			return nil, err
		}
		return &UserInfo{
			Email:     userinfo.Email,
			FirstName: userinfo.FirstName,
			LastName:  userinfo.LastName,
		}, nil
	default:
		return nil, fmt.Errorf("oauthuserinfo: userinfo endpoint not supported for provider %q", provider)
	}
}

func getUserinfoURL(provider string) (string, error) {
	switch provider {
	case "google":
		return "https://www.googleapis.com/oauth2/v3/userinfo", nil
	default:
		return "", fmt.Errorf("oauthuserinfo: userinfo endpoint not supported for provider %q", provider)
	}
}

func UserInfoGet1(ctx context.Context, tokenSource oauth2.TokenSource) (*UserInfo, error) {
	userInfoURL, err := getUserinfoURL("google")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("oauthuserinfo: create GET request: %v", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("oauthuserinfo: get access token: %v", err)
	}
	token.SetAuthHeader(req)

	resp, err := doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var userInfo userInfoRaw
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("oauthuserinfo: failed to decode userinfo: %v", err)
	}
	return &UserInfo{
		// Subject:       userInfo.Subject,
		// Profile:       userInfo.Profile,
		Email: userInfo.Email,
		// EmailVerified: bool(userInfo.EmailVerified),
		// claims:        body,
	}, nil
}

func getClient(ctx context.Context) *http.Client {
	if c, ok := ctx.Value(oauth2.HTTPClient).(*http.Client); ok {
		return c
	}
	return nil
}

func doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := http.DefaultClient
	if c := getClient(ctx); c != nil {
		client = c
	}
	return client.Do(req.WithContext(ctx))
}

type stringAsBool bool

func (sb *stringAsBool) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case "true", `"true"`:
		*sb = true
	case "false", `"false"`:
		*sb = false
	default:
		return errors.New("invalid value for boolean")
	}
	return nil
}
