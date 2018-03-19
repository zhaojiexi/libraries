package main

import (
	"github.com/go-gomail/gomail"
)

func main() {

	m := gomail.NewMessage()

	//发件人、姓名
	m.SetAddressHeader("From", "zhaojiexi@51yuyou.com", " rrrrrr")

	//设置收件人
	m.SetHeader("To",
		m.FormatAddress("8934120@qq.com", "zhxxxxx"))
	m.SetHeader("Subject", "test mail")
	//设置邮件内容格式 以及内容
	m.SetBody("text/html", "<a href=\"http://www.baidu.com\">你最宝贵de </a>")

	// 添加附件内容
	m.Attach("d://TIM截图20180319150352.png")

	d := gomail.NewDialer("smtp.ym.163.com", 465, "zhaojiexi@51yuyou.com", "*************!@#")

	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
