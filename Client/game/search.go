package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/utils"
)

func Search(c *models.TcpClient) error {
	var targetUser *entity.UserScoreInfo = nil
	for {
		utils.Cle()
		fmt.Print(view.SearchViewHead)
		if targetUser == nil {
			fmt.Print("\n\n")
		} else {
			fmt.Print(strconv.Itoa(targetUser.rank) + "\t\t" + targetUser.Name + "\t\t" + strconv.Itoa(targetUser.Score) + "\n\n")
		}
		fmt.Print(view.SearchViewTail)
		fmt.Print("请输入：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {

		}
		if text == "1" {

		}
	}
}
