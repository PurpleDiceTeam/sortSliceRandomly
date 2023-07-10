package sortSliceRandomly

import (
	"math/rand"
)

func SortSliceRandomly[T any](slice []T, newSlice ...T) []T {
	/* If original slice len == 0 */
	if len(slice) == 0 {

		/* Stop function and return the new slice */
		return newSlice

		/* If original slice > 0 */
	} else {
		/* Get random position in original array */
		radomPosition := rand.Intn(len(slice))

		/* Add random element for the original slice on the new slice */
		newSlice = append(newSlice, slice[radomPosition])

		/* Delete the element in the original slice */
		slice = append(slice[:radomPosition], slice[radomPosition+1:]...)

		/* Re-invoke function */
		return SortSliceRandomly(slice, newSlice...)
	}
}
