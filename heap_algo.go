package main

import (
	"fmt"
	"sort"
)

type Interface interface {
    		// Len is the number of elements in the collection.
    		Len() int
    		// Less reports whether the element with
    		// index i should sort before the element with index j.
    		Less(i, j int) bool
    		// Swap swaps the elements with indexes i and j.
    		Swap(i, j int)
}    	

// maintain max-heap property from index
func siftDown(data [][]byte, lo, hi, first int) {
  root := lo
  for {
    child := 2*root + 1
    if child >= hi {
      break
    }
    if child+1 < hi && less(data[first+child], data[first+child+1]) {
      child++
    }
    if !less(data[first+root], data[first+child]) {
      return
    }
    swap(data, first+root, first+child)
    root = child
  }
}

func less(a []byte, b []byte) bool {
    // I assume a and b have the same length
    n := 0
    for n<len(a) {
      if a[n] < b[n] {
        return true
      }
      if a[n] > b[n] {
        return false
      }
      n++
    }
    return false
}

// swap the index a and b of array
func swap(data [][]byte, a, b int) {
  temp := data[a]
  data[a] =  data[b]
  data[b] = temp
}

func HeapSort(data [][]byte, a, b int) {
  first := a
  lo := 0
  hi := b - a
  // build max-heap from index a to index b
  for i := (hi - 1) / 2; i >= 0; i-- {
    siftDown(data, i, hi, first)
  }
  // pop maximum to the end and rebuild the former elements
  for i := hi - 1; i >= 0; i-- {
    swap(data, first, first+i)
    siftDown(data, lo, i, first)
  }
}
func Sort(data Interface) {
   		// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
   		n := data.Len()
   		maxDepth := 0
   		for i := n; i > 0; i >>= 1 {
   			maxDepth++
   		}
   		maxDepth *= 2
   		quickSort(data, 0, n, maxDepth)
   	}

func main{
	 var arr := [}int{31,42,17,26}

	fmt.Println(arr)
	sort.Sort(arr)
	fmt.Println(arr)

}
