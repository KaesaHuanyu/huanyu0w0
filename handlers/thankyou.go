package handlers

import (
	"github.com/labstack/echo"
	"huanyu0w0/model"
	"net/http"
)

func (h *Handler) ThankYou(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	}
	return c.Render(http.StatusOK, "thankyou", data)
}
