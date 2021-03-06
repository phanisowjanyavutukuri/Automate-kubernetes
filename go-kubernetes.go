package main

import (
	"fmt"
	_ "log"
	"os"
	"os/exec"
	"Automate-kubernetes/kube-proxy"
	"Automate-kubernetes/kube-controller"
	"Automate-kubernetes/kube-scheduler"
	"Automate-kubernetes/admin-kubeconfig"
)

func main() {

	out, err := exec.Command("/bin/bash", "installcfssl.sh").Output()
	if err !=nil{
		fmt.Println(err)
		fmt.Println(out)
	}
	fmt.Println("*******")
//	fmt.Println(err)

	if _, err1 := os.Stat("ca-key.pem"); os.IsNotExist(err1) {
		out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -initca ca-csr.json | cfssljson -bare ca").Output()
	}

	if _, err2 := os.Stat("admin-key.pem"); os.IsNotExist(err2) {
		out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes  admin-csr.json | cfssljson -bare admin").Output()
	}

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -hostname=worker1,35.232.234.58,10.128.0.4 -profile=kubernetes worker1-csr.json | cfssljson -bare worker1").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-controller-manager-csr.json | cfssljson -bare kube-controller-manager").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-proxy-csr.json | cfssljson -bare kube-proxy").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kube-scheduler-csr.json | cfssljson -bare kube-scheduler").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -hostname=10.128.0.4,35.232.234.58,127.0.0.1,kubernetes.default -profile=kubernetes kubernetes-csr.json | cfssljson -bare kubernetes").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes service-account-csr.json | cfssljson -bare service-account").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=worker1.kubeconfig").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "kubectl config set-credentials system:node:worker1 --client-certificate=worker1.pem --client-key=worker1-key.pem --embed-certs=true --kubeconfig=worker1.kubeconfig").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=system:node:worker1 --kubeconfig=worker1.kubeconfig").Output()

	if err !=nil{
		fmt.Println(err)
	}
	out, err = exec.Command("/bin/sh", "-c", "kubectl config use-context default --kubeconfig=worker1.kubeconfig").Output()

	kube_proxy.Kube_proxy_kubeconfig()
	kube_controller.Kube_controller_kubeconfig()
	kube_scheduler.Kube_scheduler_kubeconfig()
	kube_admin.Kube_admin_kubeconfig()

	out, err = exec.Command("/bin/sh", "-c", "chmod 777  encryptionconfig.sh").Output()
	out, err = exec.Command("/bin/sh", "-c", "source encryptionconfig.sh; encrypt").Output()
	out, err = exec.Command("/bin/sh", "-c", "install_etcd.sh").Output()

	out, err = exec.Command("/bin/sh", "-c", "chmod 777  encryptionconfig.sh").Output()
	out, err = exec.Command("/bin/sh", "-c", "source etcd_service.sh; etcd_service").Output()
	out, err = exec.Command("/bin/sh", "-c", "install-api-controller-scheduler.sh").Output()
	out, err = exec.Command("/bin/sh", "-c", "apiserver.sh").Output()
	out, err = exec.Command("/bin/sh", "-c", "controller.sh").Output()
	out, err = exec.Command("/bin/sh", "-c", "kube-scheduler.yaml").Output()
	out, err = exec.Command("/bin/sh", "-c", "kube-scheduler-service").Output()
	out, err = exec.Command("/bin/sh", "-c", "enableservice").Output()
	out, err = exec.Command("/bin/sh", "-c", "workerbinaries.sh").Output()

	out, err = exec.Command("/bin/sh", "-c", "source bridge.sh; bridge").Output()
	out, err = exec.Command("/bin/sh", "-c", "kubelet.sh").Output()

	fmt.Println(out)

}
