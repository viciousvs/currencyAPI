package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/viciousvs/currencyAPI/config"
	"github.com/viciousvs/currencyAPI/internal/model"
	"github.com/viciousvs/currencyAPI/internal/storage"
)

func cmd(cfg *config.Config) {
	ctx := context.Background()
	db := storage.NewPostgresDB(cfg.PostgresConfig)
	res, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		log.Printf("cannot req to address %s\n", err)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("cannot readall %s\n", err)
	}
	c := model.Currencies{}
	err = json.Unmarshal(resp, &c)
	fmt.Println(c)
	stmt := "INSERT INTO currency (date, time_stamp, base, rate, value) VALUES ($1, $2, $3, $4, $5)"
	count := 0
	for key, val := range c.Rates {
		_, err := db.DB.ExecContext(ctx, stmt, c.Date, c.Timestamp, c.Base, key, val)
		if err != nil {
			log.Printf("cannot insert %s\n", err)
		}
		count++
	}
}
func main() {

	cfg := config.NewConfig()
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("cannot run cron %s", err.Error())
	}
	j, err := s.NewJob(
		gocron.DurationJob(
			time.Duration(cfg.CronConfig.Spec)*time.Hour),
		gocron.NewTask(
			cmd,
			cfg,
		),
	)
	if err != nil {
		log.Fatalf("cannot run job %s", err.Error())
	}
	log.Printf("runned job: %v", j.ID())

	s.Start()
	select {
	case <-time.After(time.Hour):
	}
	log.Fatal(s.Shutdown())
}
