package appCart

import "testing"

func TestDuplicates(t *testing.T) {
	items := []CartItemDto{
		{Id: 1, Quantity: 1},
		{Id: 2, Quantity: 1},
		{Id: 3, Quantity: 1},
		{Id: 1, Quantity: 1},
	}

	err := checkForDuplicates(items)
	// Something here is looking funny, but whatever :p
	if err == nil {
		t.Error("should be duplicates")
	}
}

func TestNoDuplicates(t *testing.T) {
	items := []CartItemDto{
		{Id: 1, Quantity: 1},
		{Id: 2, Quantity: 1},
		{Id: 3, Quantity: 1},
	}

	err := checkForDuplicates(items)
	if err != nil {
		t.Error("should be no duplicates")
	}

}
