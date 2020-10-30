package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

//얼마나 많은 페이지가 있는지 알아내는 프로그램
func main() {
	var jobs []extractedJob
	totalPages := getPages()
	//fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50) //=page*50 	string이 아니라 숫자니까 go에 포함된 패키지를 이용해 형변환해서 넣는다 생각 (int to asci/Itoa)
	fmt.Println("Requestring", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		// id, _ := card.Attr("data-jk")
		// //fmt.Println(id)
		// title := cleanString(card.Find(".title>a").Text())
		// //fmt.Println(title)
		// location := cleanString(card.Find(".sjcl").Text())
		// fmt.Println(id, title, location)
		job := extractJob(card)
		jobs = append(jobs, job) //append를 이용하여 jobs라는 배열에 추출한 job을 업데이트함
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	//fmt.Println(id)
	title := cleanString(card.Find(".title>a").Text())
	//fmt.Println(title)
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ") //TrimSpace 스페이스를 지운다
	//모든공백제거후 택스트로만 된 배열을 만든것
	//Join은 배열을 가져와서 합치는 역할 (seperator 이용)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res) //response(응답)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// fmt.Println(doc)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil { //request(의뢰)에 에러가 있다면 이렇게!
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status :", res.StatusCode)
	}
}

//read code and study
