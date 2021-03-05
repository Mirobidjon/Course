package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) signInTeacher(c *gin.Context) {
	h.signInMaster(teacherRole, c)
}

func (h Handler) signUpTeacher(c *gin.Context) {
	h.signUpMaster(teacherRole, c)
}

func (h Handler) getAllTeachers(c *gin.Context) {
	h.getAllMasters(teacherRole, c)
}

func (h Handler) getTeachers(c *gin.Context) {
	h.getMaster(teacherRole, c)
}

func (h Handler) updateTeacher(c *gin.Context) {
	h.updateMasters(teacherRole, c)
}

func (h Handler) deleteTeacher(c *gin.Context) {
	role, _, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != directorRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create book!")
		return
	}

	teacherID, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, er.Error())
		return
	}

	err = h.service.AuthMasters.DeleteTeacher(teacherRole, teacherID)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
