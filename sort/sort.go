package sort

import "github.com/go-go-data-structs/utils"

func BubbleSort(s []int) []int {
    iters := 0

    for iters < (len(s) - 1) {
        for i,x := range s[:(len(s)-1)] {
            elementTwo := s[i+1]

            if x > elementTwo {
                utils.SwapFromSlice(s, i)
            }
        }

        iters+=1
    }

    return s
}

