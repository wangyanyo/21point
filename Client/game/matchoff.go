package game

import (
	"fmt"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func MatchOff(c *models.TcpClient) error {
	utils.Cle()
	fmt.Print(view.MatchOffView)

	_, err := utils.RAL(c, enum.MatchOffPacket, c.Token, 0)
	if err != nil {
		return err
	}
	utils.PrintMessage("退出匹配成功！")
	time.Sleep(1 * time.Second)
	return nil
}
