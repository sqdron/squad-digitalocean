package digitalocean

import (
	"github.com/digitalocean/godo"
	"github.com/sqdron/squad-cloud/cloud"
	"fmt"
)

type sizeService struct {

}

func (d *sizeService) client(token string) godo.SizesService {
	return NewClient(token).Sizes
}

func PlanService() cloud.ICloudPlan {
	return &sizeService{}
}

func (d *sizeService) List(r *cloud.TokenSource) ([]*cloud.CloudPlan, error) {
	fmt.Println(r.Token)
	sizes, resp, err := d.client(r.Token).List(&godo.ListOptions{Page: 1, PerPage: 50})
	fmt.Println(sizes)
	fmt.Println(resp)
	fmt.Println(err)
	result := []*cloud.CloudPlan{}
	for _, i := range sizes{
		result = append(result, toCloudPlan(&i))
	}
	return result, err
}

func toCloudPlan(i *godo.Size) *cloud.CloudPlan{
	return &cloud.CloudPlan{Name:i.Slug, Vcpus:i.Vcpus, Memory:i.Memory, Disk:i.Disk, PriceMonthly:i.PriceMonthly, PriceHourly:i.PriceHourly, Regions:i.Regions}
}