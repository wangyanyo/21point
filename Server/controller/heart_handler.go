package controller

import (
	"context"

	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/Server/service"
	"github.com/wangyanyo/21point/common/entity"
)

func HeartHandle(ctx context.Context, req *entity.TransfeData, client *models.ClientUser) {
	ser := service.GetServer()
	ser.HeartHandler(ctx, req, client)
}
