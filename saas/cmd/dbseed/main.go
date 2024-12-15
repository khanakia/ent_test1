package main

import (
	"context"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/appperm"
	"saas/gen/ent/approle"
	"saas/gen/ent/approleperm"
	"saas/gen/ent/appuser"
	"saas/pkg/app"
	"saas/pkg/constants"
	"strings"

	. "github.com/kkdai/twitter"
)

var twitterClient *ServerClient

func main() {
	app := app.New()
	app.Cli.Execute()

	client := app.EntDB.Client()

	setupApp(constants.DefaultAppID, client)

}

func setupApp(appID string, client *ent.Client) {

	ctx := context.Background()

	createAppUser(ctx, "au1", "owner", appID, client)
	createAppUser(ctx, "au2", "readonly", appID, client)

	perms := []string{"any", "read:any", "create:any", "update:any", "delete:any", "read:post", "create:post", "update:post", "delete:post"}

	for _, permName := range perms {
		createAppPems(ctx, strings.ReplaceAll(permName, ":", "_"), permName, appID, client)
	}

	roleOwner, err := createAppRole(ctx, "owner", "Owner", constants.DefaultAppID, client)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, permName := range perms {
		createAppRolePerms(ctx, roleOwner.ID, strings.ReplaceAll(permName, ":", "_"), appID, client)
	}

	roleReadonly, err := createAppRole(ctx, "readonly", "Readonly", constants.DefaultAppID, client)

	if err != nil {
		fmt.Println(err)
		return
	}

	perms = []string{"read:any"}
	for _, permName := range perms {
		createAppRolePerms(ctx, roleReadonly.ID, strings.ReplaceAll(permName, ":", "_"), appID, client)
	}

}

func createAppRolePerms(ctx context.Context, roleID, permID, appID string, client *ent.Client) (*ent.AppRolePerm, error) {
	r, err := client.AppRolePerm.Query().
		Where(
			approleperm.AppRoleID(roleID),
			approleperm.AppPermID(permID),
		).
		Only(ctx)

	setfields := func(mut *ent.AppRolePermMutation) {
		mut.SetAppRoleID(roleID)
		mut.SetAppPermID(permID)
		mut.SetAppID(appID)
	}

	switch {
	case ent.IsNotFound(err):
		creator := client.AppRolePerm.Create()
		setfields(creator.Mutation())
		r, err = creator.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating approleperm: %v", err)
		}
	case err != nil:
		return nil, fmt.Errorf("failed querying approleperm: %v", err)

	default:
		updater := r.Update()
		setfields(updater.Mutation())
		r, err = updater.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating approleperm: %v", err)
		}
	}

	return r, nil
}

func createAppRole(ctx context.Context, id, name, appID string, client *ent.Client) (*ent.AppRole, error) {
	r, err := client.AppRole.Query().
		Where(
			approle.Name(name),
		).
		Only(ctx)

	setfields := func(mut *ent.AppRoleMutation) {
		if len(id) > 0 {
			mut.SetID(id)
		}

		mut.SetName(name)
		mut.SetAppID(appID)
		mut.SetIsGlobal(true)
	}

	switch {
	case ent.IsNotFound(err):
		creator := client.AppRole.Create()
		setfields(creator.Mutation())
		r, err = creator.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating user: %v", err)
		}
	case err != nil:
		return nil, fmt.Errorf("failed querying user: %v", err)

	default:
		updater := r.Update()
		setfields(updater.Mutation())
		r, err = updater.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating user: %v", err)
		}
	}

	return r, nil
}

func createAppPems(ctx context.Context, id, name, appID string, client *ent.Client) (*ent.AppPerm, error) {
	r, err := client.AppPerm.Query().
		Where(
			appperm.Name(name),
		).
		Only(ctx)

	setfields := func(mut *ent.AppPermMutation) {
		if len(id) > 0 {
			mut.SetID(id)
		}

		mut.SetName(name)
		mut.SetAppID(appID)
	}

	switch {
	case ent.IsNotFound(err):
		creator := client.AppPerm.Create()
		setfields(creator.Mutation())
		r, err = creator.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating permission: %v", err)
		}
	case err != nil:
		return nil, fmt.Errorf("failed querying permission: %v", err)

	default:
		updater := r.Update()
		setfields(updater.Mutation())
		r, err = updater.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating permission: %v", err)
		}
	}

	return r, nil
}

func createAppUser(ctx context.Context, adminUserID, appRoleID, appID string, client *ent.Client) (*ent.AppUser, error) {

	r, err := client.AppUser.Query().
		Where(
			appuser.AdminUserID(adminUserID),
			appuser.AppRoleID(appRoleID),
			appuser.AppID(appID),
		).
		Only(ctx)

	setfields := func(mut *ent.AppUserMutation) {

		mut.SetAdminUserID(adminUserID)
		mut.SetAppRoleID(appRoleID)
		mut.SetAppID(appID)
	}

	switch {
	case ent.IsNotFound(err):
		creator := client.AppUser.Create()
		setfields(creator.Mutation())
		r, err = creator.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating app user: %v", err)
		}
	case err != nil:
		return nil, fmt.Errorf("failed querying app user: %v", err)

	default:
		updater := r.Update()
		setfields(updater.Mutation())
		r, err = updater.Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating app user: %v", err)
		}
	}

	return r, nil
}
