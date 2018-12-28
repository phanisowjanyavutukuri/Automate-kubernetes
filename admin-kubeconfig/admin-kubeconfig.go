package kube_admin
import "os/exec"
func Kube_admin_kubeconfig () {
   
  _,_ = exec.Command("/bin/sh", "-c", "kubectl config set-cluster kubernetes-the-hard-way --certificate-authority=ca.pem --embed-certs=true --server=https://10.128.0.4:6443 --kubeconfig=admin.kubeconfig").Output()
  
  _, _ = exec.Command("/bin/sh", "-c", "kubectl config set-credentials admin --client-certificate=admin.pem --client-key=admin-key.pem --embed-certs=true --kubeconfig=admin.kubeconfig").Output()
  
  _,_ = exec.Command("/bin/sh", "-c", "kubectl config set-context default --cluster=kubernetes-the-hard-way --user=admin --kubeconfig=admin.kubeconfig").Output()
  
  _,_ = exec.Command("/bin/sh", "-c", "kubectl config use-context default --kubeconfig=admin.kubeconfig").Output()
  }
