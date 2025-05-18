package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UsecaseTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	mockRepo *MockRepository
	usecase  Usecase
}

func (suite *UsecaseTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRepo = NewMockRepository(suite.mockCtrl)
	suite.usecase = NewUsecase(suite.mockRepo)
}

func (suite *UsecaseTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func TestUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UsecaseTestSuite))
}

func (suite *UsecaseTestSuite) TestCreateProduct() {
	type args struct {
		request CreateProductInput
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "valid product input",
			args: args{
				request: CreateProductInput{
					Name:          "Test Product",
					Description:   "Test Description",
					OriginalPrice: 100,
					Discounts: []DiscountInput{
						{Type: "fixed", Value: 100},
					},
				},
			},
			mock: func() {
				suite.mockRepo.EXPECT().CreateProduct(gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed create product",
			args: args{
				request: CreateProductInput{
					Name:          "Test Product",
					Description:   "Test Description",
					OriginalPrice: 100,
					Discounts: []DiscountInput{
						{Type: "fixed", Value: 100},
					},
				},
			},
			mock: func() {
				suite.mockRepo.EXPECT().CreateProduct(gomock.Any()).Return(errors.New("failed create product"))
			},
			wantErr: true,
		},
		{
			name: "invalid product price",
			args: args{
				request: CreateProductInput{
					Name:          "Test Product",
					Description:   "Test Description",
					OriginalPrice: 0,
					Discounts: []DiscountInput{
						{Type: "fixed", Value: 100},
					},
				},
			},
			mock:    func() {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := suite.usecase.CreateProduct(tt.args.request)

			if tt.wantErr {
				suite.Error(err)
				suite.Nil(result)
			} else {
				suite.NoError(err)
				suite.NotNil(result)
				suite.Contains(result, "product")
				suite.Contains(result, "appliedDiscounts")
			}
		})
	}
}

func (suite *UsecaseTestSuite) TestListProducts() {

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "valid list products",
			mock: func() {
				suite.mockRepo.EXPECT().ListProducts().Return([]Product{}, nil)
			},
			wantErr: false,
		},
		{
			name: "failed list products",
			mock: func() {
				suite.mockRepo.EXPECT().ListProducts().Return(nil, errors.New("failed list products"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := suite.usecase.ListProducts()

			if tt.wantErr {
				suite.Error(err)
				suite.Nil(result)
			} else {
				suite.NoError(err)
				suite.NotNil(result)
				suite.Equal(result, []Product{})
			}
		})
	}
}
