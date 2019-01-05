package main

import (
	"fmt"
	"github.com/zhenghaoz/gorse/base"
	"github.com/zhenghaoz/gorse/core"
	"github.com/zhenghaoz/gorse/model"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpuperf")
	if err != nil {
		log.Fatal(err)
	}
	if err = pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	data := core.LoadDataFromBuiltIn("ml-100k")
	svd := model.NewSVDpp(base.Params{
		base.NEpochs:    100,
		base.Reg:        0.1,
		base.Lr:         0.01,
		base.NFactors:   50,
		base.InitMean:   0,
		base.InitStdDev: 0.001,
	})
	out := core.CrossValidate(svd, data, []core.Evaluator{core.RMSE, core.MAE}, core.NewKFoldSplitter(5))
	fmt.Println(out)

	// Save memory profile
	f, err = os.Create("memprof")
	if err != nil {
		log.Fatal(err)
	}
	if err = pprof.WriteHeapProfile(f); err != nil {
		log.Fatal(err)
	}
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}
}
