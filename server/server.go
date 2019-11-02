package server

import (
	"github.com/fabiosebastiano/good-news-backend/crawler"
	"github.com/fabiosebastiano/good-news-backend/db"
	"github.com/fabiosebastiano/good-news-backend/utils"
)

// Init <function>
// inizializzazione del server
func Init() {
	utils.InitEnvVars()

	db.InitService()

	db.GetClient().FillSeedsInformation() // inizializza DB

	// workers
	crawler.Start()

	r := NewRouter()
	r.Run(":80")
}
