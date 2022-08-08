package accounts

import (
	"context"
	"github.com/aacfactory/errors"
)

// CreateArgument
// @title Create account argument
// @description Create account argument
type CreateArgument struct {
	// Username
	// @title username
	// @description account's username
	Username string `json:"username" validate:"required" message:"username is invalid"`
	// Password
	// @title password
	// @description account's password
	Password string `json:"password" validate:"required" message:"password is invalid"`
}

// CreateResult
// @title Create account result
// @description Create account result
type CreateResult struct {
	// I'd
	// @title id
	// @description account id
	Id uint64 `json:"id"`
}

// create
// @fn create
// @validate true
// @authorization false
// @permission false
// @internal false
// @title Create account
// @description >>>
// Create an account
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | accounts_create_failed   | 500     | create account failed         |
// <<<
func create(ctx context.Context, argument CreateArgument) (result *CreateResult, err errors.CodeError) {

	return
}
