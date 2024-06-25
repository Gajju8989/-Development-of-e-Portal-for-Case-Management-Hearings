package handler

import (
	"Pc_Build_Service/Internal/genericresponse"
	"Pc_Build_Service/Internal/userinterface/service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *impl) CreateUserHandler(c *gin.Context) {
	var (
		req = &model.UserRegistrationRequest{}
		err = c.BindJSON(req)
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable To Bind")
		return
	}

	err = h.userInterfaceSvc.UserRegistration(c, req)
	if err != nil {
		genericErr, ok := err.(*genericresponse.GenericResponse)
		if ok {
			c.JSON(genericErr.StatusCode, genericErr)
		} else {
			c.JSON(http.StatusInternalServerError, genericresponse.GenericResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal server error",
			})
		}
		return
	}
	c.Status(http.StatusOK)
}
