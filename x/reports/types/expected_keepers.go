package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"

	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// ProfilesKeeper represents a keeper that deals with profiles
type ProfilesKeeper interface {
	// HasProfile returns true iff the given user has a profile
	HasProfile(ctx sdk.Context, user string) bool
}

// SubspacesKeeper represents a keeper that deals with subspaces
type SubspacesKeeper interface {
	// HasSubspace tells whether the subspace with the given id exists or not
	HasSubspace(ctx sdk.Context, subspaceID uint64) bool

	// HasPermission tells whether the given user has the provided permission inside the subspace with the specified id
	HasPermission(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permission subspacestypes.Permission) bool

	// IterateSubspaces iterates through the subspaces set and performs the given function
	IterateSubspaces(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool))
}

// RelationshipsKeeper represents a keeper that deals with relationships
type RelationshipsKeeper interface {
	// HasUserBlocked tells whether the given blocker has blocked the user inside the provided subspace
	HasUserBlocked(ctx sdk.Context, blocker, user string, subspaceID uint64) bool
}

// PostsKeeper represents a keeper that deals with posts
type PostsKeeper interface {
	// HasPost tells whether the given post exists or not
	HasPost(ctx sdk.Context, subspaceID uint64, postID uint64) bool

	// GetPost returns the post associated with the given id.
	GetPost(ctx sdk.Context, subspaceID uint64, postID uint64) (poststypes.Post, bool)
}
