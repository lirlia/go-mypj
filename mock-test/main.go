package main

type ApiClient interface {
	Request(string) (string, error)
}

type RegisterInterface interface {
	Register(string) (string, error)
}

// データを登録する
type DataRegister struct {
	client ApiClient
}

var _ ApiClient = &DataRegister{}
var _ RegisterInterface = &DataRegister{}

func (d *DataRegister) Register(data string) (string, error) {
	result, err := d.client.Request(data)
	if err != nil {
		return "", err
	}
	return result, nil
}

func main() {

	d := &DataRegister{}
	a, err := d.client.Request("GET")
	v, err := d.Request("GET")
}
