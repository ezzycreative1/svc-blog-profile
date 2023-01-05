package errs

import "errors"

var (
	ErrInternalServerError        = errors.New("internal Server Error")
	ErrNotFound                   = errors.New("your requested data is not found")
	ErrConflict                   = errors.New("your data already exist")
	ErrBadParamInput              = errors.New("given Param is not valid")
	ErrDuplicate                  = errors.New("duplicated entry")
	ErrEmailWrong                 = errors.New("format email wrong")
	ErrGenerateToken              = errors.New("generate token failed")
	ErrPasswordMatch              = errors.New("password do not match")
	ErrPasswordLessContain        = errors.New("password must be at least 8 characters long")
	ErrPasswordLongContain        = errors.New("password must be less than 128 characters long")
	ErrPasswordUppercaseContain   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordLowercaseContain   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordOneNumberContain   = errors.New("password must contain at least one number")
	ErrPasswordSpecialCharContain = errors.New("password must contain at least one special character")
)
