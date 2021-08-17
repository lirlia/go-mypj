package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/lirlia/go-mypj/mock-test/mock"
)

func TestSample(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApiClient := mock_main.NewMockApiClient(ctrl)
	mockApiClient.EXPECT().Request("bar").Return("bar", nil)

	d := &DataRegister{}
	d.client = mockApiClient

	expected := "bar"

	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("register error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}
