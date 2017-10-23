package main

import (
	"os"
	//"strconv"
)

type mariadb struct {
	Cluster       bool
	ClusterStatus string
}

func (m *mariadb) CreateCluster() {
	logger.Println("now into CreateCluster")
	var u util
	cmdStr := "ls"
	ret := u.RunCommand(cmdStr)
	logger.Println("show ls result:", ret)

}

func (m *mariadb) CheckEnvCluster() bool {
	envCluster := os.Getenv("CLUSTER")
	infoLog.Println("show initial env Cluster", envCluster)
	switch envCluster {
	case "true", "":
		return true
	case "false":
		return false
	default:
		return true
	}
	return true
}

func (m *mariadb) CheckMode() string {
	m.Cluster = m.CheckEnvCluster()
	infoLog.Println("shwo cluster mode:", m.Cluster)
	if !m.Cluster {
		return "single"
	} else {
		// it could be build, reboot, repair
		// check neibor node (false and file existed=reboot)
		// check neibor node (false and file not existed=build)
		// check neibor node (true and file existed=reboot)
		// check neibor node (true and file not existed=repair)
		infoLog.Println("now is in CLUSTER MODE")
		return "build"
	}

	return "no match"
}

func (m *mariadb) MariadbRun() {
	var w webserver
	go w.webserver()
	// check mode
	// select different mode
	mode := m.CheckMode()

	switch mode {
	case "single":
		infoLog.Println("now we are into single mode")
	case "build":
		//logger.Println("start to build mode ...")
		infoLog.Println("start to build mode ....")
	case "reboot":
		logger.Println("start to reboot mode ...")
	case "repair":
		logger.Println("start to repair mode ...")
	default:
		//fatal will stop the daemon
		errorLog.Fatal("not on defined mode")
	}
	m.CreateCluster()
	select {}
}

func (m *mariadb) Run() {
	m.MariadbRun()
}
