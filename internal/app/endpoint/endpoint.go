package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DBService interface {
	GetAllUsers() ([]string, error)
}

type Endpoint struct {
	db DBService
}

func New(db DBService) *Endpoint {
	return &Endpoint{db: db}
}

func (ep *Endpoint) GetUsers(ctx *gin.Context) {
	users, _ := ep.db.GetAllUsers()
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
