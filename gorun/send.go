package main

import (
   "fmt"
   "gopkg.in/gomail.v2"
)

func main() {
   m := gomail.NewMessage()

   //发送人
   m.SetHeader("From", "573916016@qq.com")
   //接收人
   m.SetHeader("To", "573916016@qq.com")
   //抄送人
   //m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
   //主题
   m.SetHeader("Subject", "小佩奇")
   //内容
   m.SetBody("text/html", "<h1>新年快乐</h1>")
   //附件
   //m.Attach("./myIpPic.png")

   //拿到token，并进行连接,第4个参数是填授权码
   d := gomail.NewDialer("smtp.qq.com", 587, "573916016@qq.com", "ioamoihljfavbedc")

   // 发送邮件
   if err := d.DialAndSend(m); err != nil {
      fmt.Printf("DialAndSend err %v:", err)
      panic(err)
   }
   fmt.Printf("send mail success\n")
}