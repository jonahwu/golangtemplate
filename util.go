package main

import (
	"os/exec"
)

type util struct {
	Cluster       bool
	ClusterStatus string
}

var test string

func init() {
	infoLog.Println("init util running---------")
	test = "initial test"
}

func (u *util) GetCluster() {
	u.ClusterStatus = test
}

func (u *util) RunCommand(commandStr string) string {
	cmdstr := commandStr
	out, _ := exec.Command("sh", "-c", cmdstr).Output()
	strout := string(out)
	return strout
}
