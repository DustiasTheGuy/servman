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
	"log"

	"github.com/DustiasTheGuy/servman/service"
)

func main() {
	s := service.Service{
		ProccessID: 0,
		Debug:      true,
		Path:       "main.exe",
		WorkingDir: "",
	}

	if err := s.StartService(); err != nil {
		// the pid will be updated if proccess
		// launches successfully
		log.Fatal(err)
	}

	fmt.Println(s.IsAlive())
}
```
## License
[MIT](https://choosealicense.com/licenses/mit/)