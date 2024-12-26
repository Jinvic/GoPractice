package api

import "github.com/gin-gonic/gin"

type todoApi struct{}

var TodoApi = &todoApi{}

func (a *todoApi) GetList(c *gin.Context) {}
func (a *todoApi) Create(c *gin.Context)  {}
func (a *todoApi) Get(c *gin.Context)     {}
func (a *todoApi) Update(c *gin.Context)  {}
func (a *todoApi) Delete(c *gin.Context)  {}
