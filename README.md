# daemon
Daemonize Go applications with 

# Usage
go get github.com/tozhengxq/godamon

``` python
// test.go
package main
import (
    "fmt"
    "time"
    _ "github.com/tozhengxq/godaemon"
)
func main(){
    for {
        time.Sleep(3*time.Second)
        fmt.Println("as")
    }
}
```
``` python
$ go run test.go -d=true
[PID] 167
```
