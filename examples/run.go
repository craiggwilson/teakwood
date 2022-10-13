package examples

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var logFlag = flag.String("logfile", "", "the file to log")

func Run(mdl tea.Model, opts ...tea.ProgramOption) {
	flag.Parse()

	if logFlag != nil && len(*logFlag) > 0 {
		f, err := os.OpenFile(*logFlag, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(fmt.Errorf("opening log file %q: %w", *logFlag, err))
		}
		defer f.Close()
		log.SetOutput(f)
	} else {
		log.SetOutput(io.Discard)
	}

	p := tea.NewProgram(mdl, opts...)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
