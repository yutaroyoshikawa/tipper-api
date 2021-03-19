package graph

import "github.com/yutaroyoshikawa/tipper-api/infrastructure"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database *infrastructure.Database
}
