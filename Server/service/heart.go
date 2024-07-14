package service

import (
	"context"
	"time"

	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) HeartHandler(ctx context.Context, req *entity.TransfeData, client *models.ClientUser) {
	client.LastTime = time.Now()
}
