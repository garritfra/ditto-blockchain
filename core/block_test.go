package core

import (
	"testing"
)

func TestBlock_AddTransaction(t *testing.T) {
	block := Block{}
	transaction := Transaction{Sender: "foo", Receiver: "bar", Message: "Test", Amount: 100}
	block.AddTransaction(transaction)

	want := transaction
	got := block.Data[0]

	if want != got {
		t.Errorf("Want: %v, but got %v", want, got)
	}
}
