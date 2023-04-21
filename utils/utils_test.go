package utils

import "testing"

func Test_swapFromSlice (t *testing.T) {
    var expectedIntSlice = []int{1,0}
    actualIntSlice := SwapFromSlice(
        []int{0,1},
        0,
    );

    for i, x := range expectedIntSlice {
        if x != actualIntSlice[i] {
            t.Errorf(
                "The Swap failed, expected %d but got %d.\n",
                x,
                actualIntSlice[i],
            )
        }
    }
}

