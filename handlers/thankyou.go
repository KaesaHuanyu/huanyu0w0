package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"huanyu0w0/model"
)

func (h *Handler) ThankYou(c echo.Context) (err error) {
	data := &struct {
		model.Cookie
	}{}
	if err = data.Cookie.ReadCookie(c); err == nil {
		data.IsLogin = true
	} else {
		return c.Redirect(http.StatusFound, "/login?path=3Q~")
	}
	return c.Render(http.StatusOK, "thankyou", data)
}