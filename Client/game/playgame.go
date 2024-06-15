package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func calc(cards []string) int {

}

func PlayGame(c *models.TcpClient) error {
	myCards := make([]string, 21)

	for {
		utils.Cle()
		fmt.Print(view.PlayGameViewHead)
		fmt.Print("你的牌：")
		for _, card := range myCards {
			fmt.Printf("%s  ", card)
		}
		fmt.Printf("\n总计：%d\n\n", calc(myCards))
		fmt.Print(view.PlayGameViewTail)
		fmt.Print("请输入：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		opt := scanner.Text()
		if opt == "0" {
			cardInfo, err := utils.RAL(c, enum.AskCardsPactet, c.Token, "")
			if err != nil {
				continue
			}
			card := cardInfo.Data.(string)
			myCards = append(myCards, card)

			point := calc(myCards)
			if point == 21 {

			} else if point > 21 {

			}

		} else if opt == "1" {

		} else if opt == "2" {

		} else if opt == "3" {

		} else {
			continue
		}
	}
}
