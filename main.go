package main

import (
    "flag"
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
    "os"
    v1types "k8s.io/client-go/kubernetes/typed/core/v1"
    "k8s.io/api/core/v1"
    "sync"
)

type MyPods []MyPod

// A collection of pods organized by node and by namespace
type PodsByNode map[string]MyPods
type PodsByNamespace map[string]MyPods

type Data struct {
    podsByNode      PodsByNode
    podsByNamespace PodsByNamespace
}

type MyPod struct {
    ID         string
    namespace  string
    containers []string
    images     []string
    labels     map[string]string
}

func main() {
    //newPodsByNode := make(map[string]Pods)
    //newPodsByNamespace := make(map[string]Pods)

    var mynodes v1types.NodeInterface
    var err error
    var list *v1.NodeList

    config := getKubeConfig()

    client := kubernetes.NewForConfigOrDie(config)

    mynodes = client.CoreV1().Nodes()
    list, err = mynodes.List(metav1.ListOptions{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "error listing nodes: %v", err)
        os.Exit(1)
    }

    data := Data { podsByNode: make(map[string]MyPods), podsByNamespace: make(map[string]MyPods) }
    //newPodsByNode := make(map[string]Pods)
    //newPodsByNamespace := make(map[string]Pods)

    for _, node := range list.Items {
        // Found a node, so record the node name
        // Later we'll add to the collection of pods
        data.podsByNode[node.Name] = MyPods{}
        fmt.Printf("Node: %s\n", node.Name)
    }

    // Iterate through namespaces
    namespacelist, err := client.CoreV1().Namespaces().List(metav1.ListOptions{})
    for _, ns := range namespacelist.Items {
        // Found a namespace, so record the node name
        // Later we'll add to the collection of pods
        data.podsByNamespace[ns.Name] = MyPods{}
        fmt.Printf("Namespace: %s\n", ns.Name)
    }
    var wg sync.WaitGroup
    wg.Add(2)
    go mapPodsToNodes(data, client, &wg)
    go mapPodsToNamespaces(data, client, &wg)
    wg.Wait()
    showPodsByNodes(data)
    showPodsByNamespace(data)
    fmt.Println("done!")

}
//-------------------------------------------------------------------------
//
//-------------------------------------------------------------------------
func showPodsByNodes(data Data) {
    for key, _ := range data.podsByNode {
        fmt.Println("----------")
        fmt.Printf("Node Name = %v\n", key)
        // Find matching key by looping through all pods
        for _, poditem := range data.podsByNode[key] {
            fmt.Println("  ----------")
            fmt.Printf("  Pod Name = %v (%v)\n", poditem.ID, poditem.namespace)

            fmt.Printf("     Labels:\n")
            for k, v := range poditem.labels {
                fmt.Printf("        Key   = %v\n", k)
                fmt.Printf("        Value = %v\n", v)
                fmt.Println("        ----------")
            }
            for i := range poditem.containers {
                fmt.Println("     ----------")
                fmt.Printf("     Container Name = %v\n", poditem.containers[i])
                fmt.Printf("     Image Name = %v\n", poditem.images[i])
            }

        }
    }
}
//-------------------------------------------------------------------------
//
//-------------------------------------------------------------------------
func showPodsByNamespace(data Data) {
    for key, _ := range data.podsByNamespace{
        fmt.Println("----------")
        fmt.Printf("Namespace = %v\n", key)
        // Find matching key by looping through all pods
        for _, poditem := range data.podsByNamespace[key] {
            fmt.Println("  ----------")
            fmt.Printf("  Pod Name = %v\n", poditem.ID)
            fmt.Printf("     Labels:\n")
            for k, v := range poditem.labels {
                fmt.Printf("        Key   = %v\n", k)
                fmt.Printf("        Value = %v\n", v)
                fmt.Println("        ----------")
            }
            for i := range poditem.containers {
                fmt.Println("     ----------")
                fmt.Printf("     Container Name = %v\n", poditem.containers[i])
                fmt.Printf("     Image Name = %v\n", poditem.images[i])
            }
        }
    }
}
//-------------------------------------------------------------------------
//
//-------------------------------------------------------------------------
func mapPodsToNodes(data Data, client *kubernetes.Clientset, wg *sync.WaitGroup) {

    defer wg.Done()
    // Get a list the pods

    podlist, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "error listing pods %v", err)
        os.Exit(1)
    }

    // ----------------------------------------------------------------
    // Loop through pods and and map to node names
    // ----------------------------------------------------------------
    var myPods MyPods
    for _, poditem := range podlist.Items {

        // Got a pod
        newMyPod := MyPod{ID: poditem.Name, namespace: poditem.Namespace}
        // Each pod has 1 or more containers with images
        for _, containeritem := range poditem.Spec.Containers {
            // For each pod, get all the containers and images
            newMyPod.containers = append(newMyPod.containers, containeritem.Name)
            newMyPod.images = append(newMyPod.images, containeritem.Image)
        }
        newMyPod.labels = poditem.Labels
        // Add to collection of pods
        myPods = append(myPods, newMyPod)
        // Add collection to podsByNode map, using NodeName as map key
        data.podsByNode[poditem.Spec.NodeName] = myPods
    }
}

//-------------------------------------------------------------------------
//
//-------------------------------------------------------------------------
func mapPodsToNamespaces(data Data, client *kubernetes.Clientset, wg *sync.WaitGroup) {
    defer wg.Done()
    var myPods MyPods
    podlist, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "error listing pods %v", err)
        os.Exit(1)
    }
    // Loop through all namespace keys (Namespace) in map
    for key, _ := range data.podsByNamespace {
        // Find matching key by looping through all pods
        for _, poditem := range podlist.Items {
            // If current pod has matching namespace, then add to pod collection
            if poditem.Namespace == key {
                // Found match so get info about current pod
                newMyPod := MyPod{ID: poditem.Name, namespace: poditem.Namespace}
                // Add all containers and images
                for _, containeritem := range poditem.Spec.Containers {
                    newMyPod.containers = append(newMyPod.containers, containeritem.Name)
                    //fmt.Println(containeritem.Name)
                    newMyPod.images = append(newMyPod.images, containeritem.Image)
                }
                newMyPod.labels = poditem.Labels


                // Add collection to map using Namespace as key
                myPods = append(myPods, newMyPod)
                data.podsByNamespace[poditem.Namespace] = myPods
            }
        }
    }
}
//-------------------------------------------------------------------------
//
//-------------------------------------------------------------------------
func getKubeConfig() *rest.Config {
    kubeconfig := ""
    flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
    flag.Parse()

    if kubeconfig == "" {
        kubeconfig = os.Getenv("KUBECONFIG")
        if kubeconfig == "" {
            kubeconfig = "/root/.kube/config"
        }
    }

    var (
        config *rest.Config
        err    error
    )

    if kubeconfig != "" {
        config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
    } else {
        config, err = rest.InClusterConfig()
    }
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating client: %v", err)
        os.Exit(1)
    }
    return config

}
