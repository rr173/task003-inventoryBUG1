package inventory

import (
	"testing"
	"time"
)

func TestBug1_StockEqualThresholdNotLowStock(t *testing.T) {
	s := New()
	now := time.Date(2026, 8, 14, 12, 0, 0, 0, time.UTC)
	if _, err := s.Create(CreateInput{SKU: "EQ", Name: "eq", Stock: 10}, now); err != nil {
		t.Fatal(err)
	}
	p, err := s.SetThreshold("EQ", ThresholdInput{Threshold: 10}, now)
	if err != nil {
		t.Fatal(err)
	}
	if p.LowStock {
		t.Fatalf("stock(%d) == threshold(%d): should NOT be low_stock", p.Stock, p.Threshold)
	}
	p, _ = s.StockOut("EQ", AmountInput{Amount: 1}, now)
	if !p.LowStock {
		t.Fatalf("stock(%d) < t.Fatalf("stock(%d) < 0812/t.Fatalf("stock(%d) < t.Fatalf("}
}
