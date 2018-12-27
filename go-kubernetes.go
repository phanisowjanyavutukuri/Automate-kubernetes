package main

import (
	"fmt"
	_ "log"
	"os/exec"
	"Automate-kubernetes/kube-proxy"
<<<<<<< HEAD
	"Automate-kubernetes/kube-scheduler"
=======
	"Automate-kubernetes/kube-scheuler"
>>>>>>> 5b95f28166079330aec64f649b79885c8cb42605
	"Automate-kubernetes/kube-controller"
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
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -hostname=10.128.0.4,35.232.234.58,127.0.0.1,kubernetes.default -profile=kubernetes kubernetes-csr.json | cfssljson -bare kubernetes"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes service-account-csr.json | cfssljson -bare service-account"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=worker1.kubeconfig"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "kubectl config set-credentials system:node:worker1 --client-certificate=worker1.pem --client-key=worker1-key.pem --embed-certs=true --kubeconfig=worker1.kubeconfig"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=system:node:worker1 --kubeconfig=worker1.kubeconfig"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	Command = "kubectl config use-context default --kubeconfig=worker1.kubeconfig"
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	
	kube_proxy.Kube_proxy_kubeconfig()
	kube_controller.Kube_controller_kubeconfig()
	kube_scheduler.Kube_scheduler_kubeconfig()
	kube_admin.Kube_admin_kubeconfig()

	Command = "source encryptionconfig.sh; encrypt"
	cmd := " chmod 777  encryptionconfig.sh"
	out, err = exec.Command("/bin/sh", "-c",cmd).Output()
	out, err = exec.Command("/bin/sh", "-c", Command).Output()
	out, err = exec.Command("/bin/sh", "-c",install_etcd.sh).Output()
	Command = "source etcd_service.sh; etcd_service"
	cmd := " chmod 777  encryptionconfig.sh"
	out, err = exec.Command("/bin/sh", "-c",cmd).Output()
	out, err = exec.Command("/bin/sh", "-c",install-api-controller-scheduler.sh).Output()
	



	fmt.Println(out)

}
