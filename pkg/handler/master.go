package handler

import (
	"github.com/Mirobidjon/course"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) signInMaster(role string, c *gin.Context) {
	var input signIn

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.AuthMasters.GenerateTokenMaster(input.Username, input.Password, role)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h Handler) signUpMaster(role string, c *gin.Context) {
	var input course.MasterInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AuthMasters.CreateMaster(input, role)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllMasters(role string, c *gin.Context) {
	masters, err := h.service.AuthMasters.GetAllMaster(role)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, masters)
}

func (h *Handler) getMaster(role string, c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	master, er := h.service.AuthMasters.GetMasterByID(role, id)
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusOK, master)
}

func (h *Handler) updateMasters(roleMaster string, c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != roleMaster {
		NewErrorResponce(c, http.StatusBadRequest, "You can't update!")
		return
	}

	var input course.MasterInput
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.AuthMasters.UpdateMaster(input, id)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
