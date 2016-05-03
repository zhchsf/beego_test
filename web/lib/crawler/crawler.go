package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  // "github.com/djimenez/iconv-go"
  "net/http"
  // "io/ioutil"
)

func craw() {
  doc, err := goquery.NewDocument("http://www.caishuo.com/topics") 
  if err != nil {
    fmt.Println(err)
  }
  doc.Find("#topics_list dt a em").Each(func(i int, s *goquery.Selection) {
    title := s.Text()
    fmt.Println(title)
  })
}

func crawFromReader() {
  res, _ := http.Get("http://www.caishuo.com/topics")
  defer res.Body.Close()
  // body, _ := ioutil.ReadAll(res.Body)
  // fmt.Println(body)
  // utfBody, _ := iconv.NewReader(res.Body, "utf-8", "utf-8")
  doc, _ := goquery.NewDocumentFromReader(res.Body)
  doc.Find("#topics_list dt a em").Each(func(i int, s *goquery.Selection) {
    title := s.Text()
    fmt.Println(title)
  })
}

func main() {
  craw()

  fmt.Println("*********************")

  crawFromReader()
}
