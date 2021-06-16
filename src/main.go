package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	data "cryptoserver-clean-app/data"
	"cryptoserver-clean-app/middleware"
	"cryptoserver-clean-app/service"
	"cryptoserver-clean-app/transport"
	"cryptoserver-clean-app/util"
	"cryptoserver-clean-app/hitbtc"
	gklog "github.com/go-kit/kit/log"
)

const (
	application = "cryptoserver-clean"
	logfile     = "/appl/cryptoserver-clean/api.log"
)

func main() {

	/* Parse command line */
	configFile := flag.String("conf", "", "Configuration File")
	flag.Parse()

	if len(*configFile) == 0 {
		println("Conf file name required\n")
		os.Exit(1)
	}

	/* Parse configuration file to config */
	config := data.NewConfig()
	if err := config.Init(*configFile); err != nil {
		println("Cannot initialize config from file:", *configFile)
		println("Error:", err)
		os.Exit(1)
	}

	log := util.NewCustomLogFormatter().GetLogger(os.Stderr, "cryptoserver", config.Logging.Level)

	/* Create service instance */
	var svc service.Service
	svc = service.NewCryptoServerSvc(config)

	/* Plug logging middleware */
	svc = middleware.LoggingMiddleware(log)(svc)


	/* Initialize go-cache */
	hitbtc.InitCache()

	/* Start Websocker readers */
	clientETH := hitbtc.NewHitBTCClient(config, "ETHBTC")
	go clientETH.Start(log)

	clientBTC := hitbtc.NewHitBTCClient(config, "BTCUSD")
	go clientBTC.Start(log)


	/* Listen and serve HTTP requests */
	url, err := url.Parse(config.Server.ListenAddress)
	if err != nil {
		log.Error("Cannot parse listen address:", config.Server.ListenAddress, err.Error())
		return
	}

	host, port, _ := net.SplitHostPort(url.Host)
	if host == "0.0.0.0" {
		host = ""
	}

	listenAddress := fmt.Sprintf("%s:%s", host, port)
	var httpHandler http.Handler
	{
		httpHandler = transport.MakeHTTPHandler(svc, gklog.NewLogfmtLogger(os.Stderr))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Info("HTTP transport listening on port:", listenAddress)
		errs <- http.ListenAndServe(listenAddress, httpHandler)
	}()

	log.Info("Exiting", <-errs)
}
