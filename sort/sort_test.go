package sort

import "testing"

func Test_bubbleSort (t *testing.T) {
    s := []int{}

    result := BubbleSort(&s)

    if result != 0 {
        t.Error("The test failed with result ", result)
    }
}
