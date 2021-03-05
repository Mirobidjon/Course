package handler

import "github.com/gin-gonic/gin"

func (h Handler) signInDirector(c *gin.Context) {
	h.signInMaster(directorRole, c)
}

func (h Handler) signUpDirector(c *gin.Context) {
	h.signUpMaster(directorRole, c)
}

func (h Handler) getAllDirectors(c *gin.Context) {
	h.getAllMasters(directorRole, c)
}

func (h Handler) getDirectors(c *gin.Context) {
	h.getMaster(directorRole, c)
}

func (h Handler) updateDirector(c *gin.Context) {
	h.updateMasters(directorRole, c)
}
