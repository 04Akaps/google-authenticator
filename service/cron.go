package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"demo-scrapping/config"
	"demo-scrapping/repository"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
)

type cronJob struct {
	cfg        *config.Config
	repository repository.RepositoryImpl
	c          *cron.Cron
}

func NewCronJob(cfg *config.Config, repository repository.RepositoryImpl) *cronJob {
	c := &cronJob{cfg: cfg, repository: repository, c: cron.New()}

	go c.runJobs()

	return c
}

func (j *cronJob) runJobs() {
	c := j.c

	c.AddFunc("*/5 * * * * *", func() {
		j.scrapping()
		fmt.Println()
	})

	c.Start()
	defer c.Stop()

	select {}
}

func (j *cronJob) scrapping() error {
	log.Println("five second job excuted from mysql for Scrapping")

	if allResult, err := j.repository.ViewAll(); err != nil {
		return err
	} else if len(allResult) == 0 {
		return errors.New("all Result zero")
	} else {
		for _, r := range allResult {
			log.Printf("Try Scrapping URL :  %s", r.URL)
			log.Printf("Try Scrapping CardSelect : %s", r.CardSelector)
			log.Printf("Try Scrapping Tag :  %s", r.Tag)
			fmt.Println()

			j.scrappingHTML(r.URL, r.CardSelector, r.InnerSelector, strings.Split(r.Tag, " "))
		}

		return nil
	}
}

func (j *cronJob) scrappingHTML(url, cardSelector, innerSelect string, tag []string) {
	client := http.Client{Timeout: time.Second * 10}

	if request, err := http.NewRequest("GET", url, nil); err != nil {
		log.Println("Failed To Make Request", "err", err)
	} else {
		request.Header.Set("User-Agent", "M")

		if response, err := client.Do(request); err != nil {
			log.Println("Failed To Call GET API", "err", err)
		} else {
			defer response.Body.Close()

			if doc, err := goquery.NewDocumentFromReader(response.Body); err != nil {
				log.Println("Failed To Read response", "err", err)
			} else {
				searchCard := doc.Find(cardSelector)

				if searchCard.Length() == 0 {
					log.Println("Failed To Search CardSelector")
				} else {
					searchCard.Each(func(_ int, card *goquery.Selection) {
						card.Find(innerSelect).Each(func(_ int, child *goquery.Selection) {
							for _, t := range tag {
								d := child.Find(t).Text()
								log.Println(d)
							}
						})
					})
				}
			}
		}
	}
}
