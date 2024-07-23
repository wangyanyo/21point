package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/wangyanyo/21point/common/enum"
)

func Cle() {
	cmd := &exec.Cmd{}
	if enum.SYS_TYPE == "windows" {
		cmd = exec.Command("cmd.exe", "/c", "cls")
	}
	if enum.SYS_TYPE == "linux" {
		cmd = exec.Command("sh", "-c", "clear")
		fmt.Println(cmd)
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func PrintMessage(text string) {
	log.Println(text)
	fmt.Println(text)
	time.Sleep(1 * time.Second)
}

func CalcPoint(cards []string) int {
	res, cnt := 0, 0
	for _, s := range cards {
		if s == "A" {
			res++
			cnt++
		} else if s == "10" || s == "J" || s == "Q" || s == "K" {
			res += 10
		} else {
			res += int(s[0] - '0')
		}
	}
	if cnt > 0 && res+10 <= 21 {
		res += 10
	}
	return res
}

func GetOpt(msg string, x int) string {
	var opt string
	for {
		fmt.Print(msg)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		opt = scanner.Text()
		flag := false
		for i := 0; i <= x; i++ {
			if opt == string('0'+i) {
				flag = true
			}
		}
		if flag {
			break
		}
		fmt.Printf("\033[1A\r")
		fmt.Printf("%*c", 40, ' ')
		fmt.Printf("\r")
	}
	return opt
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CheckGameResult(x int, y int) int {
	if x > 21 && y > 21 {
		return 0
	}
	if x > 21 {
		return -1
	}
	if y > 21 {
		return 1
	}
	t1, t2 := 21-x, 21-y
	if t1 < t2 {
		return 1
	}
	if t1 > t2 {
		return -1
	}
	return 0
}

func SetPos(row int, cls int) {
	fmt.Printf("\033[%d;%dH", row, cls)
}

var stripAnsiEscapeRegexp = regexp.MustCompile(`(\x9B|\x1B\[)[0-?]*[ -/]*[@-~]`)

func stripAnsiEscape(s string) string {
	return stripAnsiEscapeRegexp.ReplaceAllString(s, "")
}

func RealLength(s string) int {
	return runewidth.StringWidth(stripAnsiEscape(s))
}
