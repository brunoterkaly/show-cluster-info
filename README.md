# Show cluster info

This is a GoLang application that completely reads cluster information and outputs in two ways:

(1) Loop through *nodes* and show all pods in node (including containers, images, and labels per pod)
(2) Loop through *namespaces* and show all pods in node (including containers, images, and labels per pod)

### Assumptions

Before compiling and running the code:
- Make sure KUBECONFIG is properly set to point to the credentials for your Kubernetes Cluster

### Output

```
  Node: k8s-agent-dce67bdc-0
  Node: k8s-agent-dce67bdc-1
  Node: k8s-agent-dce67bdc-2
  Node: k8s-master-dce67bdc-0
  Namespace: default
  Namespace: kube-public
  Namespace: kube-system
  Namespace: staging
  ----------
  Node Name = k8s-agent-dce67bdc-1
    ----------
    Pod Name = busybox (default)
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3 (default)
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
    ----------
    Pod Name = heapster-342135353-vrqjj (kube-system)
      ----------
       Container Name = heapster
       Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
      ----------
       Container Name = heapster-nanny
       Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
    ----------
    Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-addon-manager
       Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
    ----------
    Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-apiserver
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-controller-manager
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-dns-v20-3003781527-sv2k1 (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-proxy-bfd5p (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-g536m (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-mdt7c (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-nj57n (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-scheduler-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-scheduler
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kubernetes-dashboard-924040265-vjppn (kube-system)
      ----------
       Container Name = kubernetes-dashboard
       Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
    ----------
    Pod Name = tiller-deploy-677436516-2wvg1 (kube-system)
      ----------
       Container Name = tiller
       Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
    ----------
    Pod Name = flask-monolith (staging)
      ----------
       Container Name = mymonolith-container
       Image Name = brunoterkaly/monolith
  ----------
  Node Name = k8s-agent-dce67bdc-2
    ----------
    Pod Name = busybox (default)
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3 (default)
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
    ----------
    Pod Name = heapster-342135353-vrqjj (kube-system)
      ----------
       Container Name = heapster
       Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
      ----------
       Container Name = heapster-nanny
       Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
    ----------
    Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-addon-manager
       Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
    ----------
    Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-apiserver
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-controller-manager
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-dns-v20-3003781527-sv2k1 (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-proxy-bfd5p (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-g536m (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-mdt7c (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-nj57n (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-scheduler-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-scheduler
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kubernetes-dashboard-924040265-vjppn (kube-system)
      ----------
       Container Name = kubernetes-dashboard
       Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
    ----------
    Pod Name = tiller-deploy-677436516-2wvg1 (kube-system)
      ----------
       Container Name = tiller
       Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
  ----------
  Node Name = k8s-master-dce67bdc-0
    ----------
    Pod Name = busybox (default)
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3 (default)
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
    ----------
    Pod Name = heapster-342135353-vrqjj (kube-system)
      ----------
       Container Name = heapster
       Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
      ----------
       Container Name = heapster-nanny
       Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
    ----------
    Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-addon-manager
       Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
    ----------
    Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-apiserver
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-controller-manager
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-dns-v20-3003781527-sv2k1 (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-proxy-bfd5p (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-g536m (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-mdt7c (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-nj57n (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-scheduler-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-scheduler
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Node Name = k8s-agent-dce67bdc-0
    ----------
    Pod Name = busybox (default)
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3 (default)
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
    ----------
    Pod Name = heapster-342135353-vrqjj (kube-system)
      ----------
       Container Name = heapster
       Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
      ----------
       Container Name = heapster-nanny
       Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
    ----------
    Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-addon-manager
       Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
    ----------
    Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-apiserver
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
      ----------
       Container Name = kube-controller-manager
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-dns-v20-3003781527-sv2k1 (kube-system)
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-proxy-bfd5p (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-g536m (kube-system)
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Namespace = kube-system
    ----------
    Pod Name = busybox
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
    ----------
    Pod Name = heapster-342135353-vrqjj
      ----------
       Container Name = heapster
       Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
      ----------
       Container Name = heapster-nanny
       Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
    ----------
    Pod Name = kube-addon-manager-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-addon-manager
       Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
    ----------
    Pod Name = kube-apiserver-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-apiserver
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-controller-manager-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-controller-manager
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-dns-v20-3003781527-2djkm
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-dns-v20-3003781527-sv2k1
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-proxy-bfd5p
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-g536m
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-mdt7c
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-nj57n
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-scheduler-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-scheduler
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kubernetes-dashboard-924040265-vjppn
      ----------
       Container Name = kubernetes-dashboard
       Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
    ----------
    Pod Name = tiller-deploy-677436516-2wvg1
      ----------
       Container Name = tiller
       Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
  ----------
  Namespace = staging
    ----------
    Pod Name = busybox
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
    ----------
    Pod Name = heapster-342135353-vrqjj
      ----------
       Container Name = heapster
       Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
      ----------
       Container Name = heapster-nanny
       Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
    ----------
    Pod Name = kube-addon-manager-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-addon-manager
       Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
    ----------
    Pod Name = kube-apiserver-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-apiserver
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-controller-manager-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-controller-manager
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-dns-v20-3003781527-2djkm
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-dns-v20-3003781527-sv2k1
      ----------
       Container Name = kubedns
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-kube-dns-amd64:1.14.5
      ----------
       Container Name = dnsmasq
       Image Name = gcrio.azureedge.net/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
      ----------
       Container Name = healthz
       Image Name = gcrio.azureedge.net/google_containers/exechealthz-amd64:1.2
    ----------
    Pod Name = kube-proxy-bfd5p
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-g536m
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-mdt7c
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-proxy-nj57n
      ----------
       Container Name = kube-proxy
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kube-scheduler-k8s-master-dce67bdc-0
      ----------
       Container Name = kube-scheduler
       Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
    ----------
    Pod Name = kubernetes-dashboard-924040265-vjppn
      ----------
       Container Name = kubernetes-dashboard
       Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
    ----------
    Pod Name = tiller-deploy-677436516-2wvg1
      ----------
       Container Name = tiller
       Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
    ----------
    Pod Name = flask-monolith
      ----------
       Container Name = mymonolith-container
       Image Name = brunoterkaly/monolith
  ----------
  Namespace = default
    ----------
    Pod Name = busybox
      ----------
       Container Name = busybox
       Image Name = busybox
    ----------
    Pod Name = curl-2716574283-m72z3
      ----------
       Container Name = curl
       Image Name = radial/busyboxplus:curl
  ----------
  Namespace = kube-public
  done!
```
