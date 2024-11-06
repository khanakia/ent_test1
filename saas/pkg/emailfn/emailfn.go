package emailfn

import (
	"fmt"
	"saas/gen/ent"
	"saas/pkg/appfn"
	"saas/pkg/mailer"
)

func SendForgotPassword(appID, email, code string, client *ent.Client) error {
	mailer, err := mailer.NewDefaultMailer(appID, client)
	if err != nil {
		return err
	}

	appSetting := appfn.MustGetAppSettings(appID)

	err = mailer.SendFromTempl(appID, email, "Reset your password", appSetting.AuthFpTemplID, map[string]interface{}{
		"code": code,
		"url":  fmt.Sprintf("%s/auth/reset-password?token=%s&email=%s", appSetting.SiteURL, code, email),
	})

	if err != nil {
		return err
	}

	return nil
}

func RegisterVerify(appID, email, code string, client *ent.Client) error {
	mailer, err := mailer.NewDefaultMailer(appID, client)
	if err != nil {
		return err
	}

	appSetting := appfn.MustGetAppSettings(appID)

	err = mailer.SendFromTempl(appID, email, "Verify your email to complete registration.", appSetting.AuthVerificationTemplID, map[string]interface{}{
		"code": code,
		"url":  fmt.Sprintf("%s/auth/register-verify?token=%s&email=%s", appSetting.SiteURL, code, email),
	})

	if err != nil {
		return err
	}
	return nil
}
