package service

import (
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type handlers struct {
	logger *zap.SugaredLogger
	svc    *UserService
}

func New(logger *zap.SugaredLogger, svc *UserService) *handlers {
	return &handlers{
		logger: logger,
		svc:    svc,
	}
}

type registerRequest struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	TenantId string   `json:"tenant_id"`
	Roles    []string `json:"roles"`
	IsAdmin  bool     `json:"is_admin"`
}

type registerResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	TenantID  string    `json:"tenant_id"`
	CreatedAt time.Time `json:"created_at"`
}

// @Summary Register a new user
// @Description creates a new user account using email and password
// @Tags Auth
// @Accent json
// @Produce json
// @Param request body registerRequest true "User credentials"
// @Success 201 {object} registerResponse
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (h *handlers) PostAuthRegister(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	u, err := h.svc.Register(req.Username, req.Email, req.Password, req.TenantId, req.Roles, req.IsAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registerResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		TenantID:  u.TenantID,
		CreatedAt: u.CreatedAt,
	})
}

// @Summary Login user
// @Description Authenticates a user and returns access & refresh tokens
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body map[string]string true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (handlers) PostAuthLogin(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Login endpoint"})
}

// @Summary Refresh access token
// @Description Issues new access token using refresh token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body map[string]string true "Refresh token request"
// @Success 200 {object} map[string]string
// @Router /auth/refresh [post]
func (handlers) PostAuthRefresh(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Refresh endpoint"})
}

// @Summary Logout user
// @Description Invalidates active token/session
// @Tags Auth
// @Security BearerAuth
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func (handlers) PostAuthLogout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout endpoint"})
}

// @Summary Get current user profile
// @Description Returns details of authenticated user
// @Tags Auth
// @Security BearerAuth
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /auth/me [get]
func (handlers) GetAuthMe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile endpoint"})
}

// @Summary Change password
// @Description Allows logged-in user to change password
// @Tags Auth
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param request body map[string]string true "Password change request"
// @Success 200 {object} map[string]string
// @Router /auth/change-password [post]
func (handlers) PostAuthChangePassword(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Change password endpoint"})
}

// @Summary Forgot password
// @Description Initiates password reset via email
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body map[string]string true "Forgot password request"
// @Success 200 {object} map[string]string
// @Router /auth/forgot-password [post]
func (handlers) PostAuthForgotPassword(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Forgot password endpoint"})
}

// @Summary Reset password
// @Description Completes password reset with new password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body map[string]string true "Reset password request"
// @Success 200 {object} map[string]string
// @Router /auth/reset-password [post]
func (handlers) PostAuthResetPassword(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Reset password endpoint"})
}

// @Summary Health Check
// @Description Check if API Gateway is running
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /healthz [get]
func (handlers) GetHealthz(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
