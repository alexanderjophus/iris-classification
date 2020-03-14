package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "predict",
	Short: "A cmd tool to predict IRIS class based off stored model",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	cls, err := knn.ReloadKNNClassifier("model.cls")
	if err != nil {
		log.Fatal(err)
	}
	toPredict := strings.NewReader("5.1,3.5,1.4,0.2")
	rawData, err := base.ParseCSVToInstancesFromReader(toPredict, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cls.TrainingData.AllAttributes())
	fmt.Println(rawData.AllAttributes())
	predictions, err := cls.Predict(rawData)
	if err != nil {
		log.Fatal(err)
	}

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(rawData, predictions)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
