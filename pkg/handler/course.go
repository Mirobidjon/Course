package handler

import (
	"github.com/Mirobidjon/course"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) createCourse(c *gin.Context) {
	var input course.InputCourse

	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != teacherRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create course!")
		return
	}

	if err = c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	courseID, er := h.service.Course.CreateCourse(input, id)
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": courseID,
	})
}

func (h Handler) getCourse(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	courseID, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, "Invalid id params ! ")
	}

	var courseByID course.Course
	if role == studentRole {
		courseByID, err = h.service.AuthStudents.GetCourseByID(courseID, id)
	} else if role == teacherRole {
		courseByID, err = h.service.Course.GetTeacherCourse(id, courseID)
	} else {
		courseByID, err = h.service.Course.GetCourse(courseID)
	}

	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, courseByID)
}

func (h Handler) getAllCourse(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	var allCourse []course.Course
	if role == studentRole {
		allCourse, err = h.service.AuthStudents.GetAllCourse(id)
	} else if role == teacherRole {
		allCourse, err = h.service.Course.GetAllCourse(id)
	} else {
		allCourse, err = h.service.AuthMasters.GetAllTeacherCourse()
	}

	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, allCourse)
}

func (h Handler) updateCourse(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != teacherRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't update course!")
		return
	}

	courseID, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, er.Error())
		return
	}

	var input course.UpdateCourse
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Course.UpdateCourse(input.Name, input.Description, courseID, id)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h Handler) deleteCourse(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != teacherRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't delete course!")
		return
	}

	courseID, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		NewErrorResponce(c, http.StatusBadRequest, er.Error())
		return
	}

	err = h.service.Course.DeleteCourse(courseID, id)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
