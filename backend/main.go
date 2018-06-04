package backend

import (
	"log"
	"net/http"
	"time"

	logger "github.com/gleicon/go-httplogger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Main(cmd *cobra.Command, args []string) {
	hndlr := NewHandler()

	srv := http.Server{
		Addr:         viper.GetString("web.listen-addr"),
		Handler:      logger.HTTPLogger(hndlr),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	cert := viper.GetString("web.cert-file")
	key := viper.GetString("web.key-file")
	tls := viper.GetBool("web.enable-tls")

	if tls {
		log.Fatal(srv.ListenAndServeTLS(cert, key))
	} else {
		log.Fatal(srv.ListenAndServe())
	}
}
