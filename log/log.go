package log

import (
	"github.com/godcong/elogrus"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

// InitLog ...
//docker run -p 5601:5601 -p 9200:9200 -p 5044:5044 -itd -e LOGSTASH_START=0 -e KIBANA_START=0 --name elk sebp/elk
func InitLog(index string) {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}

	t, err := elogrus.NewElasticHook(client, "localhost", log.TraceLevel, index)
	if err != nil {
		log.Panic(err)
	}
	log.AddHook(t)
}
