package graph

import (
	"firebase.google.com/go/v4/auth"
	"github.com/yutaroyoshikawa/tipper-api/infrastructure"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database  *infrastructure.Database
	LoginUser *auth.Token
}
