package logging

import (
	"github.com/briandowns/spinner"
	"time"
)

func ProgressBar(spinner_text string, spinner_duration int, spinner_delay int) {
	s := spinner.New(spinner.CharSets[9], time.Duration(spinner_delay)*time.Millisecond)
	//go func() {
	s.Prefix = "["
	s.Suffix = "] " + spinner_text
	s.Start()
	time.Sleep(time.Duration(spinner_duration) * time.Millisecond)
	s.Stop()
	//}()
}
