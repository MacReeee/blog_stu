package common

import (
	"encoding/json"
	"goblog/config"
	"goblog/models"
	"io/ioutil"
	"net/http"
	"sync"
)
var Template models.HtmlTemplate

func Load()  {
	var err error
	wg := sync.WaitGroup{}
	//加载html模板
	wg.Add(1)
	go func() {
		Template, err = models.InitHtmlTemplate(config.Cfg.System.CurrentDir + "\\template")
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Error(w http.ResponseWriter,err error){
	var ret models.Result
	ret.Code = -999
	ret.Error = err.Error()
	r,_ := json.Marshal(ret)
	w.Header().Set("Content-Type","application/json")
	w.Write(r)
}

func ReturnSuccess(w http.ResponseWriter,result interface{})  {
	var ret models.Result
	ret.Code = 200
	ret.Data = result
	w.Header().Set("Content-Type","application/json")
	r,_ := json.Marshal(ret)
	w.Write(r)
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}