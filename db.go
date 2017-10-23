package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var infoLog *log.Logger
var errorLog *log.Logger
var flagCleanConfig bool
var flagStdOut bool
var dataDir string

type writer struct {
	io.Writer
	timeFormat string
}

func (w writer) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.timeFormat)), b...))
}

func init() {
	file, err := os.Create("/var/log/template.log")
	if err != nil {
		log.Fatal(err)
	}
	flag.BoolVar(&flagCleanConfig, "c", true, "clean gluster config")
	flag.StringVar(&dataDir, "d", "/data/glusterfs", "clean gluster config")
	flag.BoolVar(&flagStdOut, "o", false, "stdout to screen for debug")
	flag.Parse()

	logger = log.New(file, "crm_", log.LstdFlags|log.Llongfile)
	infoLog = log.New(file, "[INFO]", log.Llongfile)
	infoLog.SetFlags(infoLog.Flags() | log.LstdFlags)
	//comment above if use self defined log as followed
	//hname, _ := os.Hostname()
	//infoLog.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " INFO " + hname) "
	errorLog = log.New(file, "[ERROR]", log.Llongfile)
	errorLog.SetFlags(errorLog.Flags() | log.LstdFlags)
	errorLog.SetOutput(os.Stdout)
	if flagStdOut {
		infoLog.SetOutput(os.Stdout)
		errorLog.SetOutput(os.Stdout)
		logger.SetOutput(os.Stdout)
	}

}

func main() {
	fmt.Println("vim-go")
	var m mariadb
	m.Run()
}
