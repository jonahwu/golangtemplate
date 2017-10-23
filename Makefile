# This how we want to name the binary output
# #BINARY=glustercluster
# # # Builds the project
build:
	go build db.go mariadb.go util.go common.go webserver.go
install:
	cp mariadbcluster /root/../mariadbcluster
clean:
	rm -rf /var/lib/mysql/grastate.dat

