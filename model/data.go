package model

type Data struct {
	Items []Item
}

func (d Data) FindFromId(id int) Item {
	for _, item := range d.Items {
		if item.Id == id {
			return item
		}
	}
	return Item{}
}

func (d Data) FindFromData(name string) {
	//TODO: Like search
}
