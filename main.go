package main

import (
	"flag"
	"log"
	"time"

	psutils "github.com/shirou/gopsutil/process"
)

func main() {
	processName := flag.String("processName", "GTA5.exe", "name of the process to suspend")
	suspendDuration := flag.Uint("duration", 10, "how long the process is goign to be suspended for")
	flag.Parse()

	processes, err := psutils.Processes()
	if err != nil {
		log.Fatalf("Error while getting process list : %s", err)
	}

	var process *psutils.Process
	var found bool
	for _, process = range processes {
		name, _ := process.Name()
		if name == *processName {
			found = true
			break
		}
	}

	if !found {
		log.Printf("Unable to find %s !", *processName)
		return
	}

	log.Printf("Process found ! It's PID is %6d", process.Pid)
	err = process.Suspend()
	if err != nil {
		log.Fatalf("Unable to suspend process : %s", err)
	}
	log.Printf("Process suspended.")

	// Wait a few seconds
	for i := *suspendDuration; i > 0; i-- {
		log.Printf("Resuming process in %2d seconds.", i)
		time.Sleep(time.Second)
	}

	err = process.Resume()
	if err != nil {
		log.Fatalf("Unable to resume process : %s", err)
	}

	log.Print("All good, you should be solo in your session :)")
}
