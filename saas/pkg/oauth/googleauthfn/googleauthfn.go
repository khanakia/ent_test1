package googleauthfn

import (
	"lace/httpreq"

	"github.com/ubgo/goutil"
)

type UserInfo struct {
	Sub           string `json:"sub"`         // 103936245197801500536
	Name          string `json:"name"`        // Jane Doe
	GivenName     string `json:"given_name"`  // Jane
	FamilyName    string `json:"family_name"` // Doe
	Picture       string `json:"picture"`     // url
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"` // en
}

func GetUserInfo(accessToken string) (*UserInfo, error) {
	hreq := httpreq.RequestParam{
		EndPoint: "https://www.googleapis.com/oauth2/v3/userinfo",
		Method:   "GET",
		QueryParams: map[string]string{
			"access_token": accessToken,
		},
	}

	var res UserInfo

	_, err := httpreq.Do(hreq, &res)

	if err != nil {
		return nil, err
	}

	goutil.PrintToJSON(res)

	return &res, nil
}
