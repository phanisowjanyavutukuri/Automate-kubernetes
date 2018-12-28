package kube_proxy
import "os/exec"
func Kube_proxy_kubeconfig () {

command := "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=kube-proxy.kubeconfig"
_, _ = exec.Command("/bin/sh", "-c", command).Output()

command = "kubectl config set-credentials system:kube-proxy --client-certificate=kube-proxy.pem --client-key=kube-proxy-key.pem --embed-certs=true --kubeconfig=kube-proxy.kubeconfig"
_, _ = exec.Command("/bin/sh", "-c", command).Output()

command = "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=system:kube-proxy --kubeconfig=kube-proxy.kubeconfig"
_,_ = exec.Command("/bin/sh", "-c", command).Output()
command = "kubectl config use-context default --kubeconfig=kube-proxy.kubeconfig"
_, _ = exec.Command("/bin/sh", "-c", command).Output()

}
