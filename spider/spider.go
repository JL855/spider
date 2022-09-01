package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tealeg/xlsx"
	"main.go/data"
	"net/http"
)

func Spider(page string, sheet *xlsx.Sheet) {
	//1.发送请求
	client:=http.Client{}
	req,err:=http.NewRequest("GET","https://www.aquanliang.com/blog/page/"+page,nil)
	if err!=nil {
		fmt.Println("req err",err)
	}
	//防止服务器检测爬虫访问，加了一些请求头伪造成浏览器访问
	req.Header.Set("Connection","keep-alive")
	req.Header.Set("cache-control","private, no-cache, no-store, max-age=0, must-revalidate")
	req.Header.Set("sec-ch-ua-mobile","?0")
	req.Header.Set("upgrade-insecure-requests","1")
	req.Header.Set("user-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70")
	req.Header.Set("accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site","none")
	req.Header.Set("sec-fetch-mode","navigate")
	req.Header.Set("sec-fetch-user","?1")
	req.Header.Set("sec-fetch-dest","document")
	req.Header.Set("accept-language","zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	resp,err:=client.Do(req)
	if err!=nil {
		fmt.Println("请求失败",err)
	}
	defer resp.Body.Close()

	//2.解析网页
	docDetail,err := goquery.NewDocumentFromReader(resp.Body)
	if err!=nil {
		fmt.Println("解析失败",err)
	}

	//3.获取节点信息

	//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span:nth-child(1)
	//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span:nth-child(2)
	//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span

	docDetail.Find("#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span").
		Each(func(i int, s *goquery.Selection) {  //在列表继续找
			var data data.QlData
			//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span:nth-child(1) > div > div > a > div
			title := s.Find("div > div > a > div").Text()
			//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span:nth-child(1) > div > a > div > div._2ahG-zumH-g0nsl6xhsF0s > div > img
			img := s.Find("div > a > div > div._2ahG-zumH-g0nsl6xhsF0s > div > img")  //img标签
			imgTmp, ok := img.Attr("src")


			//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span:nth-child(1) > div > div > div._1nlYtcrR408yNacE0R0s3M > div._2gvAnxa4Xc7IT14d5w8MI1
			info := s.Find("div > div > div._1nlYtcrR408yNacE0R0s3M > div._2gvAnxa4Xc7IT14d5w8MI1").Text()
			bt:=[]byte(info)
			bt=bt[len(bt)-3:]   //视情况而定，可以遍历一遍找到最后的数字就是阅读数
			readNum:=string(bt)

			//#__next > div > div:nth-child(2) > section > main > section > div:nth-child(2) > div:nth-child(1) > div._1ySUUwWwmubujD8B44ZDzy > span:nth-child(1) > div > div > div._1nlYtcrR408yNacE0R0s3M > div._3TzAhzBA-XQQruZs-bwWjE
			info1 := s.Find("div > div > div._1nlYtcrR408yNacE0R0s3M > div._3TzAhzBA-XQQruZs-bwWjE").Text()
			bt=[]byte(info1)
			bt=bt[len(bt)-10:]
			date := string(bt)
			fmt.Printf("《%s%s\n", title, "》")

			if ok {
				data.Title = title
				data.Date = date
				data.ReadNum = readNum
				data.Picture = imgTmp
				fmt.Println("data结构", data)
				//添加一行
				row1 := sheet.AddRow()
				//添加一格
				c1 := row1.AddCell()
				c1.Value = data.Title
				c2 := row1.AddCell()
				c2.Value = data.Date
				c3 := row1.AddCell()
				c3.Value = data.ReadNum
				c4 := row1.AddCell()
				c4.Value = data.Picture  //封面图没弄好
			}
		})

	// 保存信息到mysql
	//
}
