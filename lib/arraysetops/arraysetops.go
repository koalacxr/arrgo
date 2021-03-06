package arraysetops

import (
	"github.com/ledao/arrgo"
	"sort"
)

//Find the unique elements of an array.
//
//Returns the sorted unique elements of an array. There are three optional
//outputs in addition to the unique elements: the indices of the input array
//that give the unique values, the indices of the unique array that
//reconstruct the input array, and the number of times each unique value
//comes up in the input array.
//
//Parameters
//----------
//ar : array_like
//Input array. This will be flattened if it is not already 1-D.
//return_index : bool, optional
//If True, also return the indices of `ar` that result in the unique
//array.
//return_inverse : bool, optional
//If True, also return the indices of the unique array that can be used
//to reconstruct `ar`.
//return_counts : bool, optional
//If True, also return the number of times each unique value comes up
//in `ar`.
//
//.. versionadded:: 1.9.0
//
//Returns
//-------
//unique : ndarray
//The sorted unique values.
//unique_indices : ndarray, optional
//The indices of the first occurrences of the unique values in the
//(flattened) original array. Only provided if `return_index` is True.
//unique_inverse : ndarray, optional
//The indices to reconstruct the (flattened) original array from the
//unique array. Only provided if `return_inverse` is True.
//unique_counts : ndarray, optional
//The number of times each of the unique values comes up in the
//original array. Only provided if `return_counts` is True.
func Unique(a *arrgo.Arrf) *arrgo.Arrf {
	uniques := make([]float64, 0, a.Count())
	for _, v := range a.Values() {
		if !arrgo.ContainsFloat64(uniques, v) {
			uniques = append(uniques, v)
		}
	}
	sort.Float64s(uniques)
	return arrgo.Array(uniques)
}
