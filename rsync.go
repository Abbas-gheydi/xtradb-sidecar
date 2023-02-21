package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	rsyncLock sync.Mutex
)
var (
	sshPort             string
	sshidentityFile     string
	sshUsername         string
	sshServerAddr       string
	RemoteDir           string
	enablerSyncToRemote bool
)
var rsyncCmd []string

/*= []string{
	"-avh",
	"--delete",
	"-e",
	"ssh -p " + sshPort + " -i " + sshidentityFile,
	targetRootDir,
	sshUsername + "@" + sshServerAddr + ":" + RemoteDir,
}
*/

func rsync() {
	rsyncLock.Lock()

	log.Println("start: ", "rsync")
	cmd := exec.Command("rsync", rsyncCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	defer rsyncLock.Unlock()
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}

}
