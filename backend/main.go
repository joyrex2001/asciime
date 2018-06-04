package backend

import (
	"fmt"
	"log"
	"net/http"
	"time"

	logger "github.com/gleicon/go-httplogger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// healthz will start listening for /healthz http requests on given address.
func healthz(addr string) {
	go func() {
		http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "{ status: 'OK', timestamp: %d }", time.Now().Unix())
		})
		log.Fatal(http.ListenAndServe(addr, nil))
	}()
}

// Main is the main entry point for starting this service, based the settings
// initiated by cmd.
func Main(cmd *cobra.Command, args []string) {
	hlth := viper.GetString("health.listen-addr")
	if hlth != "" {
		healthz(hlth)
	}

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
