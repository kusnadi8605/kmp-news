package handler

import (
	"encoding/json"
	conf "kmp-news/config"
	dts "kmp-news/datastruct"
	mdl "kmp-news/models"
	"net/http"

	"gopkg.in/olivere/elastic.v7"
)

//GetNewsHandler ..
func GetNewsHandler(conn *conf.Connection, client *elastic.Client, index string, perpage int) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var NewsResponse dts.NewsResponse

		//listNews, err := mdl.GetNews(conn)

		page := req.URL.Query().Get("page")

		res, err := mdl.GetNewsElastic(conn, client, index, page, perpage)

		if err != nil {
			NewsResponse.ResponseCode = "301"
			NewsResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(NewsResponse)

			conf.Logf("Response News : %v", NewsResponse.ResponseDesc)

			return
		}

		if len(res) < 1 {
			NewsResponse.ResponseCode = "301"
			NewsResponse.ResponseDesc = "data not found"
			json.NewEncoder(w).Encode(NewsResponse)

			conf.Logf("Response News : %v", NewsResponse.ResponseDesc)

			return
		}

		NewsResponse.ResponseCode = "000"
		NewsResponse.ResponseDesc = "Success"
		NewsResponse.Payload = res
		json.NewEncoder(w).Encode(NewsResponse)

		conf.Logf("Response News : %v", NewsResponse.ResponseDesc)
	}

}
