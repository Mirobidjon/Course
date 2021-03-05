package handler

import (
	"github.com/Mirobidjon/course"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type signIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h Handler) signInStudent(c *gin.Context) {
	var input signIn

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h Handler) signUpStudent(c *gin.Context) {
	var input course.SignUpStudent

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateStudent(input)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h Handler) getStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	studentByID, er := h.service.AuthStudents.GetStudentByID(id)
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, studentByID)
}

func (h Handler) getAllStudent(c *gin.Context) {
	students, err := h.service.AuthStudents.GetAllStudents()
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h Handler) updateStudent(c *gin.Context) {
	var input course.UpdateStudent

	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You aren't Student!")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.AuthStudents.UpdateStudent(input, id)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h Handler) deleteStudent(c *gin.Context) {
	role, _, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != teacherRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't delete student!")
		return
	}

	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.AuthStudents.DeleteStudent(studentID)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
