package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ybalexdp/go-generate-sample/domain/model"
	mock_repository "github.com/ybalexdp/go-generate-sample/mock/repository"
)

func TestSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApiClient := mock_repository.NewMockItemGetter(ctrl)

	expected := model.ItemModel{Id: 1, Name: "hoge"}
	mockApiClient.EXPECT().Get(1).Return(model.ItemModel{Id: 1, Name: "hoge"}, nil)

	res, err := mockApiClient.Get(1)
	if err != nil {
		t.Errorf("error! %v", err)
	}
	if res != expected {
		t.Errorf("Get() = %v want %v", res, expected)
	}
}
