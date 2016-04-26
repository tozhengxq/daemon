package main

import (
    "fmt"
    "time"
    "flag"
    "os"
    "os/exec"
)

var daemon = flag.Bool("d", false, "run app as a daemon with -d=true")
func init() {
    if !flag.Parsed() {
        flag.Parse()
    }
    if *daemon {
        args := os.Args[1:]
        i := 0
        for ; i < len(args); i++ {
            if args[i] == "-d=true" {
                args[i] = "-d=false"
                break
            }
        }
        cmd := exec.Command(os.Args[0], args...)
        cmd.Start()
        fmt.Println("[PID]", cmd.Process.Pid)
        os.Exit(0)
    }
}

func main(){
    for {
        time.Sleep(3*time.Second)
        fmt.Println("as")
    }

}
