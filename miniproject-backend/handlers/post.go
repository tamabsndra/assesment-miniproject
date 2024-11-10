package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/tamabsndra/miniproject/miniproject-backend/models"
	"github.com/tamabsndra/miniproject/miniproject-backend/services"
)

type PostHandler struct {
	postService *services.PostService
	validator   *validator.Validate
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
		validator:   validator.New(),
	}
}

// @Summary      Create post
// @Description  Create a new post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Param        request body models.CreatePostRequest true "Post data"
// @Success      201  {object}  models.Post
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /posts [post]
func (h *PostHandler) Create(c *gin.Context) {
	var req models.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request body"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	userID := c.GetUint("userID")
	post, err := h.postService.Create(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// @Summary      Get all posts
// @Description  Get all posts
// @Tags         posts
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Success      200  {array}   models.Post
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /posts [get]
func (h *PostHandler) GetAll(c *gin.Context) {
	posts, err := h.postService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// @Summary      Get post by ID
// @Description  Get a post by its ID
// @Tags         posts
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  models.Post
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /posts/{id} [get]
func (h *PostHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid post id"})
		return
	}

	post, err := h.postService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary      User posts
// @Description  Get all posts of a user
// @Tags         posts
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Success      200  {array}   models.Post
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /posts/my [get]
func (h *PostHandler) GetByUserID(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "no token found"})
		return
	}

	posts, err := h.postService.GetByUserID(token.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// @Summary      Update post
// @Description  Update a post by its ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Param        id   path      int  true  "Post ID"
// @Param        request body models.UpdatePostRequest true "Post data"
// @Success      200  {object}  models.Post
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /posts/{id} [put]
func (h *PostHandler) Update(c *gin.Context) {
	var req models.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request body"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid post id"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	post, err := h.postService.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary      Delete post
// @Description  Delete a post by its ID
// @Tags         posts
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /posts/{id} [delete]
func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid post id"})
		return
	}

	if err := h.postService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "post not found"})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "post deleted successfully"})
}

func (h *PostHandler) GetPostDetail(c *gin.Context) {
	// get post with user data
	posts, err := h.postService.GetPostDetail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
