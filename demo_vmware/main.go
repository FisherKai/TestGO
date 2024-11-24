package main

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"log"
	"net/url"
	"os"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Application started")
	// vSphere 连接信息
	vcenter := "10.104.241.106"
	username := "administrator@vsphere.local"
	password := "Xxwlk@72916"

	// 创建连接
	client, err := connectToVSphere(vcenter, username, password)
	if err != nil {
		log.Fatalf("Failed to connect to vSphere: %v", err)
	}
	defer client.Logout(context.Background())

	// 获取查找器
	finder := find.NewFinder(client.Client, true)

	// 获取根文件夹
	rootFolder, err := finder.Folder(context.Background(), "/")
	if err != nil {
		log.Fatalf("Failed to get root folder: %v", err)
	}

	// 查找所有集群
	clusters, err := findClusters(context.Background(), rootFolder)
	if err != nil {
		log.Fatalf("Failed to find clusters: %v", err)
	}

	// 打印每个集群的状态
	for _, cluster := range clusters {
		printClusterStatus(cluster)
	}
}

func connectToVSphere(vcenter, username, password string) (*govmomi.Client, error) {
	ctx := context.Background()
	vsphereURL := fmt.Sprintf("https://%s/sdk", vcenter)
	u, err := url.Parse(vsphereURL)
	if err != nil {
		return nil, err
	}
	client, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		return nil, err
	}

	// 创建 Userinfo 对象
	userInfo := url.UserPassword(username, password)

	// 登录到 vSphere
	err = client.SessionManager.Login(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func findClusters(ctx context.Context, rootFolder *object.Folder) ([]*object.ClusterComputeResource, error) {
	var clusters []*object.ClusterComputeResource

	// 查找所有数据中心
	datacenters, err := rootFolder.Children(ctx)
	if err != nil {
		return nil, err
	}

	for _, dc := range datacenters {
		dcFolder, ok := dc.(*object.Folder)
		if !ok {
			continue
		}

		// 查找所有集群
		dcClusters, err := dcFolder.Children(ctx)
		if err != nil {
			return nil, err
		}
		for _, entity := range dcClusters {
			if cluster, ok := entity.(*object.ClusterComputeResource); ok {
				clusters = append(clusters, cluster)
			}
		}
	}

	return clusters, nil
}

func printClusterStatus(cluster *object.ClusterComputeResource) {
	ctx := context.Background()

	// 获取集群的详细信息
	var clusterInfo mo.ClusterComputeResource
	err := cluster.Properties(ctx, cluster.Reference(), []string{"name", "summary"}, &clusterInfo)
	if err != nil {
		log.Printf("Failed to get cluster properties: %v", err)
		return
	}

	// 打印集群名称和状态
	fmt.Printf("Cluster Name: %s\n", clusterInfo.Name)
	fmt.Printf("Cluster Summary: %+v", clusterInfo.Summary)
	fmt.Println("-----------------------------")
}
