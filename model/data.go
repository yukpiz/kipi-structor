package model

type Data struct {
	Texts []Text
	Items []Item
}

func (d Data) FindFromId(id int) Text {
	for _, text := range d.Texts {
		if text.TextId == id {
			return text
		}
	}
	return Text{}
}

func (d Data) FindFromData(name string) {
	//TODO: Like search
}
