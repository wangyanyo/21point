package ral

import "github.com/wangyanyo/21point/Client/models"

func Connect() {
	if models.Connecting {
		return
	}
	models.ReConnChan <- struct{}{}
	<-models.EndConnChan
}
