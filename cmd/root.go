package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/happyRip/Box-Tailor/box"
	"github.com/happyRip/Box-Tailor/box/lidded"
	"github.com/happyRip/Box-Tailor/box/utility"
	"github.com/happyRip/Box-Tailor/plotter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	output         string
	dimensions     []float64
	buffer         []float64
	kerf           float64
	boardThickness float64
	debug          bool
)

var rootCmd = &cobra.Command{
	Use:   "Box-Tailor",
	Short: "Simple CLI app that is used to create parametric boxes.",
	Run: func(cmd *cobra.Command, _ []string) {
		p := box.Product{}
		if len(dimensions) == 0 {
			p.ProcessUserInput()
		} else {
			if len(dimensions) != 3 {
				log.Fatal("Incorrect number of dimensions provided (must be 3, eg. `--dimensions 6,4.5,3`).")
			}
			if len(buffer) != 0 && (len(buffer) < 2 || len(buffer) > 3) {
				log.Fatal("Incorrect number of buffer dimensions provided (must be either 0, 2 or 3, eg. `--buffer 2.3,4`)")
			}

			for i := range dimensions {
				dimensions[i] *= 10
			}
			for i := range buffer {
				if i >= 3 {
					break
				}
				dimensions[i] += buffer[i] * 10
			}

			p.Name = output
			d := dimensions
			x, y, z := d[0], d[1], d[2]
			if x > y {
				x, y = y, x
			}
			p.SetSize(x, y, z)
		}

		bottom := lidded.Box{
			Content:        p,
			BoardThickness: boardThickness,
			Kerf:           kerf,
			Origin: utility.Pair{
				X: 0,
				Y: -12,
			},
			Debug: debug,
		}
		origin, _ := bottom.CalculateSize()
		lid := lidded.Box{
			Content: box.Product{
				Size: utility.Triad{
					X: p.Size.X + 2*(boardThickness+1),
					Y: p.Size.Y + 2*(boardThickness+1) + boardThickness,
					Z: p.Size.Z,
				},
			},
			Origin: utility.Pair{
				X: origin + 5,
				Y: -12,
			},
			BoardThickness: boardThickness,
			Kerf:           kerf,
			Debug:          debug,
		}

		outAbs, err := drawToSingleFile(
			p.Name,
			bottom,
			lid,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(outAbs)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Box-Tailor.yaml)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Name/path of the output file")
	rootCmd.PersistentFlags().Float64SliceVarP(&dimensions, "dimensions", "d", []float64{}, "Dimensions of object inside the box")
	rootCmd.PersistentFlags().Float64SliceVarP(&buffer, "buffer", "b", []float64{}, "Additional space around the product")
	rootCmd.PersistentFlags().Float64VarP(&kerf, "kerf", "k", 2., "Offset to account for width of material removed by the manufacturing process")
	rootCmd.PersistentFlags().Float64VarP(&boardThickness, "board-thickness", "t", 6., "Thickness of the material boxes are cut from")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Turn on debug mode")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".Box-Tailor")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func drawToSingleFile(name string, bottom box.Drafter, lid box.Drafter) (string, error) {
	outputFile, err := plotter.NewPltFile(name, "", "")
	if err != nil {
		return "", err
	}
	outputFile.Initialize()

	outputFile.WriteString(
		plotter.SelectPen(5),
		plotter.DefineTerminator('$'),
		plotter.CharacterSize(6*0.75, 6*1.5),
		plotter.ConstructCommand("PU", 0, -150*40),
		plotter.Label(name),
	)

	for _, s := range bottom.Draw() {
		outputFile.WriteString(s)
	}

	x, _ := bottom.CalculateSize()
	pen := plotter.Pen{}
	outputFile.WriteString(pen.MoveAbsolute(x+1, 0))
	for _, s := range lid.Draw() {
		outputFile.WriteString(s)
	}

	outAbs, err := filepath.Abs(outputFile.Pointer.Name())
	if err != nil {
		return "", err
	}

	err = outputFile.Close()
	return outAbs, err
}
