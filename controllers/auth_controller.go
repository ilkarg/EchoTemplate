package controllers

import (
	models "api/models"
	systems "api/systems"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Registration(ctx echo.Context) (err error) {
	db, err := systems.ConnectToDb()

	if err != nil {
		return err
	}

	user := new(models.RegisteringUser)
	if err := ctx.Bind(&user); err != nil {
		return err
	}

	if !user.Checkbox {
		return ctx.JSON(http.StatusOK, models.Response{
			Status:  "ERROR",
			Message: "Для продолжения необходимо установить галочку",
			Data:    nil,
		})
	}

	if valid := systems.ValidateData(user); !valid {
		return ctx.JSON(http.StatusOK, models.Response{
			Status:  "ERROR",
			Message: "Некорректные данные",
			Data:    nil,
		})
	}

	if user.Password != user.PasswordSubmit {
		return ctx.JSON(http.StatusOK, models.Response{
			Status:  "ERROR",
			Message: "Пароль и повтор пароля не совпадают",
			Data:    nil,
		})
	}

	userData := new(models.Users)

	db.First(&userData, "email = ?", user.Email)

	if userData.Email != "" {
		return ctx.JSON(http.StatusOK, models.Response{
			Status:  "ERROR",
			Message: "Аккаунт с указанным Email уже зарегистрирован",
			Data:    nil,
		})
	}

	newUser := models.Users{
		Email:          user.Email,
		Password:       systems.PasswordHashing(user.Password),
		EmailConfirmed: false,
	}

	db.Create(&newUser)

	message := "<p align='center'><a href='http://127.0.0.1:8000/submit/" + newUser.Email + "' style='text-decoration: none; font-size: 20px; color: white; font-weight: bold; background-color: green; border: 1px solid black; border-radius: 5px; padding: 5px'>Подтвердить</a></p>"
	systems.SendEmail(newUser.Email, "Heroes - Подтверждение почты", message)

	if err = systems.WriteInSession(ctx, "user", newUser); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  "OK",
		Message: "Аккаунт успешно зарегистрирован. Для подтверждения перейдите по ссылке в письме отправленном на указанный вами Email",
		Data:    newUser,
	})
}

func Login(ctx echo.Context) error {
	db, err := systems.ConnectToDb()

	if err != nil {
		return err
	}

	user := new(models.AuthorizedUser)
	userData := new(models.Users)

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	if valid := systems.ValidateData(user); !valid {
		return ctx.JSON(http.StatusOK, models.Response{
			Status:  "ERROR",
			Message: "Некорректные логин или пароль",
			Data:    nil,
		})
	}
	user.Password = systems.PasswordHashing(user.Password)

	db.First(&userData, "email = ? AND password = ?", user.Email, user.Password)

	if valid := systems.ValidateData(userData); !valid || (userData.Email != user.Email || userData.Password != user.Password) {
		return ctx.JSON(http.StatusOK, models.Response{
			Status:  "ERROR",
			Message: "Неверные логин или пароль",
			Data:    nil,
		})
	}

	if err = systems.WriteInSession(ctx, "user", userData); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  "OK",
		Message: "Вы успешно вошли в аккаунт",
		Data:    userData,
	})
}

func Logout(ctx echo.Context) error {
	if err := systems.RemoveFromSession(ctx, "user"); err != nil {
		fmt.Println(err)

		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  "ERROR",
			Message: "Во время выхода из аккаунта произошла ошибка",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  "OK",
		Message: "Вы успешно вышли из аккаунта",
		Data:    nil,
	})
}

func EmailConfirmation(ctx echo.Context) error {
	var emailConfirmation models.EmailConfirmation
	if err := ctx.Bind(&emailConfirmation); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  "ERROR",
			Message: "Bad Request",
			Data:    nil,
		})
	}

	userData := new(models.Users)
	db.First(&userData, "email = ?", emailConfirmation.Email)

	if userData.Email == "" {
		return ctx.JSON(http.StatusNotFound, models.Response{
			Status:  "ERROR",
			Message: "Указанный Email не найден",
			Data: 	 nil,
		})
	}

	userData.EmailConfirmed = true
	db.Save(userData)

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  "OK",
		Message: "Email успешно подтвержден",
		Data:    nil,
	})
}
