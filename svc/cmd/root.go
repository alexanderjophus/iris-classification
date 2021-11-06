package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
	pb "github.com/trelore/iris-classification/proto/gen/go/iris_classification/v1"
	"github.com/trelore/iris-classification/svc/server"
	"google.golang.org/grpc"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "svc",
	Short: "A service to predict iris classifications",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	s := server.New()
	grpcS := grpc.NewServer()
	defer grpcS.GracefulStop()

	pb.RegisterIrisClassificationServiceServer(grpcS, &s)

	address := ":32400"
	log.Printf("listening to address %s", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	grpcS.Serve(listener)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
