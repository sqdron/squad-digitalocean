package digitalocean

import (
	"github.com/digitalocean/godo"
	"github.com/sqdron/squad-cloud/cloud"
	"fmt"
)

type keysService struct {

}

func (d *keysService) client(token string) godo.KeysService {
	return NewClient(token).Keys
}

func KeysService() cloud.ICloudKey {
	return &keysService{}
}

func (d *keysService) List(r *cloud.TokenSource) ([]*cloud.CloudKey, error) {
	fmt.Println(r.Token)
	keys, resp, err := d.client(r.Token).List(&godo.ListOptions{Page: 1, PerPage: 50})
	fmt.Println(keys)
	fmt.Println(resp)
	fmt.Println(err)
	if (err != nil) {
		return nil, err
	}
	result := []*cloud.CloudKey{}
	for _, i := range keys {
		result = append(result, toKey(&i))
	}
	return result, nil
}

func (d *keysService) Get(r *cloud.KeyGetRequest) (*cloud.CloudKey, error) {
	key, resp, err := d.client(r.Token).GetByID(r.ID)
	if (err != nil) {
		return nil, err
	}
	fmt.Println(key)
	fmt.Println(resp)
	fmt.Println(err)
	return toKey(key), nil

}

func (d *keysService) Create(r *cloud.KeyCreateRequest) (*cloud.CloudKey, error) {
	key, resp, err := d.client(r.Token).Create(&godo.KeyCreateRequest{Name:r.Name, PublicKey:r.PublicKey})
	if (err != nil) {
		return nil, err
	}
	fmt.Println(key)
	fmt.Println(resp)
	fmt.Println(err)
	return toKey(key), nil

}

func toKey(i *godo.Key) *cloud.CloudKey {
	return &cloud.CloudKey{ID:i.ID, Name:i.Name, Fingerprint:i.Fingerprint, PublicKey:i.PublicKey}
}
