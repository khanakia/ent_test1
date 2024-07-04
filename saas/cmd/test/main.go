package main

import (
	"saas/pkg/app"
	"saas/pkg/emailfn"
)

func main() {
	app := app.New()
	app.Cli.Execute()

	emailfn.SendForgotPassword("khanakia@gmail.com", "34343", app.EntDB.Client())

	// mailer, err := mailer.NewMailer("fake", app.EntDB.Client())
	// fmt.Println(err)
	// // fmt.Println(mailer)

	// err = mailer.SendFromTempl("khanakia@gmail.com", "Testing", "forgotpass", nil)
	// fmt.Println(err)

	// url := googleauthfn.GetAuthUrl("", "http://localhost:2304/auth/oauth/callback")
	// fmt.Println(url)

	// userinfo, err := googleauthfn.GetUserInfo("ya29.a0AXooCguPvhSOE9KULC-koxf20BEvabh0kk-yIg3rzPQdXnYJFMP_2-bxhNdQkbvV-9EFba9wHv3rXB5qvIgZC7BzQrN6cNz4DXagO_LBjpHpkUp7QlgbiTePrnoCLGGUkK7_by4UScpzwNVSFxqGLkDfgpi7k-Y0a8l1aCgYKAT4SARISFQHGX2MinHqkMzENepEk1hC4cXowiA0171")

	// fmt.Println(userinfo)
	// fmt.Println(err)

	// fmt.Println(authfn.Hash("admin"))

	// googleauthfn.GetUserInfoFromToken("ya29.a0AXooCgsu1qcofeZ83FOYQ7pW2VdNoozlrY-67yLqOhkVEfEu1lAOAxa-VUalU4TiNvaIBiiz8ilN9196N7Ozo1nHCvknyDJjIAja1Rvf_6o0IM8jUecNAMTRdupAmfOdWWE12iHPtHmPkDHHIGG0gnrfdzLIp1qEP0TJaCgYKAQoSARISFQHGX2Mipryr-eWwWR-mloUeQmtJ7w0171")

	// app.EntDB.Client().OauthConnection.Create().
	// 	SetClientID("266391746836-3femdil6844oj9q27snvo4ojp7ac1a1q.apps.googleusercontent.com").
	// 	SetClientSecret("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").
	// 	SetName("jeoga").
	// 	SetScopes("https://www.googleapis.com/auth/userinfo.email,https://www.googleapis.com/auth/userinfo.profile").
	// 	Save(context.Background())

}
