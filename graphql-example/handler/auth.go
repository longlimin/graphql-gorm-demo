package handler

import (
	"server_common/graphql-example/conf"
	"server_common/graphql-example/db"
	"server_common/graphql-example/model"
	"server_common/graphql-example/util"

	"github.com/graphql-go/graphql"
)

// AuthType handler
var AuthType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Auth",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if auth, ok := p.Source.(*model.Auth); ok {
					return auth.Token, nil
				}
				return nil, nil
			},
		},
		"user": &graphql.Field{
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if auth, ok := p.Source.(*model.Auth); ok {
					return &auth.User, nil
				}
				return nil, nil
			},
		},
	},
})

func auth(token string) (bool, error) {
	claims := util.CustomClaims{}
	ok, err := util.ValidateJWTToken(token, conf.Config.JWT.Secret, &claims)
	if !ok || err != nil {
		return ok, err
	}
	if ok, checkErr := db.CheckUserValid(claims.UserID, claims.Email); !ok || checkErr != nil {
		return false, checkErr
	}
	return true, nil
}

var loginArgs = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Description: "User Email to login",
		Type:        graphql.NewNonNull(graphql.String),
	},
	"pw": &graphql.ArgumentConfig{
		Description: "User PW to login",
		Type:        graphql.NewNonNull(graphql.String),
	},
}

func login(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)
	pw := p.Args["pw"].(string)
	auth := &model.Auth{
		User: model.User{
			Email: email,
			PW:    pw,
		},
	}
	if ok, err := loginCheck(auth); !ok {
		return nil, err
	}
	return auth, nil
}

func loginCheck(auth *model.Auth) (bool, error) {

	return true, nil
}
