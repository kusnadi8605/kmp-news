# kmp-news

# kmp-news-producer
Golang &amp; Kafka

## add lib
go get -v github.com/go-sql-driver/mysql
go get -v gopkg.in/olivere/elastic.v7
go get -v github.com/segmentio/kafka-go
go get -v kmp-news-consumer/parser
go get -v github.com/go-yaml/yaml

# running kafka
## running zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
## running kafka server
bin/kafka-server-start.sh config/server.properties

# Running App
go run main.go

## Get News
curl -X GET 'http://localhost:3000/api/news?page=1' 

curl -X GET 'http://localhost:3000/api/news?page=2' 
