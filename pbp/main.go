package pbp

import (
	"fmt"
	"os/exec"
)

func GenPbpFromIso(isoPath string, outputPBP string, contentID, string, versionKey string, compress bool) string {
	compressVal := ""
	if compress {
		compressVal = "-c"
	}

	fmt.Println("@sign_np@", "-pbp", compressVal, isoPath, outputPBP, contentID, versionKey)
	signnp := exec.Command("@sign_np@", "-pbp", compressVal, isoPath, outputPBP, contentID, versionKey)

	signnp.Start()
	fmt.Println(signnp.CombinedOutput())
	return ""
}
