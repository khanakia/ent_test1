package constants

const (
	// RoleSaID     = 1
	// RoleMemberID = 2
	WorkspaceRoleOwner = "owner"

	UserRoleMember = "member"
	UserRoleSa     = "sa"

	DefaultAppID = "a1"

	// check if the current user has permission to perform the admin operations
	DirectiveCanAdmin = "canAdmin"

	// check if current user belongs to the app or has permission for the query he requested
	DirectiveCanApp = "canApp"

	// DirectivePublicSkip = "publicSkip"

	// define if the e.g. type Post {} can be in `gql/gql` go module eng.graphql
	DirectiveAppGql = "appGql"

	_ = iota
	LockReadonly
	LockU
	LockD

	_ = iota
	A
	B
)

const (
	PrefixPostType   = "a"
	PrefixPostStatus = "b"
)
