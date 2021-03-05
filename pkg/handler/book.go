package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type bookInput struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (h Handler) createBook(c *gin.Context) {
	var input bookInput

	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create book!")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	bookID, err := h.service.Book.CreateBook(input.Name, input.Author, id)

	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": bookID,
	})
}

func (h Handler) getBook(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create book!")
		return
	}

	bkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.service.GetBookByID(bkID, id)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h Handler) getAllBook(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create book!")
		return
	}

	books, err := h.service.Book.GetAllBooks(id)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h Handler) updateBook(c *gin.Context) {
	var input bookInput

	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create book!")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	bkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Book.UpdateBook(input.Name, input.Author, id, bkID)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h Handler) deleteBook(c *gin.Context) {
	role, id, err := h.getRoleAndID(c)
	if err != nil {
		return
	}

	if role != studentRole {
		NewErrorResponce(c, http.StatusBadRequest, "You can't create book!")
		return
	}

	bkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Book.DeleteBook(id, bkID)

	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
