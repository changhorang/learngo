package scrapper // src에서 main.go로 실행 가능
import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id            string
	title         string
	job_condition string
	job_date      string
	job_sector    string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://saramin.co.kr/zf_user/search/recruit?searchword=" + term
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 1; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 1; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("DONE JOB Scrapper!")
}

func getPage(page int, url string, mainC chan []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := url + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting ", pageURL)
	res, err := http.Get(pageURL)

	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searcItems := doc.Find(".item_recruit")

	searcItems.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searcItems.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

// CleanString clearns a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".job_tit>a").Text())
	job_condition := CleanString(card.Find(".job_condition").Text())
	job_date := CleanString(card.Find(".job_date").Text())
	job_sector := CleanString(card.Find(".job_sector").Text())

	c <- extractedJob{
		id:            id,
		title:         title,
		job_condition: job_condition,
		job_date:      job_date,
		job_sector:    job_sector,
	}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID (https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx={id})", "Title", "Job_condition", "Job_date", "Job_sector"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.job_condition, job.job_date, job.job_sector}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}
