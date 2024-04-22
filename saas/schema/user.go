package schema

import (
	"lace/publicid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	// incrementalEnabled := true

	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				// An example of a dumb ID generator - use a production-ready alternative instead.
				return publicid.Must()
			}),
		field.Time("created_at").
			Optional().
			Default(time.Now),
		field.Time("updated_at").
			Optional().
			Default(time.Now).
			UpdateDefault(time.Now),
		// field.Int("id_num").Annotations(entsql.Annotation{
		// 	Incremental: &incrementalEnabled,
		// }).DefaultFunc(func() int64 {
		// 	return tools.GenSnowflakeID()
		// }),
		// field.Int("id_num").SchemaType(map[string]string{
		// 	dialect.Postgres: postgres.TypeBigSerial,
		// }),
		// field.Int("role_id").Optional().Annotations(
		// 	entgql.MapsTo("roleId"),
		// ),
		field.String("email").Unique(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("company").Optional(),
		field.Bool("status").Optional().Default(false),
		field.String("password").
			// StructTag(`json:"-"`).
			// StructTag(`gqlgen:"-" json="-"`).
			Sensitive().
			Optional(),
		field.String("secret").Sensitive().Optional(),
		field.Int("role_id").Optional(),
		field.String("api_key").Optional(),
		field.Bool("welcome_email_sent").Annotations(
		// entgql.Skip(),
		).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("websites", Website.Type).
		// 	Field("website_id").
		// 	Ref("users").
		// 	Through("websiteuser", WebsiteUser.Type),
	}
	// return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.Skip(entgql.SkipWhereInput),
		// entgql.Skip(entgql.SkipAll),
	}
}
