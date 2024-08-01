package appfn

import (
	"context"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/workspaceuser"
	"saas/pkg/app"
	"saas/pkg/constants"
)

func WithTx[T any](ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (*T, error)) (*T, error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	result, err := fn(tx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}
	return result, nil
}

func CanAccessWorkspace(workspaceID, userID string, client *ent.Client, ctx context.Context) bool {
	_, err := client.WorkspaceUser.Query().Where(workspaceuser.WorkspaceID(workspaceID), workspaceuser.UserID(userID)).First(ctx)
	fmt.Println("sfs", err)
	return err == nil
}

// we will use this functio to create default workspace for the new user registered
func CreateDefaultWorkspaceForUser(userID string, client *ent.Client, ctx context.Context) error {
	workspace, err := client.Workspace.Create().SetName("Personal").SetIsPersonal(true).SetUserID(userID).Save(ctx)
	if err != nil {
		return err
	}

	_, err = client.WorkspaceUser.Create().
		SetWorkspaceID(workspace.ID).
		SetUserID(userID).
		SetRole(constants.WorkspaceRoleOwner).
		Save(ctx)

	return err
}

func GetAppSettings(client *ent.Client) (*ent.AppSetting, error) {
	return client.AppSetting.Query().First(context.Background())
}

func MustGetAppSettings() *ent.AppSetting {
	client := app.GetPlugins().EntDB.Client()
	record, err := client.AppSetting.Query().First(context.Background())
	if err != nil {
		panic(err)
	}
	return record
}

func IsUserSA(user *ent.User) bool {
	return user.RoleID == constants.UserRoleSa
}
