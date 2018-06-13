# Show cluster info

This is a GoLang application that completely reads cluster information and outputs in two ways:

(1) Loop through **nodes** and show all pods in node (including containers, images, and labels per pod)

(2) Loop through **namespaces** and show all pods in node (including containers, images, and labels per pod)

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
Node Name = k8s-agent-dce67bdc-0
  ----------
  Pod Name = busybox (default)
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3 (default)
     Labels:
        Key   = pod-template-hash
        Value = 2716574283
        ----------
        Key   = run
        Value = curl
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
  ----------
  Pod Name = heapster-342135353-vrqjj (kube-system)
     Labels:
        Key   = pod-template-hash
        Value = 342135353
        ----------
        Key   = k8s-app
        Value = heapster
        ----------
     ----------
     Container Name = heapster
     Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
     ----------
     Container Name = heapster-nanny
     Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
  ----------
  Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
     ----------
     Container Name = kube-addon-manager
     Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
  ----------
  Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-apiserver
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-apiserver
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-controller-manager
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-controller-manager
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
     Labels:
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
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
     Labels:
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
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
     Labels:
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-g536m (kube-system)
     Labels:
        Key   = tier
        Value = node
        ----------
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
----------
Node Name = k8s-agent-dce67bdc-1
  ----------
  Pod Name = busybox (default)
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3 (default)
     Labels:
        Key   = pod-template-hash
        Value = 2716574283
        ----------
        Key   = run
        Value = curl
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
  ----------
  Pod Name = heapster-342135353-vrqjj (kube-system)
     Labels:
        Key   = k8s-app
        Value = heapster
        ----------
        Key   = pod-template-hash
        Value = 342135353
        ----------
     ----------
     Container Name = heapster
     Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
     ----------
     Container Name = heapster-nanny
     Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
  ----------
  Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
     ----------
     Container Name = kube-addon-manager
     Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
  ----------
  Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = tier
        Value = control-plane
        ----------
        Key   = component
        Value = kube-apiserver
        ----------
     ----------
     Container Name = kube-apiserver
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-controller-manager
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-controller-manager
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
     Labels:
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
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
     Labels:
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
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
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-g536m (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-mdt7c (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-nj57n (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-scheduler-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = tier
        Value = control-plane
        ----------
        Key   = component
        Value = kube-scheduler
        ----------
     ----------
     Container Name = kube-scheduler
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kubernetes-dashboard-924040265-vjppn (kube-system)
     Labels:
        Key   = k8s-app
        Value = kubernetes-dashboard
        ----------
        Key   = pod-template-hash
        Value = 924040265
        ----------
     ----------
     Container Name = kubernetes-dashboard
     Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
  ----------
  Pod Name = tiller-deploy-677436516-2wvg1 (kube-system)
     Labels:
        Key   = app
        Value = helm
        ----------
        Key   = name
        Value = tiller
        ----------
        Key   = pod-template-hash
        Value = 677436516
        ----------
     ----------
     Container Name = tiller
     Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
  ----------
  Pod Name = flask-monolith (staging)
     Labels:
        Key   = name
        Value = mymonolith
        ----------
        Key   = app
        Value = mymonolith-app
        ----------
     ----------
     Container Name = mymonolith-container
     Image Name = brunoterkaly/monolith
----------
Node Name = k8s-agent-dce67bdc-2
  ----------
  Pod Name = busybox (default)
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3 (default)
     Labels:
        Key   = pod-template-hash
        Value = 2716574283
        ----------
        Key   = run
        Value = curl
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
  ----------
  Pod Name = heapster-342135353-vrqjj (kube-system)
     Labels:
        Key   = k8s-app
        Value = heapster
        ----------
        Key   = pod-template-hash
        Value = 342135353
        ----------
     ----------
     Container Name = heapster
     Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
     ----------
     Container Name = heapster-nanny
     Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
  ----------
  Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
     ----------
     Container Name = kube-addon-manager
     Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
  ----------
  Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-apiserver
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-apiserver
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-controller-manager
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-controller-manager
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
     Labels:
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
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
     Labels:
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
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
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-g536m (kube-system)
     Labels:
        Key   = tier
        Value = node
        ----------
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-mdt7c (kube-system)
     Labels:
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
        Key   = component
        Value = kube-proxy
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-nj57n (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-scheduler-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-scheduler
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-scheduler
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kubernetes-dashboard-924040265-vjppn (kube-system)
     Labels:
        Key   = k8s-app
        Value = kubernetes-dashboard
        ----------
        Key   = pod-template-hash
        Value = 924040265
        ----------
     ----------
     Container Name = kubernetes-dashboard
     Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
  ----------
  Pod Name = tiller-deploy-677436516-2wvg1 (kube-system)
     Labels:
        Key   = app
        Value = helm
        ----------
        Key   = name
        Value = tiller
        ----------
        Key   = pod-template-hash
        Value = 677436516
        ----------
     ----------
     Container Name = tiller
     Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
----------
Node Name = k8s-master-dce67bdc-0
  ----------
  Pod Name = busybox (default)
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3 (default)
     Labels:
        Key   = pod-template-hash
        Value = 2716574283
        ----------
        Key   = run
        Value = curl
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
  ----------
  Pod Name = heapster-342135353-vrqjj (kube-system)
     Labels:
        Key   = k8s-app
        Value = heapster
        ----------
        Key   = pod-template-hash
        Value = 342135353
        ----------
     ----------
     Container Name = heapster
     Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
     ----------
     Container Name = heapster-nanny
     Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
  ----------
  Pod Name = kube-addon-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
     ----------
     Container Name = kube-addon-manager
     Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
  ----------
  Pod Name = kube-apiserver-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-apiserver
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-apiserver
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-controller-manager-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-controller-manager
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-controller-manager
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-dns-v20-3003781527-2djkm (kube-system)
     Labels:
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
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
     Labels:
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
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
     Labels:
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-g536m (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-mdt7c (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-nj57n (kube-system)
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-scheduler-k8s-master-dce67bdc-0 (kube-system)
     Labels:
        Key   = component
        Value = kube-scheduler
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-scheduler
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
----------
Namespace = kube-public
----------
Namespace = kube-system
  ----------
  Pod Name = busybox
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3
     Labels:
        Key   = pod-template-hash
        Value = 2716574283
        ----------
        Key   = run
        Value = curl
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
  ----------
  Pod Name = heapster-342135353-vrqjj
     Labels:
        Key   = k8s-app
        Value = heapster
        ----------
        Key   = pod-template-hash
        Value = 342135353
        ----------
     ----------
     Container Name = heapster
     Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
     ----------
     Container Name = heapster-nanny
     Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
  ----------
  Pod Name = kube-addon-manager-k8s-master-dce67bdc-0
     Labels:
     ----------
     Container Name = kube-addon-manager
     Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
  ----------
  Pod Name = kube-apiserver-k8s-master-dce67bdc-0
     Labels:
        Key   = tier
        Value = control-plane
        ----------
        Key   = component
        Value = kube-apiserver
        ----------
     ----------
     Container Name = kube-apiserver
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-controller-manager-k8s-master-dce67bdc-0
     Labels:
        Key   = component
        Value = kube-controller-manager
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-controller-manager
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-dns-v20-3003781527-2djkm
     Labels:
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
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
     Labels:
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
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
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-g536m
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-mdt7c
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-nj57n
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-scheduler-k8s-master-dce67bdc-0
     Labels:
        Key   = component
        Value = kube-scheduler
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-scheduler
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kubernetes-dashboard-924040265-vjppn
     Labels:
        Key   = pod-template-hash
        Value = 924040265
        ----------
        Key   = k8s-app
        Value = kubernetes-dashboard
        ----------
     ----------
     Container Name = kubernetes-dashboard
     Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
  ----------
  Pod Name = tiller-deploy-677436516-2wvg1
     Labels:
        Key   = app
        Value = helm
        ----------
        Key   = name
        Value = tiller
        ----------
        Key   = pod-template-hash
        Value = 677436516
        ----------
     ----------
     Container Name = tiller
     Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
----------
Namespace = staging
  ----------
  Pod Name = busybox
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3
     Labels:
        Key   = pod-template-hash
        Value = 2716574283
        ----------
        Key   = run
        Value = curl
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
  ----------
  Pod Name = heapster-342135353-vrqjj
     Labels:
        Key   = k8s-app
        Value = heapster
        ----------
        Key   = pod-template-hash
        Value = 342135353
        ----------
     ----------
     Container Name = heapster
     Image Name = gcrio.azureedge.net/google_containers/heapster-amd64:v1.4.2
     ----------
     Container Name = heapster-nanny
     Image Name = gcrio.azureedge.net/google_containers/addon-resizer:1.7
  ----------
  Pod Name = kube-addon-manager-k8s-master-dce67bdc-0
     Labels:
     ----------
     Container Name = kube-addon-manager
     Image Name = gcrio.azureedge.net/google_containers/kube-addon-manager-amd64:v6.4-beta.2
  ----------
  Pod Name = kube-apiserver-k8s-master-dce67bdc-0
     Labels:
        Key   = component
        Value = kube-apiserver
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-apiserver
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-controller-manager-k8s-master-dce67bdc-0
     Labels:
        Key   = component
        Value = kube-controller-manager
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-controller-manager
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-dns-v20-3003781527-2djkm
     Labels:
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
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
     Labels:
        Key   = pod-template-hash
        Value = 3003781527
        ----------
        Key   = version
        Value = v20
        ----------
        Key   = k8s-app
        Value = kube-dns
        ----------
        Key   = kubernetes.io/cluster-service
        Value = true
        ----------
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
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-g536m
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-mdt7c
     Labels:
        Key   = component
        Value = kube-proxy
        ----------
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-proxy-nj57n
     Labels:
        Key   = controller-revision-hash
        Value = 402111355
        ----------
        Key   = pod-template-generation
        Value = 1
        ----------
        Key   = tier
        Value = node
        ----------
        Key   = component
        Value = kube-proxy
        ----------
     ----------
     Container Name = kube-proxy
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kube-scheduler-k8s-master-dce67bdc-0
     Labels:
        Key   = component
        Value = kube-scheduler
        ----------
        Key   = tier
        Value = control-plane
        ----------
     ----------
     Container Name = kube-scheduler
     Image Name = gcrio.azureedge.net/google_containers/hyperkube-amd64:v1.7.7
  ----------
  Pod Name = kubernetes-dashboard-924040265-vjppn
     Labels:
        Key   = k8s-app
        Value = kubernetes-dashboard
        ----------
        Key   = pod-template-hash
        Value = 924040265
        ----------
     ----------
     Container Name = kubernetes-dashboard
     Image Name = gcrio.azureedge.net/google_containers/kubernetes-dashboard-amd64:v1.6.3
  ----------
  Pod Name = tiller-deploy-677436516-2wvg1
     Labels:
        Key   = app
        Value = helm
        ----------
        Key   = name
        Value = tiller
        ----------
        Key   = pod-template-hash
        Value = 677436516
        ----------
     ----------
     Container Name = tiller
     Image Name = gcrio.azureedge.net/kubernetes-helm/tiller:v2.6.1
  ----------
  Pod Name = flask-monolith
     Labels:
        Key   = app
        Value = mymonolith-app
        ----------
        Key   = name
        Value = mymonolith
        ----------
     ----------
     Container Name = mymonolith-container
     Image Name = brunoterkaly/monolith
----------
Namespace = default
  ----------
  Pod Name = busybox
     Labels:
        Key   = run
        Value = busybox
        ----------
     ----------
     Container Name = busybox
     Image Name = busybox
  ----------
  Pod Name = curl-2716574283-m72z3
     Labels:
        Key   = run
        Value = curl
        ----------
        Key   = pod-template-hash
        Value = 2716574283
        ----------
     ----------
     Container Name = curl
     Image Name = radial/busyboxplus:curl
done!

```
