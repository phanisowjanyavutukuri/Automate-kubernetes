package kube_scheduler
import "os/exec"
func Kube_scheduler_kubeconfig () {

   
   _, _ = exec.Command("/bin/sh", "-c", "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=kube-scheduler.kubeconfig").Output()
  
    _, _ = exec.Command("/bin/sh", "-c", "kubectl config set-credentials system:kube-scheduler --client-certificate=kube-scheduler.pem --client-key=kube-scheduler-key.pem --embed-certs=true --kubeconfig=kube-scheduler.kubeconfig").Output()
    
    _, _ = exec.Command("/bin/sh", "-c", "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=system:kube-scheduler --kubeconfig=kube-scheduler.kubeconfig").Output()
    
    _,_ = exec.Command("/bin/sh", "-c", "kubectl config use-context default --kubeconfig=kube-scheduler.kubeconfig").Output()
}
