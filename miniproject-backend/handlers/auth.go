package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/tamabsndra/miniproject/miniproject-backend/models"
	"github.com/tamabsndra/miniproject/miniproject-backend/services"
)

type AuthHandler struct {
    authService  *services.AuthService
    tokenService *services.TokenService
    validator    *validator.Validate
}

func NewAuthHandler(authService *services.AuthService, tokenService *services.TokenService) *AuthHandler {
    return &AuthHandler{
        authService:  authService,
        tokenService: tokenService,
        validator:    validator.New(),
    }
}


// @Summary      Login user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.LoginRequest true "Login credentials"
// @Success      200  {object}  models.LoginResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Router       /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var req models.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request body"})
        return
    }

    if err := h.validator.Struct(req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    response, err := h.authService.Login(req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
        return
    }

    c.SetCookie("authToken", response.Token, int(24*time.Hour.Seconds()), "/", "", true, true)

    c.JSON(http.StatusOK, response)
}

// @Summary      Logout user
// @Description  Invalidate the current JWT token
// @Tags         auth
// @Produce      json
// @Param Authorization header string true "Authorization"
// @Success      200  {object}  models.SuccessResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
    token, exists := c.Get("token")
    if !exists {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "no token found"})
        return
    }

    err := h.tokenService.BlacklistToken(token.(string))
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to logout"})
        return
    }

    c.SetCookie("authToken", "", -1, "/", "", true, true)

    c.JSON(http.StatusOK, models.SuccessResponse{Message: "successfully logged out"})
}

// @Summary      Validate token
// @Description  Validate JWT token and return its metadata
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.ValidateTokenRequest true "Token to validate"
// @Success      200  {object}  models.TokenValidationResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /validate-token [post]
func (h *AuthHandler) ValidateToken(c *gin.Context) {
    var req models.ValidateTokenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request body"})
        return
    }

    if err := h.validator.Struct(req); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }

    result, err := h.tokenService.ValidateToken(req.Token)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "failed to validate token"})
        return
    }

    c.JSON(http.StatusOK, result)
}

// @Summary      Register user
// @Description  Register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.RegisterRequest true "User data"
// @Success      201  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /register [post]
func (h *AuthHandler) Register(c *gin.Context) {

	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid request body"})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.authService.Register(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "user created successfully"})
}

// @Summary      GetMe
// @Description  Get current user
// @Tags         auth
// @Accept       json
// @Param Authorization header string true "Authorization"
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      401  {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	token, exists := c.Get("token")
    if !exists {
        c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "no token found"})
        return
    }

	user, err := h.authService.GetMe(token.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) VerifyCookieToken(c *gin.Context) {
    response := models.VerifyCookieToken{}
    token, err := c.Cookie("authToken")

    if err != nil {
        response.IsAuthenticated = false
        response.Token = ""
        response.IsLoading = false
        c.JSON(http.StatusOK, response)
        return
    }

    _, err = h.tokenService.ValidateToken(token)
    if err != nil {
        response.IsAuthenticated = false
        response.Token = ""
        response.IsLoading = false
        c.JSON(http.StatusOK, response)
        return
    }

    user, err := h.authService.GetMe(token)
    if err != nil {
        response.IsAuthenticated = false
        response.Token = ""
        response.IsLoading = false
        c.JSON(http.StatusOK, response)
        return
    }


    response.User.ID = user.ID
    response.User.Email = user.Email
    response.User.Name = user.Name
    response.User.CreatedAt = user.CreatedAt
    response.User.UpdatedAt = user.UpdatedAt

    response.Token = token
    response.IsAuthenticated = true
    response.IsLoading = false

    c.JSON(http.StatusOK, response)
}
