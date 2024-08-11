package githubapi

import (
	"lace/httpreq"

	"github.com/ubgo/goutil"
)

type UserInfo struct {
	Name      string `json:"name"` // Jane Doe
	Company   string `json:"company"`
	AvatarURL string `json:"avatarUrl"` // https://avatars.githubusercontent.com/u/2174170?v=4

	Email string `json:"email"`

	// extract
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetUserInfo(accessToken string) (*UserInfo, error) {
	hreq := httpreq.RequestParam{
		EndPoint: "https://api.github.com/user",
		Method:   "GET",
		Token:    accessToken,
	}

	var res UserInfo

	_, err := httpreq.Do(hreq, &res)

	if err != nil {
		return nil, err
	}

	// goutil.PrintToJSON(res)

	name := goutil.ParseName(res.Name)
	res.FirstName = name.FirstName
	res.LastName = name.LastName

	// github does not returns the primary email so run a separate api to get all the emails
	// associted with account and then extract the primary email
	userEmails, err := GetUserEmails(accessToken)
	if err != nil {
		return nil, err
	}

	for _, userEmail := range userEmails {
		if userEmail.Primary {
			res.Email = userEmail.Email
			break
		}
	}

	return &res, nil
}

type UserEmail struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

func GetUserEmails(accessToken string) ([]UserEmail, error) {
	hreq := httpreq.RequestParam{
		EndPoint: "https://api.github.com/user/emails",
		Method:   "GET",
		Token:    accessToken,
	}

	var res []UserEmail

	_, err := httpreq.Do(hreq, &res)

	if err != nil {
		return nil, err
	}

	// goutil.PrintToJSON(res)

	return res, nil
}
