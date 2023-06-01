package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Mirobidjon/course"
	"github.com/gin-gonic/gin"
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

type fileUrl struct {
	FileUrl string `json:"file_url"`
}

func (h Handler) addFileToCourse(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You aren't Student!")
		return
	}

	var input fileUrl
	err = c.BindJSON(&input)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	course, err := h.service.AuthStudents.UpdateCourseFileUrl(id, input.FileUrl)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h Handler) uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// id := uuid.NewString()

	if err := c.SaveUploadedFile(file, "public/"+file.Filename); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"url":    "/download-file/" + file.Filename,
	})
}

func (h Handler) downloadFile(c *gin.Context) {
	fileUrl := c.Param("id")
	fmt.Println(fileUrl)

	pwd, _ := os.Getwd()

	c.File(pwd + "/public/" + fileUrl)
}
