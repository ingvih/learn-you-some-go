package shapes

import "testing"

func TestArea(t *testing.T) {
	got := Area(4.0, 5.0)
	want := 20.0

	if(got != want){
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
