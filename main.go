package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type parametrsJson struct {
	FIO string
	Age int
	PassportData struct{
		Seria string
		Number string
}
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", Handler())
	managementServer := &http.Server{
		Addr: ":5505",
	}
	defer managementServer.Close()
	go func() {
		if err := managementServer.ListenAndServe(); err != nil {
			fmt.Errorf("Ошибка сервера управления: ", err)
		}
	}()

	log.Info("Приложение запущено")

	<-signals
}

func putUserData(w http.ResponseWriter, r *http.Request) {
	body,err:= ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var param parametrsJson

	err=json.Unmarshal(body,&param)
if err!=nil {
	log.Error(err)
	return
}

	log.Info("ФИО:",param.FIO)
	log.Info("Возраст:",param.Age)
	log.Info("Паспортные данные:")
	log.Info("Серия:",param.PassportData.Seria)
	log.Info("Номер:",param.PassportData.Number)
}


func Handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"метод": r.Method,
			"путь":  r.URL.Path,
		}).Info("Получен управляющий запрос")

		path := r.URL.Path
		switch path {
		case "/ping":
			log.Info("OK")
		case "/putUserData":
			putUserData(w,r)
		default:
			log.Info("метод не найден")
		}
	}
}