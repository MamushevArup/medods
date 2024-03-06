package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Access string `json:"access"`
}

func (h *handler) createToken(c *gin.Context) {

	guid := c.Param("guid")
	if guid == "" {
		newErrorResponse(c, http.StatusBadRequest, "user id not provided")
		return
	}

	tokenModel, err := h.service.Token.Generate(c, guid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r := response{Access: tokenModel.AccessToken}

	c.SetCookie("refresh_token", tokenModel.RefreshToken, maxAge, "/auth", "localhost", false, true)
	c.JSON(200, r)
}

// refresh get guid from url and refresh token from the cookie
func (h *handler) refresh(c *gin.Context) {

	guid := c.Param("guid")
	if guid == "" {
		newErrorResponse(c, http.StatusBadRequest, "user id not provided")
		return
	}

	refresh, err := c.Cookie("refresh_token")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if refresh == "" {
		newErrorResponse(c, http.StatusBadRequest, "no refresh provided")
		return
	}

	token, err := h.service.Token.UpdateToken(c, guid, refresh)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r := response{Access: token.AccessToken}

	// set cookie http-only which will be safe from js scripts
	c.SetCookie("refresh_token", token.RefreshToken, maxAge, "/auth", "localhost", false, true)

	c.JSON(200, r)
}
