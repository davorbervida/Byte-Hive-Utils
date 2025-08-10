package startup

import (
	"fmt"
	"os"
	"strings"
)

func GetPort() string {

	// Default given port
	var port string = ":8081"

	if len(os.Args) == 1 {
		return port
	}

	if idx := strings.IndexByte(os.Args[1], '='); idx != -1 {
		port = fmt.Sprintf(":%s", os.Args[1][idx+1:])
	}

	return port
}
