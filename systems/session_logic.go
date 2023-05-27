package systems

import (
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

func WriteInSession(ctx echo.Context, key string, value interface{}) error {
	store := echosession.FromContext(ctx)
	store.Set(key, value)
	return store.Save()
}

func RemoveFromSession(ctx echo.Context, key string) error {
	store := echosession.FromContext(ctx)
	store.Delete(key)
	return store.Save()
}
