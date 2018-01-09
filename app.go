package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/julienschmidt/httprouter"

	"github.com/tokopedia/sqlt"

	"github.com/antony/polling/config"
	"github.com/google/gops/agent"

	"github.com/tokopedia/logging/tracer"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"

	_pdaRepository "github.com/antony/polling/polling_defined_answer/repository"
	_puaRepository "github.com/antony/polling/polling_user_answer/repository"

	pollingDelivery "github.com/antony/polling/polling/delivery/http"
	pollingRepository "github.com/antony/polling/polling/repository"
	pollingUsecase "github.com/antony/polling/polling/usecase"

	_ "github.com/lib/pq"
)

var cfg *config.Config

func init() {
	var ok bool
	cfg, ok = config.NewConfig([]string{"./files/etc/polling"}...)
	if !ok {
		fmt.Println("Error opening config files")
	}
}

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	if err := agent.Listen(&agent.Options{}); err != nil {
		log.Fatal(err)
	}

	//Open connection to db
	db, err := sqlt.Open("postgres", cfg.Polling.DB)
	if err != nil {
		log.Println("Error opening db")
		return
	}

	router := httprouter.New()

	//init polling handler
	pp := pollingRepository.NewPostgresPollingRepository(db)
	pdap := _pdaRepository.NewPostgresPollingDefinedAnswerRepository(db)
	puap := _puaRepository.NewPollingUserAnswerPostgresRepository(db)
	pu := pollingUsecase.NewPollingUsecase(pp, pdap, puap)
	pollingDelivery.NewPollingHttpHandler(router, pu)

	go logging.StatsLog()

	tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	log.Fatal(grace.Serve(":9000", router))
}
