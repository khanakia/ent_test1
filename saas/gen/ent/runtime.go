// Code generated by ent, DO NOT EDIT.

package ent

import (
	"saas/gen/ent/media"
	"saas/gen/ent/mediable"
	"saas/gen/ent/post"
	"saas/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	mediaMixin := schema.Media{}.Mixin()
	mediaMixinFields0 := mediaMixin[0].Fields()
	_ = mediaMixinFields0
	mediaFields := schema.Media{}.Fields()
	_ = mediaFields
	// mediaDescStatus is the schema descriptor for status field.
	mediaDescStatus := mediaFields[16].Descriptor()
	// media.DefaultStatus holds the default value on creation for the status field.
	media.DefaultStatus = mediaDescStatus.Default.(bool)
	// mediaDescID is the schema descriptor for id field.
	mediaDescID := mediaMixinFields0[0].Descriptor()
	// media.DefaultID holds the default value on creation for the id field.
	media.DefaultID = mediaDescID.Default.(func() string)
	mediableMixin := schema.Mediable{}.Mixin()
	mediableMixinFields0 := mediableMixin[0].Fields()
	_ = mediableMixinFields0
	mediableFields := schema.Mediable{}.Fields()
	_ = mediableFields
	// mediableDescID is the schema descriptor for id field.
	mediableDescID := mediableMixinFields0[0].Descriptor()
	// mediable.DefaultID holds the default value on creation for the id field.
	mediable.DefaultID = mediableDescID.Default.(func() string)
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescID is the schema descriptor for id field.
	postDescID := postMixinFields0[0].Descriptor()
	// post.DefaultID holds the default value on creation for the id field.
	post.DefaultID = postDescID.Default.(func() string)
}
