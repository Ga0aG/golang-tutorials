package handler

import (
	"gin-app/internal/model"
	"gin-app/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户HTTP处理器
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		service: svc,
	}
}

// CreateUser 创建用户
// @Summary 创建用户
// @Accept json
// @Produce json
// @Param body body model.CreateUserRequest true "用户信息"
// @Success 201 {object} model.User
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		Code:    0,
		Message: "创建成功",
		Data:    user,
	})
}

// GetUser 获取用户
// @Summary 获取用户详情
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} model.User
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, model.Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "获取成功",
		Data:    user,
	})
}

// ListUsers 获取用户列表
// @Summary 获取用户列表
// @Produce json
// @Success 200 {array} model.User
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "获取成功",
		Data:    users,
	})
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param body body model.UpdateUserRequest true "用户信息"
// @Success 200 {object} model.User
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, model.Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "更新成功",
		Data:    user,
	})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Produce json
// @Param id path int true "用户ID"
// @Success 204
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
