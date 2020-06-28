package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
	mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification function")
	fmt.Println("Value passed in: %d\n", value)

	args := m.Called(value)

	return args.Bool(0)
}

func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

func TestChargeCustomer(t *testing.T){
	smsService := new(smsServiceMock)

	smsService.On("SendChargeNotification", 100).Return(true)

	myService := MyService{smsService}
	myService.ChargeCustomer(100)

	smsService.AssertExpectations(t)
}

//func TestCalculate(t *testing.T) {
//	assert := assert.New(t)
//
//	var tests = []struct{
//		input int 
//		expected int
//	}{
//		{2, 4},
//        {-1, 1},
//        {0, 2},
//        {-5, -3},
//        {99999, 100001},
//	}
//
//	for _, test := range tests {
//		assert.Equal(Caluculate(test.input), test.expected)
//	}
//}



//func TestCalculate(t *testing.T){
//	if Calculate(2) != 4 {
//		t.Error("Expecter 2 + 2 to equal 4")
//	}
//}
//
//func TestTableCalculate(t *testing.T){
//	var tests = []struct {
//		input int
//		expect int
//	}{
//		{2, 4},
//		{-1, 1},
//		{0, 2},
//		{-5, -3},
//		{99999, 100001},
//	}
//
//	for _, test := range tests {
//		if output := Calculate(test.input); output != test.expect {
//			t.Error("Test Failed: {} inputted, {} expected, recived: {}", test.input, test.expect, output)
//		}
//	}
//}