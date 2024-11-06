// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/app"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// App is the model entity for the App schema.
type App struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Copyright holds the value of the "copyright" field.
	Copyright string `json:"copyright,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// SocialTw holds the value of the "social_tw" field.
	SocialTw string `json:"social_tw,omitempty"`
	// SocialFb holds the value of the "social_fb" field.
	SocialFb string `json:"social_fb,omitempty"`
	// SocialIn holds the value of the "social_in" field.
	SocialIn string `json:"social_in,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL string `json:"logo_url,omitempty"`
	// SiteURL holds the value of the "site_url" field.
	SiteURL string `json:"site_url,omitempty"`
	// DefaultMailConnID holds the value of the "default_mail_conn_id" field.
	DefaultMailConnID string `json:"default_mail_conn_id,omitempty"`
	// MailLayoutTemplID holds the value of the "mail_layout_templ_id" field.
	MailLayoutTemplID string `json:"mail_layout_templ_id,omitempty"`
	// WsapceInviteTemplID holds the value of the "wsapce_invite_templ_id" field.
	WsapceInviteTemplID string `json:"wsapce_invite_templ_id,omitempty"`
	// WsapceSuccessTemplID holds the value of the "wsapce_success_templ_id" field.
	WsapceSuccessTemplID string `json:"wsapce_success_templ_id,omitempty"`
	// AuthFpTemplID holds the value of the "auth_fp_templ_id" field.
	AuthFpTemplID string `json:"auth_fp_templ_id,omitempty"`
	// AuthWelcomeEmailTemplID holds the value of the "auth_welcome_email_templ_id" field.
	AuthWelcomeEmailTemplID string `json:"auth_welcome_email_templ_id,omitempty"`
	// AuthVerificationTemplID holds the value of the "auth_verification_templ_id" field.
	AuthVerificationTemplID string `json:"auth_verification_templ_id,omitempty"`
	// AuthEmailVerify holds the value of the "auth_email_verify" field.
	AuthEmailVerify string `json:"auth_email_verify,omitempty"`
	// AdminUserID holds the value of the "admin_user_id" field.
	AdminUserID  string `json:"admin_user_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*App) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case app.FieldID, app.FieldName, app.FieldCopyright, app.FieldEmail, app.FieldAddress, app.FieldSocialTw, app.FieldSocialFb, app.FieldSocialIn, app.FieldLogoURL, app.FieldSiteURL, app.FieldDefaultMailConnID, app.FieldMailLayoutTemplID, app.FieldWsapceInviteTemplID, app.FieldWsapceSuccessTemplID, app.FieldAuthFpTemplID, app.FieldAuthWelcomeEmailTemplID, app.FieldAuthVerificationTemplID, app.FieldAuthEmailVerify, app.FieldAdminUserID:
			values[i] = new(sql.NullString)
		case app.FieldCreatedAt, app.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the App fields.
func (a *App) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case app.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case app.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case app.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case app.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case app.FieldCopyright:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field copyright", values[i])
			} else if value.Valid {
				a.Copyright = value.String
			}
		case app.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case app.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				a.Address = value.String
			}
		case app.FieldSocialTw:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field social_tw", values[i])
			} else if value.Valid {
				a.SocialTw = value.String
			}
		case app.FieldSocialFb:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field social_fb", values[i])
			} else if value.Valid {
				a.SocialFb = value.String
			}
		case app.FieldSocialIn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field social_in", values[i])
			} else if value.Valid {
				a.SocialIn = value.String
			}
		case app.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				a.LogoURL = value.String
			}
		case app.FieldSiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field site_url", values[i])
			} else if value.Valid {
				a.SiteURL = value.String
			}
		case app.FieldDefaultMailConnID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_mail_conn_id", values[i])
			} else if value.Valid {
				a.DefaultMailConnID = value.String
			}
		case app.FieldMailLayoutTemplID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mail_layout_templ_id", values[i])
			} else if value.Valid {
				a.MailLayoutTemplID = value.String
			}
		case app.FieldWsapceInviteTemplID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wsapce_invite_templ_id", values[i])
			} else if value.Valid {
				a.WsapceInviteTemplID = value.String
			}
		case app.FieldWsapceSuccessTemplID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wsapce_success_templ_id", values[i])
			} else if value.Valid {
				a.WsapceSuccessTemplID = value.String
			}
		case app.FieldAuthFpTemplID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_fp_templ_id", values[i])
			} else if value.Valid {
				a.AuthFpTemplID = value.String
			}
		case app.FieldAuthWelcomeEmailTemplID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_welcome_email_templ_id", values[i])
			} else if value.Valid {
				a.AuthWelcomeEmailTemplID = value.String
			}
		case app.FieldAuthVerificationTemplID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_verification_templ_id", values[i])
			} else if value.Valid {
				a.AuthVerificationTemplID = value.String
			}
		case app.FieldAuthEmailVerify:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_email_verify", values[i])
			} else if value.Valid {
				a.AuthEmailVerify = value.String
			}
		case app.FieldAdminUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field admin_user_id", values[i])
			} else if value.Valid {
				a.AdminUserID = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the App.
// This includes values selected through modifiers, order, etc.
func (a *App) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this App.
// Note that you need to call App.Unwrap() before calling this method if this App
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *App) Update() *AppUpdateOne {
	return NewAppClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the App entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *App) Unwrap() *App {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: App is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *App) String() string {
	var builder strings.Builder
	builder.WriteString("App(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("copyright=")
	builder.WriteString(a.Copyright)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(a.Address)
	builder.WriteString(", ")
	builder.WriteString("social_tw=")
	builder.WriteString(a.SocialTw)
	builder.WriteString(", ")
	builder.WriteString("social_fb=")
	builder.WriteString(a.SocialFb)
	builder.WriteString(", ")
	builder.WriteString("social_in=")
	builder.WriteString(a.SocialIn)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(a.LogoURL)
	builder.WriteString(", ")
	builder.WriteString("site_url=")
	builder.WriteString(a.SiteURL)
	builder.WriteString(", ")
	builder.WriteString("default_mail_conn_id=")
	builder.WriteString(a.DefaultMailConnID)
	builder.WriteString(", ")
	builder.WriteString("mail_layout_templ_id=")
	builder.WriteString(a.MailLayoutTemplID)
	builder.WriteString(", ")
	builder.WriteString("wsapce_invite_templ_id=")
	builder.WriteString(a.WsapceInviteTemplID)
	builder.WriteString(", ")
	builder.WriteString("wsapce_success_templ_id=")
	builder.WriteString(a.WsapceSuccessTemplID)
	builder.WriteString(", ")
	builder.WriteString("auth_fp_templ_id=")
	builder.WriteString(a.AuthFpTemplID)
	builder.WriteString(", ")
	builder.WriteString("auth_welcome_email_templ_id=")
	builder.WriteString(a.AuthWelcomeEmailTemplID)
	builder.WriteString(", ")
	builder.WriteString("auth_verification_templ_id=")
	builder.WriteString(a.AuthVerificationTemplID)
	builder.WriteString(", ")
	builder.WriteString("auth_email_verify=")
	builder.WriteString(a.AuthEmailVerify)
	builder.WriteString(", ")
	builder.WriteString("admin_user_id=")
	builder.WriteString(a.AdminUserID)
	builder.WriteByte(')')
	return builder.String()
}

// Apps is a parsable slice of App.
type Apps []*App
