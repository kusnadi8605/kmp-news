package models

import (
	conf "kmp-news/config"
	dts "kmp-news/datastruct"
)

//GetNews ..
func GetNews(conn *conf.Connection) ([]dts.News, error) {
	arrNews := []dts.News{}
	strNews := dts.News{}
	strQuery := `select * from news`
	rows, err := conn.Query(strQuery)

	conf.Logf("Query News : %s", strQuery)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(
			&strNews.ID,
			&strNews.Author,
			&strNews.Body,
			&strNews.Created,
		)

		if err != nil {
			return nil, err
		}

		arrNews = append(arrNews, strNews)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return arrNews, nil
}

//GetNewsDetail ..
func GetNewsDetail(conn *conf.Connection, ID int64) ([]dts.NewsDetail, error) {
	arrNewsDetail := []dts.NewsDetail{}
	strNewsDetail := dts.NewsDetail{}

	strQuery := `select * from news where id=?`

	rows, err := conn.Query(strQuery, ID)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(&strNewsDetail.ID,
			&strNewsDetail.Author,
			&strNewsDetail.Body,
			&strNewsDetail.Created,
		)

		if err != nil {
			return nil, err
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		arrNewsDetail = append(arrNewsDetail, strNewsDetail)
	}

	return arrNewsDetail, nil
}
