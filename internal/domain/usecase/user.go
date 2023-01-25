package usecase

import (
	"context"
	"regexp"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/dtos"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/model/entities"
	"github.com/ezzycreative1/svc-blog-profile/internal/domain/ports"
	"github.com/ezzycreative1/svc-blog-profile/pkg/errs"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mid"
	"github.com/ezzycreative1/svc-blog-profile/pkg/mvalidator"
)

type userUseCase struct {
	Repo ports.IUserRepository
}

func NewUserUsecase(repo ports.IUserRepository) *userUseCase {
	return &userUseCase{
		Repo: repo,
	}
}

// PasswordValidator - validates the password.
func (uc *userUseCase) passwordValidator(password string) (bool, error) {
	switch {
	case len(password) < 8:
		return false, errs.ErrPasswordLessContain
	case len(password) > 128:
		return false, errs.ErrPasswordLongContain
	case !regexp.MustCompile(`[A-Z]+`).MatchString(password):
		return false, errs.ErrPasswordUppercaseContain
	case !regexp.MustCompile(`[a-z]+`).MatchString(password):
		return false, errs.ErrPasswordLowercaseContain
	case !regexp.MustCompile(`\d+`).MatchString(password):
		return false, errs.ErrPasswordOneNumberContain
	case !regexp.MustCompile(`[!@#~$%^&*()+|_]{1}`).MatchString(password):
		return false, errs.ErrPasswordSpecialCharContain
	default:
		return true, nil
	}
}

// func getRoleName(id int64) string {
// 	var roleName string
// 	switch id {
// 	case 0:
// 		roleName = "user"
// 	case 1:
// 		roleName = "superadmin"
// 	case 2:
// 		roleName = "admin"
// 	case 3:
// 		roleName = "viewer"
// 	}

// 	return roleName
// }

// func (uc *userUseCase) getRoleID(name string) int64 {
// 	var roleID int64
// 	switch name {
// 	case "User":
// 		roleID = 0
// 	case "Superadmin":
// 		roleID = 1
// 	case "Admin":
// 		roleID = 2
// 	case "Viewer":
// 		roleID = 3
// 	}
// 	return roleID
// }

func (uc *userUseCase) Register(ctx context.Context, input dtos.RegisterRequestBody) error {
	if input.Password != input.ConfirmPassword {
		return errs.ErrPasswordMatch
	}

	if validated, err := uc.passwordValidator(input.Password); !validated {
		return err
	}

	password, err := mid.HashPassword(input.Password)
	if err != nil {
		return err
	}

	data := entities.Users{
		FirstName:   input.Firstname,
		LastName:    input.Lastname,
		Email:       input.Email,
		Password:    password,
		PhoneNumber: input.PhoneNumber,
		IsActive:    1,
		CreatedAt:   time.Now().UnixNano(),
		UpdatedAt:   time.Now().UnixNano(),
	}

	if err := uc.Repo.StoreUser(ctx, data); err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) Login(ctx context.Context, input dtos.LoginRequestBody) (*dtos.LoginResponseBody, error) {
	if input.Email == "" && input.Password == "" {
		return nil, errs.ErrBadParamInput
	}

	checkEmail := mvalidator.ValidEmail(input.Email)
	if !checkEmail {
		return nil, errs.ErrEmailWrong
	}

	//validation user
	userData, err := uc.Repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, errs.ErrNotFound
	}

	checkPassword := mid.CheckPasswordHash(input.Password, userData.Password)
	if !checkPassword {
		// check key exis
		return nil, errs.ErrBadParamInput
	}

	// Getting Access Token
	access_token, err := mid.GenerateToken(userData.ID, "access")
	if err != nil {
		return nil, err
	}

	// Getting Refresh Token
	refresh_token, err := mid.GenerateToken(userData.ID, "refresh")
	if err != nil {
		return nil, err
	}

	response := dtos.LoginResponseBody{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}
	return &response, nil
}

func (uc *userUseCase) UpdateUser(ctx context.Context, id int64, input dtos.UpdateUserRequestBody) error {
	userdata, err := uc.Repo.GetUserById(ctx, id)
	if err != nil {
		return errs.ErrNotFound
	}

	data := entities.Users{
		FirstName:   input.Firstname,
		LastName:    input.Lastname,
		Email:       input.Email,
		Password:    userdata.Password,
		PhoneNumber: input.PhoneNumber,
		IsActive:    userdata.IsActive,
		CreatedAt:   userdata.CreatedAt,
		UpdatedAt:   time.Now().UnixNano(),
	}

	if err := uc.Repo.UpdateUser(ctx, data); err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	userdata, err := uc.Repo.GetUserById(ctx, id)
	if err != nil {
		return errs.ErrNotFound
	}

	if err := uc.Repo.DeleteUser(ctx, userdata, id); err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id int64) (*dtos.UserResponseBody, error) {
	userdata, err := uc.Repo.GetUserById(ctx, id)
	if err != nil {
		return nil, errs.ErrNotFound
	}
	result := dtos.UserResponseBody{
		Firstname:   userdata.FirstName,
		Lastname:    userdata.LastName,
		Email:       userdata.Email,
		PhoneNumber: userdata.PhoneNumber,
		IsActive:    userdata.IsActive,
	}
	return &result, nil
}
