package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	s := `
	[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
	Invoke-WebRequest -Uri "https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-all-3.6.1.tar.gz" -OutFile ".\protobuf-all-3.6.1.tar.gz"
	tar -zxvf ".\protobuf-all-3.6.1.tar.gz"
	Move-Item -Path protobuf-3.6.1 -Destination $ALLUSERSPROFILE/protobuf-3.6.1
	cd $ALLUSERSPROFILE/protobuf-3.6.1
	`

	cmd := exec.Command("powershell", "-c", s)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
