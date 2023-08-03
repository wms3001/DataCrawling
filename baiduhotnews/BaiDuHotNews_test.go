package baiduhotnews

import (
	"log"
	"testing"
)

func TestBaiDuHotNews_GetBaiDuHostNews(t *testing.T) {
	var bn BaiDuHotNews
	res := bn.GetBaiDuHostNews()
	log.Println(res)
}
