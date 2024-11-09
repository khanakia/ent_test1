package constants

const (
	// RoleSaID     = 1
	// RoleMemberID = 2
	WorkspaceRoleOwner = "owner"

	UserRoleMember = "member"
	UserRoleSa     = "sa"

	DirectiveCanAdmin = "canAdmin"
	DirectiveCanApp   = "canApp"

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
