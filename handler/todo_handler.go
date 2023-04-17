package handler

import (
	"errors"
	"github.com/thalerngsak/todoapplication/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thalerngsak/todoapplication/service"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (h *TodoHandler) Create(c *gin.Context) {

	v, _ := c.Get("user_id")

	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(v.(uint), todo.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h *TodoHandler) Update(c *gin.Context) {

	v, _ := c.Get("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input struct {
		Message string `json:"message"`
		Done    *bool  `json:"done"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Message == "" && input.Done == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message or done flag must be provided"})
		return
	}

	message := input.Message
	done := false
	if input.Done != nil {
		done = *input.Done
	}

	if err := h.service.Update(uint(id), v.(uint), message, done); err != nil {
		if errors.Is(err, errors.New("record not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *TodoHandler) Delete(c *gin.Context) {

	v, _ := c.Get("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id), v.(uint)); err != nil {
		if errors.Is(err, errors.New("record not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *TodoHandler) GetByID(c *gin.Context) {

	v, _ := c.Get("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.service.GetByID(uint(id), v.(uint))
	if err != nil {
		if errors.Is(err, errors.New("record not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      todo.ID,
		"message": todo.Message,
		"done":    todo.Done,
		"created": todo.CreatedAt,
	})
}

func (h *TodoHandler) List(c *gin.Context) {

	v, _ := c.Get("user_id")

	todos, err := h.service.List(v.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, todo := range todos {
		response = append(response, gin.H{
			"id":      todo.ID,
			"message": todo.Message,
			"done":    todo.Done,
			"created": todo.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *TodoHandler) MarkAsDone(c *gin.Context) {

	v, _ := c.Get("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.service.GetByID(uint(id), v.(uint))

	if todo.Done {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "todo already marked as done"})
		return
	}

	todo.Done = true

	if err := h.service.Update(uint(id), v.(uint), todo.Message, todo.Done); err != nil {
		if errors.Is(err, errors.New("record not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
