# daemon
Daemonize Go applications with 

# Usage
go get github.com/tozhengxq/damon

``` python
// test.go
package main
import (
    "fmt"
    "time"
    _ "github.com/tozhengxq/daemon"
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
