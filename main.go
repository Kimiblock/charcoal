package main

import (
	//"fmt"
	"log"
	//"os"
	"golang.org/x/sys/unix"
	//"strings"
	"net"

	"github.com/google/nftables"
	"github.com/Kimiblock/pecho"
)

const (
	version		float32	=	0.1
)

var (
	connNft		*nftables.Conn
	err		error
	logChan		= pecho.MkChannel()
)

/* Special strings may be interpreted
	private		10.0.0.0 - 10.255.255.255, 172.16.0.0 - 172.31.255.255, 192.168.0.0 - 192.168.255.255 and fd00::/8
	Custom IPs not supported yet.
*/
type appOutPerms struct {
	allowIP		[]string
	denyIP		[]string
}

func echo(lvl string, msg string) {
	logChan <- []string{lvl, msg}
}

/* Returns whether the operation is success or not */
func setAppPerms(appCgroup string, outperm appOutPerms, appID string, sandboxEng string) bool {
	logChan <- []string{
		"debug",
		"Got firewall rules for " + appID + " from " + sandboxEng,
	}
	var table = nftables.Table {
		Name:	sandboxEng + "-" + appID,
		Family:	unix.NFPROTO_INET,
	}
	tableExt, errList := connNft.ListTableOfFamily(
		sandboxEng + "-" + appID,
		unix.NFPROTO_INET,
	)
	if errList != nil {
		log.Println("Error listing table: " + errList.Error() + ", treating as non-existent")
	} else if tableExt == nil {
		log.Println("Got nil from ListTable")
	} else {
		connNft.DelTable(&table)
		log.Println("Deleted previous table")
	}



	return true
}

func main() {
	go pecho.StartDaemon(logChan)
	log.Println("Starting charcoal", version, ", establishing connection to nftables")
	connNft, err = nftables.New()
	if err != nil {
		log.Fatalln("Could not establish connection to nftables: " + err.Error())
	}
	log.Println("Established connection")
}