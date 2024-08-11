package resolverfn

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"
	"gql/graph/model"
	"lace/publicid"
	"lace/util"
	"net/url"
	"saas/pkg/oauth/oauthconnectionfn"

	"github.com/ubgo/goutil"
)

// OauthRequest is the resolver for the oauthRequest field.
func (r *mutationResolver) OauthRequest(ctx context.Context, input *model.OauthRequestInput) (string, error) {
	oauthRequester, err := oauthconnectionfn.NewOauthRequester(input.OatConnectionID, r.Plugin.EntDB.Client())

	if err != nil {
		return "", fmt.Errorf("something went wrong")
	}

	if input.RedirectURL != nil {
		oauthRequester.SetRedirectURL(*input.RedirectURL)
	}

	uid := publicid.Must()

	oauthRequester.SetState(uid)

	r.Plugin.Cache.Put(uid, oauthconnectionfn.OauthRequestCache{
		OauthConnectionID: input.OatConnectionID,
		Metadata:          input.Metadata,
		RedirectURL:       oauthRequester.GetRedirectURL(),
		// Provider:          oauthRequester.GetOauthConnection().Provider,
	}, 15000)

	url := oauthRequester.GetAuthUrl()
	return url, nil
}

// OauthCallback is the resolver for the oauthCallback field.
func (r *mutationResolver) OauthCallback(ctx context.Context, callbackURL string) (*model.OauthCallbackResponse, error) {
	myUrl, _ := url.Parse(callbackURL)
	params, _ := url.ParseQuery(myUrl.RawQuery)
	// fmt.Println(params)
	code := params.Get("code")
	cacheId := params.Get("state")

	oauthReqCacheAny := r.Plugin.Cache.Get(cacheId)
	goutil.PrintToJSON(oauthReqCacheAny)

	var oauthReqCache oauthconnectionfn.OauthRequestCache
	err := util.InterfaceToStruct(oauthReqCacheAny, &oauthReqCache)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("something went wrong")
	}

	oauthRequester, err := oauthconnectionfn.NewOauthRequester(oauthReqCache.OauthConnectionID, r.Plugin.EntDB.Client())

	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	if len(oauthReqCache.RedirectURL) > 0 {
		oauthRequester.SetRedirectURL(oauthReqCache.RedirectURL)
	}

	token, err := oauthRequester.TokenFromCode(code)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("something went wrong")
	}

	r.Plugin.Cache.Del(cacheId)

	uid := publicid.Must()
	r.Plugin.Cache.Put(uid, oauthconnectionfn.OauthResponseCache{
		Code:              code,
		OauthConnectionID: oauthReqCache.OauthConnectionID,
		Token:             token,
		Provider:          oauthRequester.GetOauthConnection().Provider,
	}, 10000)

	oauthConnection := oauthRequester.GetOauthConnection()

	// return the provider so at frontend we can decide where to redirect
	return &model.OauthCallbackResponse{
		Provider: oauthConnection.Provider,
		CacheID:  uid,
		Metadata: oauthReqCache.Metadata,
	}, nil
}