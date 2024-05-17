package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/utils"
)

func Home() {
	for {
		utils.Cle()
		fmt.Print(view.HomeView)
		fmt.Print("请输入：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {

		}
		if text == "1" {

		}
		if text == "2" {

		}
	}
}
