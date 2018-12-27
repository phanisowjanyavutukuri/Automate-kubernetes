package kube-controller
import "os/exec"
func Kube-controller-kubeconfig () {
command := "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=kube-controller-manager.kubeconfig"
out, err := exec.Command("/bin/sh", "-c", command).Output()
command = "kubectl config set-credentials system:kube-controller-manager --client-certificate=kube-controller-manager.pem --client-key=kube-controller-manager-key.pem --embed-certs=true --kubeconfig=kube-controller-manager.kubeconfig"
out, err = exec.Command("/bin/sh", "-c", command).Output()
command = "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=system:kube-controller-manager --kubeconfig=kube-controller-manager.kubeconfig"
out, err = exec.Command("/bin/sh", "-c", command).Output()
command = "kubectl config use-context default --kubeconfig=kube-controller-manager.kubeconfig"
out, err = exec.Command("/bin/sh", "-c", command).Output()
}