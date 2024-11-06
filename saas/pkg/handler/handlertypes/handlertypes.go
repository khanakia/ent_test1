package handlertypes

import "saas/gen/ent"

type WorkspaceInviteInput struct {
	WorkspaceID string `json:"workspaceId"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type WorkspaceAddUserInput struct {
	WorkspaceID string `json:"workspaceId"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type WorkspaceUpdateUserInput struct {
	ID   string `json:"id"`
	Role string `json:"userId"`
}

type WorkspaceRemoveUserInput struct {
	ID string `json:"id"`
}

type WorkspaceUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RegisterVerifyInput struct {
	AppID string
	Token string `json:"token"` // temp.id
	Email string `json:"email"`

	// need for creating login session
	IP        string
	UserAgent string
	Payload   string
}

type LoginResponse struct {
	Token string    `json:"token"`
	Me    *ent.User `json:"me"`
}
