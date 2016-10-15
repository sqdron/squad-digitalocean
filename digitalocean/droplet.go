package digitalocean

type dropletApi struct {

}

func DropletUnit() *dropletApi {
	return &dropletApi{}
}

//func (d *dropletApi) Create() (model.Unit, error){
//
//}
//func (d *dropletApi) List() ([]*model.Unit, error)
//func (d *dropletApi) Delete(model.Unit) error