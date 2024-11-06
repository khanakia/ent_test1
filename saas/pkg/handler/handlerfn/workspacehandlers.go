package handlerfn

import (
	"context"
	"fmt"
	"saas/gen/ent"

	"saas/gen/ent/user"
	"saas/gen/ent/workspace"
	"saas/gen/ent/workspaceinvite"
	"saas/gen/ent/workspaceuser"
	"saas/pkg/appfn"
	"saas/pkg/handler/handlertypes"
)

func WorkspaceInviteHandler(params handlertypes.WorkspaceInviteInput, cuser *ent.User, client *ent.Client, ctx context.Context) (*ent.WorkspaceInvite, error) {
	canAccess := appfn.CanAccessWorkspace(params.WorkspaceID, cuser.ID, client, ctx)

	if !canAccess {
		return nil, fmt.Errorf("access denied")
	}

	_, err := client.WorkspaceInvite.Query().Where(workspaceinvite.WorkspaceID(params.WorkspaceID), workspaceinvite.Email(params.Email)).First(ctx)

	if err != nil {
		return nil, fmt.Errorf("already invited")
	}

	workspaceInvite, err := client.WorkspaceInvite.Create().
		SetWorkspaceID(params.WorkspaceID).
		SetEmail(params.Email).
		SetRole(params.Role).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return workspaceInvite, nil
}

func WorkspaceInviteCancelHandler(id string, cuser *ent.User, client *ent.Client, ctx context.Context) (bool, error) {
	workspaceInvite, err := client.WorkspaceInvite.Query().Where(workspaceinvite.ID(id)).First(ctx)

	if err != nil {
		return false, fmt.Errorf("already invited")
	}

	canAccess := appfn.CanAccessWorkspace(workspaceInvite.WorkspaceID, cuser.ID, client, ctx)

	if !canAccess {
		return false, fmt.Errorf("access denied")
	}

	err = client.WorkspaceInvite.DeleteOne(workspaceInvite).Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

func WorkspaceAddUserHandler(params handlertypes.WorkspaceAddUserInput, cuser *ent.User, client *ent.Client, ctx context.Context) (*ent.WorkspaceUser, error) {
	canAccess := appfn.CanAccessWorkspace(params.WorkspaceID, cuser.ID, client, ctx)

	if !canAccess {
		return nil, fmt.Errorf("access denied")
	}

	user, err := client.User.Query().Where(user.Email(params.Email)).First(ctx)

	if err != nil && ent.IsNotFound(err) {
		return nil, fmt.Errorf("user not found")
	}

	_, err = client.WorkspaceUser.Query().Where(workspaceuser.WorkspaceID(params.WorkspaceID), workspaceuser.UserID(user.ID)).First(ctx)

	if err == nil {
		return nil, fmt.Errorf("already a part of workspace")
	}

	workspaceUser, err := client.WorkspaceUser.Create().
		SetWorkspaceID(params.WorkspaceID).
		SetUserID(user.ID).
		SetRole(params.Role).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return workspaceUser, nil
}

func WorkspaceRemoveUserHandler(id string, cuser *ent.User, client *ent.Client, ctx context.Context) (bool, error) {
	workspaceUser, err := client.WorkspaceUser.Query().WithWorkspace().Where(workspaceuser.ID(id)).First(ctx)

	if err != nil {
		return false, fmt.Errorf("no user found for workspace")
	}

	canAccess := appfn.CanAccessWorkspace(workspaceUser.WorkspaceID, cuser.ID, client, ctx)

	if !canAccess {
		return false, fmt.Errorf("access denied")
	}

	// workspaceUser, err := client.WorkspaceUser.Query().Where(workspaceuser.WorkspaceID(params.WorkspaceID), workspaceuser.UserID(params.UserID)).First(ctx)

	// if err != nil {
	// 	return false, fmt.Errorf("no user found for workspace")
	// }

	err = client.WorkspaceUser.DeleteOne(workspaceUser).Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

func WorkspaceUpdateUserHandler(params handlertypes.WorkspaceUpdateUserInput, cuser *ent.User, client *ent.Client, ctx context.Context) (*ent.WorkspaceUser, error) {
	workspaceUser, err := client.WorkspaceUser.Query().WithWorkspace().Where(workspaceuser.ID(params.ID)).First(ctx)

	if err != nil {
		return nil, fmt.Errorf("no user found for workspace")
	}

	canAccess := appfn.CanAccessWorkspace(workspaceUser.WorkspaceID, cuser.ID, client, ctx)

	if !canAccess {
		return nil, fmt.Errorf("access denied")
	}

	workspaceUser, err = workspaceUser.Update().
		SetRole(params.Role).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return workspaceUser, nil
}

// leave workspace by workspaceId
func WorkspaceLeaveHandler(id string, cuser *ent.User, client *ent.Client, ctx context.Context) (bool, error) {
	workspaceUser, err := client.WorkspaceUser.Query().WithWorkspace().Where(workspaceuser.WorkspaceID(id), workspaceuser.UserID(cuser.ID)).First(ctx)

	if err != nil {
		return false, fmt.Errorf("no user found for workspace")
	}

	err = client.WorkspaceUser.DeleteOne(workspaceUser).Exec(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

func WorkspaceUpdateHandler(input handlertypes.WorkspaceUpdateInput, cuser *ent.User, client *ent.Client, ctx context.Context) (*ent.Workspace, error) {
	canAccess := appfn.CanAccessWorkspace(input.ID, cuser.ID, client, ctx)

	if !canAccess {
		return nil, fmt.Errorf("access denied")
	}

	workspace, err := client.Workspace.Query().Where(workspace.ID(input.ID)).First(ctx)

	if err != nil {
		return nil, err
	}

	workspace, err = workspace.Update().SetName(input.Name).Save(ctx)

	if err != nil {
		return nil, err
	}

	return workspace, err
}

func WorkspaceGetHandler(id string, cuser *ent.User, client *ent.Client, ctx context.Context) (*ent.Workspace, error) {
	canAccess := appfn.CanAccessWorkspace(id, cuser.ID, client, ctx)

	if !canAccess {
		return nil, fmt.Errorf("access denied")
	}

	workspace, err := client.Workspace.Query().Where(workspace.AppID(cuser.AppID), workspace.ID(id)).First(ctx)

	return workspace, err
}

func WorkspaceUsersHandler(appID, workspaceID string, client *ent.Client, ctx context.Context) ([]*ent.WorkspaceUser, error) {
	return client.WorkspaceUser.Query().WithWorkspace().WithUser().Where(workspaceuser.AppID(appID), workspaceuser.WorkspaceID(workspaceID)).Limit(100).All(ctx)
}

func WorkspaceInvitesHandler(workspaceID string, cuser *ent.User, client *ent.Client, ctx context.Context) ([]*ent.WorkspaceInvite, error) {
	canAccess := appfn.CanAccessWorkspace(workspaceID, cuser.ID, client, ctx)

	if !canAccess {
		return nil, fmt.Errorf("access denied")
	}

	return client.WorkspaceInvite.Query().WithWorkspace().Where(workspaceinvite.WorkspaceID(workspaceID)).All(ctx)
}
