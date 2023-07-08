package exception

import (
	"buckmate/structs"
	"log"
	"os"
)

func Handle(exception structs.Exception) {
	if exception.Err != nil {
		exitCode := 1
		if exception.ExitCode != 0 {
			exitCode = exception.ExitCode
		}
		log.Fatalf("Oops: %s, %v", exception.Message, exception.Err)
		os.Exit(exitCode)
	}
}
