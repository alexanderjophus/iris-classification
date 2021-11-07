// modified from
// https://github.com/gorgonia/gorgonia/blob/v0.9.17/examples/iris/cmd/main.go
package cmd

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/spf13/cobra"
	"github.com/trelore/iris-classification/cmd/predict/models"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "predict",
	Short: "A cmd tool to predict IRIS class based off stored model",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	b, err := models.Data.ReadFile("theta.bin")
	if err != nil {
		log.Fatal(err)
	}
	dec := gob.NewDecoder(bytes.NewReader(b))
	var thetaT *tensor.Dense
	err = dec.Decode(&thetaT)
	if err != nil {
		log.Fatal(err)
	}
	g := gorgonia.NewGraph()
	theta := gorgonia.NodeFromAny(g, thetaT, gorgonia.WithName("theta"))
	values := make([]float64, 5)
	xT := tensor.New(tensor.WithBacking(values))
	x := gorgonia.NodeFromAny(g, xT, gorgonia.WithName("x"))
	y, err := gorgonia.Mul(x, theta)
	if err != nil {
		log.Fatal(err)
	}
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()
	values[4] = 1.0
	values[0] = 5.1
	values[1] = 3.5
	values[2] = 1.4
	values[3] = 0.2

	if err = machine.RunAll(); err != nil {
		log.Fatal(err)
	}
	switch math.Round(y.Value().Data().(float64)) {
	case 1:
		fmt.Println("setosa")
	case 2:
		fmt.Println("virginica")
	case 3:
		fmt.Println("versicolor")
	default:
		fmt.Println("unknown iris")
	}
	machine.Reset()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
