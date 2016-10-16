package digitalocean

import (
	"github.com/sqdron/squad-cloud/cloud"
	"github.com/digitalocean/godo"
	"fmt"
)

type droplet struct {

}

func (d *droplet) client(token string) godo.DropletsService {
	return NewClient(token).Droplets
}

func UnitService() cloud.ICloudUnit {
	return &droplet{}
}

func (d *droplet) Create(r *cloud.UnitCreateRequest) (*cloud.CloudUnit, error) {
	createRequest := &godo.DropletCreateRequest{
		Name:   r.Name,
		Region: r.Region,
		Size:   r.Size,
		Image: godo.DropletCreateImage{
			Slug: r.Image,
		},
		IPv6: true,
	}
	droplet, req, err := d.client(r.Token).Create(createRequest)
	fmt.Println(droplet)
	fmt.Println(req)
	fmt.Println(err)
	if (err != nil){
		return nil, err
	}
	return toUnit(droplet), nil
}

func (d *droplet) List(r *cloud.TokenSource) ([]*cloud.CloudUnit, error) {
	droplets, _, err := d.client(r.Token).List(&godo.ListOptions{Page: 1, PerPage: 10})
	if (err != nil){
		return nil, err
	}
	units := []*cloud.CloudUnit{}
	for _, d := range droplets{
		units = append(units, toUnit(&d))
	}
	return units, nil
}

func (d *droplet) Delete(*cloud.CloudUnit) (error) {
	return nil
}

func toUnit(d *godo.Droplet) *cloud.CloudUnit {
	return &cloud.CloudUnit{Name:d.Name}
}