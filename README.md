# Server Manager

## Usage

```bash
go run main/main.go

or

go build main/main.go
```

```GO
package main

import (
	"fmt"

	"github.com/DustiasTheGuy/servman/service"
)

func main() {
	s := service.Service{
		ProccessID: 0,
		Debug:      false,
		Path:       "main.exe",
		WorkingDir: "D:/Development/GO/isak_tech/server",
	}

	if err := s.StartService(); err != nil { // the pid will get updated if the proccess was successfully started
	    log.Fatal(err)
	}

	fmt.Println(s.IsAlive())
}

```
## License
[MIT](https://choosealicense.com/licenses/mit/)