package facebookapi

import (
	"lace/httpreq"
)

type UserInfo struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Picture    Picture `json:"picture"`
	LastName   string  `json:"last_name"`
	FirstName  string  `json:"first_name"`
	NameFormat string  `json:"name_format"`
	ID         string  `json:"id"`
}
type Data struct {
	Height       int    `json:"height"`
	IsSilhouette bool   `json:"is_silhouette"`
	URL          string `json:"url"`
	Width        int    `json:"width"`
}
type Picture struct {
	Data Data `json:"data"`
}

func GetUserInfo(accessToken string) (*UserInfo, error) {
	hreq := httpreq.RequestParam{
		EndPoint: "https://graph.facebook.com/me",
		Method:   "GET",
		// Token:    accessToken,
		QueryParams: map[string]string{
			"access_token": accessToken,
			"fields":       "name,email,picture,last_name,first_name,middle_name",
		},
	}

	var res UserInfo

	_, err := httpreq.Do(hreq, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
