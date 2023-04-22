package sort

import "testing"

var sliceToTestSortFromLeastToGreatest = []int{10,9,8,7,6,5,4,3,2,1}
var expedtedLeastToGreatestSortedSlice = []int{1,2,3,4,5,6,7,8,9,10}

func Test_bubbleSort (t *testing.T) {
    actualSliceResult := BubbleSort(sliceToTestSortFromLeastToGreatest)

    for i,x := range expedtedLeastToGreatestSortedSlice {
        if x != actualSliceResult[i] {
            t.Errorf(
                "Expected bubble sort to have value %d in index %d. Instead it returns %d.\n",
                expedtedLeastToGreatestSortedSlice[i],
                i,
                actualSliceResult[i],
            )
        }
    }
}
