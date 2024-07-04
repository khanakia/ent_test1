package emailfn

import (
	"fmt"
	"saas/gen/ent"
	"saas/pkg/appfn"
	"saas/pkg/mailer"
)

func SendForgotPassword(email, code string, client *ent.Client) error {
	mailer, err := mailer.NewDefaultMailer(client)
	if err != nil {
		return err
	}

	appSetting := appfn.MustGetAppSettings()

	err = mailer.SendFromTempl(email, "Reset your password", appSetting.AuthFpTemplID, map[string]interface{}{
		"code": code,
		"url":  fmt.Sprintf("%s/auth/reset-password?token=%s&email=%s", appSetting.SiteURL, code, email),
	})

	if err != nil {
		return err
	}

	return nil
}

func RegisterVerify(email, code string, client *ent.Client) error {
	mailer, err := mailer.NewDefaultMailer(client)
	if err != nil {
		return err
	}

	appSetting := appfn.MustGetAppSettings()

	err = mailer.SendFromTempl(email, "Verify your email to complete registration.", appSetting.AuthVerificationTemplID, map[string]interface{}{
		"code": code,
		"url":  fmt.Sprintf("%s/auth/register-verify?token=%s&email=%s", appSetting.SiteURL, code, email),
	})

	if err != nil {
		return err
	}
	return nil
}
