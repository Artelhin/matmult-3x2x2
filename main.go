package main

import (
	"diplom/src"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		kFlag   = flag.Int("k", 2, "size of the field (simple number)")
		outFlag = flag.String("o", "", "name of the output file")
		outLog  = flag.String("name ol", "", "file to output each different solution")
	)

	flag.Parse()

	var (
		out, log *os.File
		err      error
	)

	if *outFlag != "" {
		out, err = os.Create(*outFlag)
		if err != nil {
			fmt.Printf("can't create output file: %s\nusing std output for report", err)
		}
	}
	if *outLog != "" {
		out, err = os.Create(*outLog)
		if err != nil {
			fmt.Printf("can't create log file: %s\nwon't return any logs", err)
		}
	}

	src.SetOutput(out, log)

	src.CheckFieldK(*kFlag)
	//fmt.Println(src.CheckParams(-1, -1, 1, 0, 0, 1, 0))
}
