package controllers

import "github.com/gin-gonic/gin"

type IGroup interface {
	CreateGroup(ctx *gin.Context)
	DeleteGroup(ctx *gin.Context)
	UpdateGroup(ctx *gin.Context)
	ViewAllGroups(ctx *gin.Context)
}
