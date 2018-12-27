package main

import (
	"fmt"
	_ "log"
	"os/exec"
)

func main() {
	Command := "mkdir ~/bin"
	out, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err == nil {
		Command = "curl -s -L -o ~/bin/cfssl https://pkg.cfssl.org/R1.2/cfssl_linux-amd64"
		out, err = exec.Command("/bin/sh", "-c", Command).Output()
		if err == nil {
			Command = "curl -s -L -o ~/bin/cfssljson https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64"
			out, err = exec.Command("/bin/sh", "-c", Command).Output()
			if err == nil {
				Command = "chmod +x ~/bin/{cfssl,cfssljson}"
				out, err = exec.Command("/bin/sh", "-c", Command).Output()
				if err == nil {
					Command = "export PATH=$PATH:~/bin"
					out, err = exec.Command("/bin/sh", "-c", Command).Output()
				}
			}
		}
	}

	Command = "cfssl gencert -initca ca-csr.json | cfssljson -bare ca"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes  admin-csr.json | cfssljson -bare admin"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -hostname=worker1,35.232.234.58,10.128.0.4 -profile=kubernetes worker1-csr.json | cfssljson -bare worker1"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-controller-manager-csr.json | cfssljson -bare kube-controller-manager"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-proxy-csr.json | cfssljson -bare kube-proxy"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-scheduler-csr.json | cfssljson -bare kube-scheduler"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()

	fmt.Println(out)

}
