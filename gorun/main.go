package main

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	
	"strings"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	// 创建邮件消息1
	m := gomail.NewMessage()

	// 设置发送人
	m.SetHeader("From", "573916016@qq.com")
	// 设置接收人
	m.SetHeader("To", "573916016@qq.com")
	// 设置主题
	m.SetHeader("Subject", "推特信息")

	// 创建邮件正文内容
	var body strings.Builder

	// 抓取推特信息并保存在文件中
	scraper := twitterscraper.New()

	usernames := map[string]string{
		"ylecun":          "Yann LeCun (NYU教授、深度学习创始人)",
		"elonmusk":        "Elon Musk (特斯拉、SpaceX 创始人)",
		"BillGates":       "Bill Gates (微软创始人)",
		"AndrewYNg":       "吴恩达 (Coursera联合创始人、前百度首席科学家)",
		"karpathy":        "Andrej Karpathy (前Tesla AI 负责人)",
		"pmddomingos":     "Pedro Domingos (华盛顿大学教授、机器学习专家)",
		"GaryMarcus":      "Gary Marcus (纽约大学教授、AI 评论家)",
		"jackclarkSF":     "Jack Clark (OpenAI通讯主任)",
		"geoffreyhinton":  "Geoffrey Hinton (多伦多大学教授、深度学习先驱)",
		"ilyasut":         "Ilya Sutskever (OpenAI联合创始人、深度学习专家)",
		"JeffDean":        "Jeff Dean (Google人工智能部门负责人)",
		"OriolVinyalsML":  "Oriol Vinyals (Google Brain研究员、AlphaGo团队成员)",
		"sama":            "Sam Altman (OpenAI 创始人)",
		"gdb":             "Greg Brockman (OpenAI 联合创始人兼主席)",
		"lexfridman":      "Lex Fridman (MIT研究员、AI 评论家)",
		"jbrowder1":       "Joshua Browder (DoNotPay创始人、AI 律师)",
		"mattshumer_":     "Matt Shumer (SociallyMined创始人、AI 创业者)",
		"gregisenberg":    "Greg Isenberg (Late Checkout创始人、AI 创业者)",
		"SullyOmarr":      "Sully Omarr (AI 研究员、OpenAI 成员)",
		"ESYudkowsky":     "Eliezer Yudkowsky (人工智能风险研究员、机器超级智能理论家)",
		 "tegmark":         "Max Tegmark (MIT教授、宇宙学家)",
		 "Google":          "Google (全球最大的搜索引擎和互联网技术公司)",
		 "DeepMind":        "DeepMind (Google旗下的人工智能研究公司)",
		 "OpenAI":          "OpenAI (GPT、DALL-E 缔造者)",
	}

	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	for username, name := range usernames {
		var tweetCount int
		var userTweets []string
		for tweet := range scraper.GetTweets(context.Background(), username, 3) {
			if tweet.Error != nil {
				panic(tweet.Error)
			}
			if strings.HasPrefix(tweet.Text, "RT @") {
				continue
			}
			createdAt := time.Unix(tweet.Timestamp, 0)
			if !strings.HasPrefix(createdAt.Format("2006-01-02"), yesterday) &&
				!strings.HasPrefix(createdAt.Format("2006-01-02"), today) {
				continue
			}
			userTweets = append(userTweets, fmt.Sprintf("%d. %s -- %s", tweetCount+1, tweet.Text, createdAt.Format("2006-01-02")))
			tweetCount++
		}
		if tweetCount > 0 {
			body.WriteString(fmt.Sprintf("%s:\n", name))
			for _, tweet := range userTweets {
				body.WriteString(fmt.Sprintf("%s\n", tweet))
			}
			body.WriteString("\n")
		}
	}

	// 设置邮件正文内容
	m.SetBody("text/plain", body.String())

	// 拿到token，并进行连接，第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "573916016@qq.com", "ioamoihljfavbedc")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}

	fmt.Printf("send mail success\n")
	fmt.Println("Done!")
}
