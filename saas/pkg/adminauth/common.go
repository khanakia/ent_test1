package adminauth

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/adminuser"
	"saas/gen/ent/session"
)

func ValidateSession(sessionId string, client *ent.Client) (*ent.AdminUser, error) {
	session, err := client.Session.Query().Where(session.ID(sessionId)).First(context.Background())

	if err != nil {
		return nil, fmt.Errorf("session not found")
	}

	userRec, err := client.AdminUser.
		Query().
		Where(adminuser.ID(session.UserID)).
		Only(context.Background())

	if err != nil && ent.IsNotFound(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	if !userRec.Status {
		return nil, errors.New("user is disabled")
	}

	return userRec, nil
}
