package main

import (
	"encoding/json"
	"strings"
	"time"

	"code.wirelessregistry.com/data/readers/acl"

	"fmt"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	persistentPath              = "/testPersistent"
	ephemeralPath               = "/testEphemeral"
	aclPath                     = "/acl"
	role1          acl.RoleName = "role1"
	role2          acl.RoleName = "role2"
	admin          acl.RoleName = "admin"
	cmd1                        = "cmd1"
	cmd2                        = "cmd2"
	cmd3                        = "cmd3"
	const2                      = "constraint1"
)

func main() {
	conn := connect()
	defer conn.Close()
	//listChildren(conn)
	//createPersistent(persistentPath, conn)
	//listChildren(conn)
	//createEphemeral(ephemeralPath, conn)
	//listChildren(conn)
	//printData(ephemeralPath, conn)
	//delete(ephemeralPath, conn)
	//listChildren(conn)
	go createACLWatchClient(aclPath)
	createPersistent(aclPath, conn)
	updateData(aclPath, getTestDataACL(), conn)
	printData(aclPath, conn)
	testACL(aclPath, conn)
}

func connect() *zk.Conn {
	zksStr := "127.0.0.1:2181,127.0.0.1:2182,127.0.0.1:2183"
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	must(err)
	return conn
}

func listChildren(conn *zk.Conn) {
	childs, stat, _ := conn.Children("/")
	fmt.Printf("Children: %v, %v\n", childs, stat.Version)
}

func createPersistent(path string, conn *zk.Conn) {
	fmt.Printf("Creating path %s\n", path)
	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	conn.Create(path, getTestData(), flags, acl)
}

func createEphemeral(path string, conn *zk.Conn) {
	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)
	conn.Create(path, getTestData(), flags, acl)
}

func updateData(path string, data []byte, conn *zk.Conn) {
	fmt.Printf("Updating path %s\n", path)
	_, version := getData(path, conn)
	conn.Set(path, data, version)
}

func getData(path string, conn *zk.Conn) (string, int32) {
	data, stat, _ := conn.Get(path)
	return convertTestData(data), stat.Version
}
func printData(path string, conn *zk.Conn) {
	data, version := getData(path, conn)
	fmt.Printf("%s, %d\n", data, version)
}

func delete(path string, conn *zk.Conn) {
	_, version := getData(path, conn)
	conn.Delete(path, version)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func getTestData() []byte {
	data := fmt.Sprintf("%d", time.Now().Unix())
	return []byte(data)
}

func convertTestData(data []byte) string {
	return string(data[:])
}

func getTestDataACL() []byte {
	anAcl := acl.NewACL()
	r := acl.NewRole(role1)
	p := &acl.Permission{cmd1}
	c := acl.NewAppListConstraint([]string{const2})
	anAcl.AddPermission(p)

	r.AddPermission(p, c)
	anAcl.AddRole(r)

	jsonDoc, _ := json.Marshal(anAcl)
	return jsonDoc
}

func readACL(path string, conn *zk.Conn) *acl.ACL {
	data, _, _ := conn.Get(path)
	newACL := acl.NewACL()
	json.Unmarshal(data, newACL)
	return newACL
}

func testACL(path string, conn *zk.Conn) {
	readACL := readACL(path, conn)
	allowed, cons := readACL.Can(role1, cmd1)
	fmt.Printf("%v: %v\n", allowed, cons)
}

func createACLWatchClient(path string) {
	wConn := connect()
	fmt.Printf("Starting watch for ACL change...\n")
	defer wConn.Close()
	data, _, ech, _ := wConn.GetW(path)
	evt := <-ech
	fmt.Printf("Catching event ACL change: %v: %s\n", evt, string(data[:]))
}
