package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/xuri/excelize/v2"
	"lczx/utility/logger"
	"log"
	"net"
	"os/exec"
)

// GetClientIp 获取客户端IP
func GetClientIp(req *ghttp.Request) (ip string) {
	ip = req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = req.GetClientIp()
	}
	return
}

// GetLocalIp 获取服务端IP
func GetLocalIp() (ip string, err error) {
	var addrs []net.Addr
	addrs, err = net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		ip = ipAddr.IP.String()
		return
	}
	return
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ctx context.Context, ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	byteList := g.Client().GetBytes(ctx, url)
	src := string(byteList)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		logger.Error(ctx, "GetCityByIp Error: ", err.Error())
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		return city
	} else {
		return ""
	}
}

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

// Reverse 反转切片
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 执行shell命令
func interactiveToexec(commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)
	log.Println("interactiveToexec cmd: ", cmd)
	_, err := cmd.Output()
	cmd.Stderr = bytes.NewBuffer(nil)
	return err
}

// Doexec 直接通过字符串执行shell命令，不拼接命令
func Doexec(cmdStr string) error {
	cmd := exec.Command("bash", "-c", cmdStr)
	log.Println("Doexec cmd: ", cmd)
	_, err := cmd.CombinedOutput()
	return err
}

// RedisKey 组装redis key
func RedisKey(keys ...string) string {
	return gstr.Join(keys, ":")
}

// GetExcelHeadStyle 获取Excel文件表头样式
func GetExcelHeadStyle(file *excelize.File) (style int, err error) {
	style, err = file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#DFEBF6"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "微软雅黑",
			Size:   14,
			Color:  "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	return
}

// GetExcelBodyStyle 获取Excel文件表体样式
func GetExcelBodyStyle(file *excelize.File) (style int, err error) {
	style, err = file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#DFEBF6"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Family: "微软雅黑",
			Size:   11,
			Color:  "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	return
}
