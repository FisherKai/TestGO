package utils

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
)

// ConnectToVSphere 连接到 vSphere
func ConnectToVSphere(vcenter, username, password string) (*govmomi.Client, error) {
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

// FindClusters 查找所有集群
func FindClusters(client *govmomi.Client) ([]*object.ClusterComputeResource, error) {
	ctx := context.Background()
	finder := find.NewFinder(client.Client, true)

	// 获取根文件夹
	rootFolder, err := finder.Folder(ctx, "/")
	if err != nil {
		return nil, err
	}

	return findClusters(ctx, rootFolder)
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

// PrintClusterStatus 打印集群状态
func PrintClusterStatus(cluster *object.ClusterComputeResource) {
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
	fmt.Printf("Cluster OverallStatus: %s\n", clusterInfo.Summary.GetComputeResourceSummary().OverallStatus)
	fmt.Println("-----------------------------")
}
