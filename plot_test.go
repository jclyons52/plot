package plot

// test plotting sin function
import (
	"io"
	"math"
	"testing"
)

func TestPlotSinCosGraphs(t *testing.T) {
	WriteToFile("sin.gif", func(out io.Writer) {
		plot := NewPlot(-math.Pi, math.Pi, -1, 1, 500)
		plot.Draw(out, math.Sin, math.Cos)
	})
}

func TestPlotSinCosGraphsMoving(t *testing.T) {
	WriteToFile("sinmove.gif", func(out io.Writer) {
		plot := NewPlot(-math.Pi, math.Pi, -1, 1, 500)
		plot.DrawMoving(out, math.Sin, math.Cos)
	})
}

func TestPlotExponentialGraph(t *testing.T) {
	WriteToFile("exp.gif", func(out io.Writer) {
		eq := func(x float64) float64 {
			return math.Exp(x)
		}
		plot := NewPlot(0, math.Pi, 0, math.Exp(math.Pi), 500)
		plot.Draw(out, eq)
	})
}

func TestPlotExponentialGraphMoving(t *testing.T) {
	WriteToFile("expmove.gif", func(out io.Writer) {
		eq := func(x float64) float64 {
			return math.Exp(x)
		}
		plot := NewPlot(0, math.Pi, 0, math.Exp(math.Pi), 500)
		plot.DrawMoving(out, eq)
	})
}
