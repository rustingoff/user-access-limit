package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupControllers struct {
	IGroup
}

func (GroupControllers *GroupControllers) CreateGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Create Group Defined")
}

func (GroupControllers *GroupControllers) DeleteGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Delete Group Defined")
}

func (GroupControllers *GroupControllers) UpdateGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Update Group Defined")
}

func (GroupControllers *GroupControllers) ViewAllGroups(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "View all Groups Defined")
}

func GetGroupControllers() *GroupControllers {
	return &GroupControllers{new(GroupControllers)}
}
