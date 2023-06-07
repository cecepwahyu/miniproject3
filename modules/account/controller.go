package account

import (
	"crud/dto"
	"crud/middleware"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type ControllerAccount interface {
	CreateAccount(req AccountParam) (any, error)
	GetAccountById(id uint) (FindAccount, error)
	UpdateAccount(req AccountParam, id uint) (any, error)
	DeleteAccount(username string) (any, error)
	LoginAccount(account LoginAccount) (any, error)
	VerifyAccount(req AccountParam, id uint) (any, error)
	ActivateAccount(req AccountParam, id uint) (any, error)
}

type controllerAccount struct {
	accountUseCase UseCaseAccount
}

func (uc controllerAccount) CreateAccount(req AccountParam) (any, error) {

	account, err := uc.accountUseCase.CreateAccount(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create account",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: AccountParam{
			Username: account.Username,
			Password: account.Password,
			RoleID:   account.RoleID,
			Verified: account.Verified,
			Active:   account.Active,
		},
	}
	return res, nil
}

func (uc controllerAccount) GetAccountById(id uint) (FindAccount, error) {
	var res FindAccount
	account, err := uc.accountUseCase.GetAccountById(id)
	if err != nil {
		return FindAccount{}, err
	}
	res.Data = account
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get account",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerAccount) UpdateAccount(req AccountParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.accountUseCase.UpdateAccount(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerAccount) DeleteAccount(username string) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.accountUseCase.DeleteAccount(username)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"

	return res, nil
}

func (uc controllerAccount) LoginAccount(account LoginAccount) (any, error) {
	var res dto.ResponseMeta
	login, err := uc.accountUseCase.GetAccountByUsername(account.Username)
	if err != nil {
		return res, errors.New("Declined")
	}
	if login.Password != account.Password {
		return res, errors.New("Declined")
	}
	if login.Verified != "true" {
		return res, errors.New("Declined")
	}
	if login.Active != "true" {
		return res, errors.New("Declined")
	}
	res = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Selamat sudah berhasil login!",
		ResponseTime: "",
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = account.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set the token expiration time (e.g., 24 hours)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(middleware.JwtSecret))
	if err != nil {
		return res, errors.New("failed to generate JWT token")
	}

	// Return the token in the response
	res.Token = tokenString
	// Other response data

	return res, nil
}

// JWT
//func (uc controllerAccount) HashPassword(password string) (any, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	if err != nil {
//		return res, nil
//	}
//	account.Password = string(bytes)
//	return nil
//}
//func (uc controllerAccount) CheckPassword(providedPassword string) (any, error) {
//	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(providedPassword))
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (uc controllerAccount) VerifyAccount(req AccountParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.accountUseCase.VerifyAccount(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerAccount) ActivateAccount(req AccountParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.accountUseCase.ActivateAccount(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}
