package main

type ApiClient interface {
	Request(string) (string, error)
}

type DataRegister struct {
	client ApiClient
}

func (d *DataRegister) Register(data string) (string, error) {
	result, err := d.client.Request(data)
	if err != nil {
		return "", err
	}
	return result, nil
}
