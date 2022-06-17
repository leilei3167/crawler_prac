package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

/*
爬取豆瓣top250电影

重点:
1.模拟登录,未登录状态会很快被判定为异常
输入错误密码,在浏览器控制台中定位到提交账号密码的url,如豆瓣就是 https://accounts.douban.com/j/mobile/login/basic



2.用goquery抓取需要的元素 及元素值

3.分词,去除无用的数据
*/
func GetCollector() *colly.Collector {
	/* 配置collector
	限制访问超时,模拟随机延迟等
	*/
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout: time.Second * 3, KeepAlive: time.Second * 10,
		}).DialContext,
	}
	c.WithTransport(tr)
	c.Limit(&colly.LimitRule{RandomDelay: time.Second * 3})

	/*
		//如果实现登录?某些有滑动验证码
		err := c.Post("https://accounts.douban.com/j/mobile/login/basic", map[string]string{
			"name":     "17708037113",
			"password": "a123456789",
		})
		if err != nil {
			log.Fatal(err)
		} */
	jar, err := cookiejar.New(&cookiejar.Options{}) //自动管理获得的cookie
	if err != nil {
		log.Fatal(err)
	}
	c.SetCookieJar(jar)
	return c
}

func main() {
	c := GetCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("开始访问:%s\n", r.URL.String())
	})
	//抓取信息(元素选择器和class选择器结合使用,ol元素中,class为grid_view的结果)
	//等效于ol[class=grid_view]
	c.OnHTML("ol.grid_view", func(h *colly.HTMLElement) {
		dom := h.DOM

		dom.Find("li").Each(func(i int, s *goquery.Selection) {
			//TODO:抓取到的数据包含大量无效的字符 空行 换行等需要对数据进行处理
			//编号
			fullurl, _ := s.Find(".hd a").Attr("href")
			fmt.Println("编号:", filepath.Base(fullurl))
			//名称
			fmt.Println("名称:", s.Find(".title").Text())
			//评分
			fmt.Println("评分:", s.Find("span.rating_num").Text())
			//信息

			fmt.Println("信息:", strings.TrimSpace(s.Find("p").Text()))
		})

	})

	//实现翻页(.后面跟class的值,实现整个页面中class值为paginator的结果)
	c.OnHTML(".paginator", func(h *colly.HTMLElement) {
		h.DOM.Find("a").Each(func(i int, s *goquery.Selection) {
			uurl, ok := s.Attr("href") //attr用于获取元素的值(会首先判断该元素是否有值)
			if !ok {
				fmt.Println("获取下一页失败")
				return
			} else {
				h.Request.Visit(uurl)
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("code:%d,msg:%v\n", r.StatusCode, err)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("抓取完成")
	})

	err := c.Visit("https://movie.douban.com/top250")
	if err != nil {
		log.Fatal(err)
	}
}
