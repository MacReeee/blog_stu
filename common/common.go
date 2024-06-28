package common

import (
	"goblog/config"
	"goblog/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
