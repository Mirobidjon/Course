package handler

import (
	"net/http"

	"github.com/Mirobidjon/course/pkg/service"
	"github.com/gin-gonic/gin"
)

const (
	teacherRole  = "teacher"
	studentRole  = "student"
	directorRole = "director"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", h.getMenu)

	auth := router.Group("/auth")
	{
		student := auth.Group("/student")
		{
			student.POST("/sign-up", h.signUpStudent)
			student.POST("/sign-in", h.signInStudent)
		}

		master := auth.Group("/director")
		{
			master.POST("/sign-up", h.signInDirector)
			master.POST("/sign-in", h.signInDirector)
		}

		teacher := auth.Group("/teacher")
		{
			teacher.POST("/sign-up", h.signUpTeacher)
			teacher.POST("/sign-in", h.signInTeacher)
		}

	}

	router.GET("/download-file/:id", h.downloadFile)

	router.GET("/api", h.getApi)

	api := router.Group("/api", h.identify)
	{
		student := api.Group("/students")
		{
			student.GET("/", h.getAllStudent)
			student.GET("/:id", h.getStudent)
			student.PUT("/", h.updateStudent)
			student.DELETE("/:id", h.deleteStudent)
		}

		book := api.Group("/book")
		{
			book.GET("/", h.getAllBook)
			book.GET("/:id", h.getBook)
			book.POST("/", h.createBook)
			book.PUT("/:id", h.updateBook)
			book.DELETE("/:id", h.deleteBook)
		}

		director := api.Group("/director")
		{
			director.GET("/", h.getAllDirectors)
			director.GET("/:id", h.getDirectors)
			director.PUT("/", h.updateDirector)
		}

		teacher := api.Group("/teacher")
		{
			teacher.GET("/", h.getAllTeachers)
			teacher.GET("/:id", h.getTeachers)
			teacher.PUT("/", h.updateTeacher)
			teacher.DELETE("/:id", h.deleteTeacher)
		}

		course := api.Group("/course")
		{
			course.GET("/", h.getAllCourse)
			course.POST("/", h.createCourse)
			course.GET("/:id", h.getCourse)
			course.PUT("/:id", h.updateCourse)
			course.DELETE("/:id", h.deleteCourse)
		}

		api.PUT("/add-file-to-course", h.addFileToCourse)
		api.POST("/upload-file", h.uploadFile)
	}

	return router
}

func (h *Handler) getMenu(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"You are at the menu": "ok",
		"You get api":         "ok",
	})
}

func (h *Handler) getApi(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"You are at api the menu": "ok",
	})
}
