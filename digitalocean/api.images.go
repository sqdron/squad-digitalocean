package digitalocean

import (
	"github.com/digitalocean/godo"
	"github.com/sqdron/squad-cloud/cloud"
	"fmt"
)

type imageService struct {

}

func (d *imageService) client(token string) godo.ImagesService {
	return NewClient(token).Images
}

func ImageService() cloud.IImage {
	return &imageService{}
}

func (d *imageService) List(r *cloud.TokenSource) ([]*cloud.Image, error) {
	fmt.Println(r.Token)
	images, resp, err := d.client(r.Token).List(&godo.ListOptions{Page: 1, PerPage: 50})
	fmt.Println(images)
	fmt.Println(resp)
	fmt.Println(err)
	imgs := []*cloud.Image{}
	for _, i := range images{
		imgs = append(imgs, toImage(&i))
	}
	return imgs, err
}

func toImage(i *godo.Image) *cloud.Image{
	return &cloud.Image{ID:i.ID, Name:i.Name, Distribution:i.Distribution, Slug:i.Slug}
}