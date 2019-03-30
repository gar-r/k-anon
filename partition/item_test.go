package partition

import "testing"

func TestItem_Contains(t *testing.T) {
	p := &Item{item: 1}

	t.Run("contains item", func(t *testing.T) {
		i := 1
		if !p.Contains(i) {
			t.Errorf("expected partition to contain item: %v", i)
		}
	})

	t.Run("does not contain item", func(t *testing.T) {
		i := 2
		if p.Contains(i) {
			t.Errorf("expected partition not to contain item: %v", i)
		}
	})

}

func TestItem_ContainsPartition(t *testing.T) {
	p := &Item{item: 1}
	if p.ContainsPartition(NewSet()) {
		t.Errorf("expected false")
	}
}

func TestItem_Equals(t *testing.T) {

	p := &Item{item: 1}

	t.Run("equals other partition", func(t *testing.T) {
		q := &Item{item: 1}
		if !p.Equals(q) {
			t.Errorf("expected true")
		}
	})

	t.Run("does not equal other partition", func(t *testing.T) {
		q := &Item{item: 2}
		if p.Equals(q) {
			t.Errorf("expected false")
		}
	})

	t.Run("incompatible type", func(t *testing.T) {
		q := NewSet()
		if p.Equals(q) {
			t.Errorf("expected false")
		}
	})

}

func TestItem_String(t *testing.T) {
	q := &Item{item: 1}
	actual := q.String()
	expected := "1"
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
