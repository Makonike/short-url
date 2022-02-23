package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"math"
	"strings"
	"time"
)

type Short struct {
	Id        int64  `orm:"column(id)"`
	Lurl      string `orm:"column(lurl)"`
	Md5       string `orm:"column(md5)"`
	GmtCreate int64  `orm:"column(gmt_create)"`
}

var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// ToShort 将长链转换为短链，并存储于数据库中
func ToShort(url string) (shortUrl string) {
	var short Short
	// 获取分号器id(短链id)
	id := GetTicket()
	// 将长链md5压缩，利于索引搜索
	h := md5.New()
	_, err := io.WriteString(h, url)
	if err != nil {
		logs.Error("write md5 string io error")
		return ""
	}
	short.Lurl = url
	short.Md5 = fmt.Sprintf("%x", h.Sum(nil))
	short.GmtCreate = time.Now().Unix()
	short.Id = id
	//logs.Info(short)
	// get short url
	shortUrl = changeShort(id)
	//logs.Info(shortUrl)
	// save information of short url
	o := orm.NewOrm()
	_, err2 := o.Insert(&short)
	if err2 != nil {
		logs.Error("save information of short url error")
		return ""
	}
	return shortUrl
}

// 将短链id转换为短链（10进制转62进制）
func changeShort(id int64) (url string) {
	var bytes []byte
	for id > 0 {
		bytes = append(bytes, chars[id%62])
		id = id / 62
	}
	reverse(bytes)
	return string(bytes)
}

// 反转字符串
func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

// 62进制转10进制
func decode(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

// GetShort 根据短链(62)获取长链
func GetShort(url string) (longUrl string) {
	var short Short
	o := orm.NewOrm()
	err := o.QueryTable("t_short").Filter("id", decode(url)).One(&short)
	if err != nil {
		return ""
	}
	return short.Lurl
}
