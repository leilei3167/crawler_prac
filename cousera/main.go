package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Course stores information about a coursera course
type Course struct {
	Title       string
	Description string
	Creator     string

	URL string

	Rating string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("coursera.org", "www.coursera.org"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./coursera_cache"), colly.MaxDepth(6),
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{RandomDelay: time.Second * 3, Parallelism: 3})
	// Create another collector to scrape course details
	detailCollector := c.Clone()

	courses := make([]Course, 0, 200)

	// 选择出具有href属性的所有a元素,每一个元素被找到时执行此回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//a包含href但是同时具有class属性为这个值的排除
		if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}
		link := e.Attr("href")
		// 如果该连接包含登录,等无关前缀,直接跳过
		if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}
		// 继续探测
		e.Request.Visit(link)
	})

	// 当具备CardText-link 属性的元素被找到时,使用细节收集者进行挖掘
	c.OnHTML(`.CardText-link`, func(e *colly.HTMLElement) {
		courseURL := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Println("get url:", courseURL)
		if strings.Index(courseURL, "/learn") != -1 {
			detailCollector.Visit(courseURL) //调用另一个细节收集者(使用绝对URL,因为跨越收集者的话url不是完整的)
		}
	})

	// Extract details of the course
	detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
		log.Println("Course found", e.Request.URL)
		title := e.ChildText("h1") //ChildText会将所有选择结果所有的Text()去空格后返回
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		course := Course{
			Title:       title,
			URL:         e.Request.URL.String(),
			Description: e.ChildText("div.content"), //div元素中,class值为content的结果
			Creator:     e.ChildText(".instructor-name"),
			Rating:      e.ChildText("div.rc-ReviewsOverview__totals__rating"),
		}

		courses = append(courses, course)
	})

	// Start scraping on http://coursera.com/browse
	c.Visit("https://coursera.org/browse")
	c.Wait()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(courses)
}
