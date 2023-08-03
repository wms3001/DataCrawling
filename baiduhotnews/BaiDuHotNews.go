package baiduhotnews

import (
	"context"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

type BaiDuHotNews struct {
	News string `json:"news"`
	Url  string `json:"url"`
}

func (b *BaiDuHotNews) GetBaiDuHostNews() string {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://top.baidu.com/board?tab=realtime`),
		//chromedp.Text(`.Logo`, &res, chromedp.NodeVisible),
		chromedp.InnerHTML(`.container-bg_lQ801`, &res, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		log.Fatal(err)
	}
	var news []BaiDuHotNews
	doc.Find(".content_1YWBm").Each(func(i int, s *goquery.Selection) {
		new := BaiDuHotNews{}
		new.News = strings.ReplaceAll(s.Find("a").Text(), "查看更多>", "")
		new.Url, _ = s.Find("a").Attr("href")
		news = append(news, new)
	})
	nn, _ := json.Marshal(news)
	return string(nn)
}
