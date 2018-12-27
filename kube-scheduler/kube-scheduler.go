package kube-controller
import "os/exec"
func Kube-scheduler-kubeconfig () {

    command := "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=kube-scheduler.kubeconfig"
    out, err := exec.Command("/bin/sh", "-c", command).Output()
    command = "kubectl config set-credentials system:kube-scheduler --client-certificate=kube-scheduler.pem --client-key=kube-scheduler-key.pem --embed-certs=true --kubeconfig=kube-scheduler.kubeconfig"
    out, err = exec.Command("/bin/sh", "-c", command).Output()
    command = "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=system:kube-scheduler --kubeconfig=kube-scheduler.kubeconfig"
    out, err = exec.Command("/bin/sh", "-c", command).Output()
    command = "kubectl config use-context default --kubeconfig=kube-scheduler.kubeconfig"
    out, err = exec.Command("/bin/sh", "-c", command).Output()
}