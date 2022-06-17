package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

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

	jar, err := cookiejar.New(&cookiejar.Options{}) //自动管理获得的cookie
	if err != nil {
		log.Fatal(err)
	}
	c.SetCookieJar(jar)
	return c
}
func main() {
	c := GetCollector()

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("done!")
	})



	
}
