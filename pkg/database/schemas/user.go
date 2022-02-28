package schemas

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/efectn/library-management/pkg/database/ent"
	"github.com/efectn/library-management/pkg/database/ent/hook"
	"github.com/efectn/library-management/pkg/database/ent/user"
	"github.com/efectn/library-management/pkg/globals/api"
	"github.com/efectn/library-management/pkg/utils/convert"
	"github.com/efectn/library-management/pkg/utils/database"
	"golang.org/x/crypto/bcrypt"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("password").
			NotEmpty(),
		field.String("name").
			NotEmpty(),
		field.String("avatar").
			Optional(),
		field.String("phone").
			Optional(),
		field.String("city").
			Optional(),
		field.String("state").
			Optional(),
		field.String("country").
			Optional(),
		field.String("address").
			Optional(),
		field.Int("zip_code").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

// Hook of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					for _, field := range m.Fields() {
						if v, _ := m.Field(field); v == "" {
							switch field {
							case "avatar":
								m.ResetAvatar()
							case "phone":
								m.ResetPhone()
							case "city":
								m.ResetCity()
							case "state":
								m.ResetState()
							case "country":
								m.ResetCountry()
							case "address":
								m.ResetAddress()
							case "zip_code":
								m.ResetZipCode()
							}
						}
					}

					// Hash password
					if v, _ := m.Field("password"); v != "" {
						password, err := bcrypt.GenerateFromPassword(convert.UnsafeBytes(v.(string)), api.App.Config.App.Hash.BcryptCost)
						if err != nil {
							return "", err
						}

						m.SetPassword(convert.UnsafeString(password))
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					for _, field := range m.Fields() {
						if v, _ := m.Field(field); v == "" {
							switch field {
							case "email":
								m.ResetEmail()
							case "password":
								m.ResetPassword()
							case "name":
								m.ResetName()
							case "avatar":
								m.ResetAvatar()
							case "phone":
								m.ResetPhone()
							case "city":
								m.ResetCity()
							case "state":
								m.ResetState()
							case "country":
								m.ResetCountry()
							case "address":
								m.ResetAddress()
							case "zip_code":
								m.ResetZipCode()
							}
						}
					}

					// Hash password
					if v, _ := m.Field("password"); v != nil {
						password, err := bcrypt.GenerateFromPassword(convert.UnsafeBytes(v.(string)), api.App.Config.App.Hash.BcryptCost)
						if err != nil {
							return "", err
						}

						m.SetPassword(convert.UnsafeString(password))
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					if id, e := m.ID(); e {
						tx, err := m.Tx()
						if err != nil {
							return "", database.EntRollback(tx, err)
						}

						user, err := tx.User.Query().Where(user.IDEQ(id)).First(ctx)
						if err != nil {
							return "", database.EntRollback(tx, err)
						}

						_ = api.App.DB.S3.Delete("avatars/" + user.Avatar)
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}
