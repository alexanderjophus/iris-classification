package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	irisclassificationpbconnect "github.com/trelore/iris-classification/proto/gen/go/iris_classification/v1/irisclassificationpbconnect"
	"github.com/trelore/iris-classification/svc/server"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "svc",
	Short: "A service to predict iris classifications",
	Run:   run,
}

const (
	serviceName = "iris-classification"
)

func run(cmd *cobra.Command, args []string) {
	s := server.New()

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer zapLogger.Sync()

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollections(collectors.GoRuntimeMemStatsCollection | collectors.GoRuntimeMetricsCollection),
	))
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "calls_to_iris_classification",
		Help: "the number of calls to get an iris classification",
	}, []string{"rpc"})
	reg.Register(counter)

	interceptors := connect.WithInterceptors(
		NewPromInterceptor(counter),
		NewLoggingInterceptor(zapLogger),
	)

	go http.ListenAndServe(":2112", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	mux := http.NewServeMux()
	mux.Handle(irisclassificationpbconnect.NewIrisClassificationServiceHandler(
		s,
		interceptors,
	))
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
