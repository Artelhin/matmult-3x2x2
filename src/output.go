package src

import (
	"fmt"
	"log"
	"os"
)

type Outputer struct {
	out, log *os.File
}

var Out *Outputer

func (o *Outputer) Report(s string) {
	if o.out == nil {
		fmt.Print(s)
	}
	if _, err := o.out.WriteString(s); err != nil {
		fmt.Printf("error: can't write to output file: %s", err)
	}
}

func (o *Outputer) Log(s string) {
	if o.log == nil {
		return
	}
	if _, err := o.log.WriteString(s); err != nil {
		fmt.Printf("error: can't write to log file: %s", err)
	}
}

func (o *Outputer) Info(s string) {
	log.Println(s)
}

//logFile can be passed as nil
func SetOutput(outFile, logFile *os.File) {
	Out = new(Outputer)
	Out.out = outFile
	Out.log = logFile
}
