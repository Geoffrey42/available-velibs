package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("returns total number of available bikes", func(t *testing.T) {
		records := []Records{
			{
				Fields{
					NumBikesAvailable: 6,
				},
			},
			{
				Fields{
					NumBikesAvailable: 4,
				},
			},
		}

		got := GlobalResponse{
			Total:   0,
			NHits:   9,
			Records: records,
		}
		got.Sum()

		want := 10

		if got.Total != want {
			t.Errorf("got %d, want %d", got.Total, want)
		}
	})
}
