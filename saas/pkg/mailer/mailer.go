package mailer

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"saas/gen/ent"
	"saas/pkg/appfn"
	"strings"
	"time"

	"github.com/ubgo/goutil"
	mail "github.com/xhit/go-simple-mail/v2"
)

// Mailer will take the mailConn record and set the all the smtp server settings and can use the same
// to send email mailer.SendFromTempl which will take templ record and get the body
type Mailer struct {
	mailConn   *ent.MailConn
	client     *ent.Client
	smtpClient *mail.SMTPClient
}

// func (cls *Mailer) SendMail(from, to, subject, body string) error {
// 	// Send email using SMTP or any other library
// 	return nil
// }

func (cls *Mailer) SendFromTempl(toEmail, subject, templID string, templParams map[string]interface{}) error {
	var err error

	templRender, err := NewTemplRender(templID, cls.client)
	if err != nil {
		return err
	}
	body, err := templRender.Render(templParams)
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("%s <%s>", cls.mailConn.FromName, cls.mailConn.FromEmail)).
		AddTo(toEmail).
		SetSubject(subject)

	email.SetBody(mail.TextHTML, body)

	if email.Error != nil {
		log.Fatal(email.Error)
		return email.Error
	}

	err = email.Send(cls.smtpClient)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (cls *Mailer) SetMailConn(id string) error {
	mailConn, err := cls.client.MailConn.Get(context.Background(), id)

	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = mailConn.Host
	server.Port = mailConn.Port
	server.Username = mailConn.Username
	server.Password = mailConn.Password
	server.Encryption = mail.Encryption(mailConn.Encryption)
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
		return err
	}

	cls.smtpClient = smtpClient
	cls.mailConn = mailConn

	return nil
}

func NewMailer(mailConnID string, client *ent.Client) (*Mailer, error) {
	mailr := &Mailer{
		client: client,
	}

	err := mailr.SetMailConn(mailConnID)
	if err != nil {
		return nil, err
	}

	return mailr, nil
}

func NewDefaultMailer(client *ent.Client) (*Mailer, error) {
	appSetting, err := appfn.GetAppSettings(client)

	if err != nil {
		return nil, err
	}

	if len(appSetting.DefaultMailConnID) == 0 {
		return nil, fmt.Errorf("default mail connection not found")
	}

	return NewMailer(appSetting.DefaultMailConnID, client)
}

// This will take the body of the templ table record and parse and return as string
type TemplRender struct {
	templ      *ent.Templ
	client     *ent.Client
	appSetting *ent.App
}

func (cls *TemplRender) defaultData(data map[string]interface{}) {

	data["app"] = map[string]interface{}{
		"name":     cls.appSetting.Name,
		"email":    cls.appSetting.Email,
		"socialFb": cls.appSetting.SocialFb,
		"socialTw": cls.appSetting.SocialTw,
		"socialIn": cls.appSetting.SocialIn,
		"address":  cls.appSetting.Address,
		"logoUrl":  cls.appSetting.LogoURL,
	}
}

func (cls *TemplRender) wrapInLayout(body string) string {
	layout, err := cls.client.Templ.Get(context.Background(), cls.appSetting.MailLayoutTemplID)
	if err != nil {
		return body
	}

	body = strings.Replace(layout.Body, "{slot}", body, 1)

	return body
}

func (cls *TemplRender) Render(data map[string]interface{}) (string, error) {
	cls.defaultData(data)

	templBody := cls.wrapInLayout(cls.templ.Body)

	var err error
	t1 := template.New("t1")
	t1, err = t1.Parse(templBody)
	if err != nil {
		return "", err
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	goutil.PrintToJSON(data)

	buf := new(bytes.Buffer)
	err = t1.Execute(buf, data)
	if err != nil {

		return "", err
	}

	body := buf.String()

	return body, nil
}

func NewTemplRender(id string, client *ent.Client) (*TemplRender, error) {
	templ, err := client.Templ.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}

	appSetting, err := appfn.GetAppSettings(client)
	if err != nil {
		log.Println("Error getting app settings", err)
		return nil, err
	}

	return &TemplRender{
		templ:      templ,
		client:     client,
		appSetting: appSetting,
	}, nil
}
