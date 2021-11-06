package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	"github.com/spf13/cobra"
	"github.com/trelore/iris-classification/cmd/train/datasets"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your application",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	f, err := datasets.Data.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := f.(io.ReadSeeker); !ok {
		log.Fatal("file is not a ReadSeeker")
	}

	rawData, err := base.ParseCSVToInstancesFromReader(f.(io.ReadSeeker), false)
	if err != nil {
		log.Fatal(err)
	}

	//Initialises a new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	//Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		log.Fatal(err)
	}

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))

	cls.Save("model.cls")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
