package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

func RankList(c *models.TcpClient) error {
	cnt := 1
	for {
		utils.Cle()
		log.Println("查看排行榜", cnt)
		fmt.Print(view.RankListViewHead)
		if _, err := c.Send(entity.NewTransfeData(enum.RankListPactet, "", cnt)); err != nil {
			models.Rconn <- true
			log.Println("断线重连", err)
			fmt.Println("连接已断开，正在尝试重连...")
			time.Sleep(1 * time.Second)
			return err
		}
		rankList := <-c.CmdChan
		if rankList.Cmd == enum.RankListPactet {
			for i, v := range rankList.Data.([]entity.UserScoreInfo) {
				fmt.Println(cnt+i+1, "\t\t", v.Name, "\t\t", v.Score)
			}
		} else {
			log.Println("获取排行榜错误")
			fmt.Println("获取排行榜错误！")
			time.Sleep(1 * time.Second)
			return &myerror.RankListError{}
		}
		fmt.Print(view.RankListViewTail)
		fmt.Print("请输入: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {
			if cnt == 1 {
				fmt.Println("这是第一页！")
				time.Sleep(1 * time.Second)
				continue
			}
			cnt = max(1, cnt-10)
			continue
		}
		if text == "1" {
			if _, err := c.Send(entity.NewTransfeData(enum.UserCountPacket, "", 0)); err != nil {
				models.Rconn <- true
				log.Println("断线重连", err)
				fmt.Println("连接已断开，正在尝试重连...")
				time.Sleep(1 * time.Second)
				return err
			}
			userCount := <-c.CmdChan
			var num int
			if userCount.Cmd == enum.UserCountPacket {
				num = userCount.Data.(int)
			} else {
				log.Println("获取玩家人数错误")
				fmt.Println("获取玩家人数错误！")
				time.Sleep(1 * time.Second)
				return &myerror.UserCountError{}
			}
			if cnt+10 > num {
				fmt.Println("这是最后一页！")
				time.Sleep(1 * time.Second)
				continue
			}
			cnt = min(num, cnt+10)
			continue
		}
		if text == "2" {
			if err := Search(c); err != nil {
				continue
			}
			continue
		}
		if text == "3" {
			log.Println("退出排行榜")
			fmt.Println("退出成功！")
			time.Sleep(1 * time.Second)
			return nil
		}
	}
}
