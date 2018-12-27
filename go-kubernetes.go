package main
import (
    "fmt"
    "log"
    "os/exec"
)
func main(){
	Command := "mkdir ~/bin"
		out, err := exec.Command(Command).Output()
		if err == nil {
	Command = "curl -s -L -o ~/bin/cfssl https://pkg.cfssl.org/R1.2/cfssl_linux-amd64"
		out, err = exec.Command(Command).Output()
		if err == nil {
	Command = "curl -s -L -o ~/bin/cfssljson https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64"
		out, err = exec.Command(Command).Output()
		if err == nil {
	Command = "chmod +x ~/bin/{cfssl,cfssljson}"
		out, err = exec.Command(Command).Output()
		if err == nil {
	Command = export PATH=$PATH:~/bin
		out, err = exec.Command(Command).Output()
		}
		}
		}
		}
		

		Command = "cfssl gencert -initca ca-csr.json | cfssljson -bare ca"
		out,err = exec.Command(Command).Output()




}