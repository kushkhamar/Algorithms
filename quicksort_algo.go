func medianOfThree(data Interface, a, b, c int) {
    80		m0 := b
    81		m1 := a
    82		m2 := c
    83		// bubble sort on 3 elements
    84		if data.Less(m1, m0) {
    85			data.Swap(m1, m0)
    86		}
    87		if data.Less(m2, m1) {
    88			data.Swap(m2, m1)
    89		}
    90		if data.Less(m1, m0) {
    91			data.Swap(m1, m0)
    92		}
    93		// now data[m0] <= data[m1] <= data[m2]
    94	}
    95	
    96	func swapRange(data Interface, a, b, n int) {
    97		for i := 0; i < n; i++ {
    98			data.Swap(a+i, b+i)
    99		} 
   100	}
   101	
   102	func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
   103		m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
   104		if hi-lo > 40 {
   105			// Tukey's ``Ninther,'' median of three medians of three.
   106			s := (hi - lo) / 8
   107			medianOfThree(data, lo, lo+s, lo+2*s)
   108			medianOfThree(data, m, m-s, m+s)
   109			medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
   110		}
   111		medianOfThree(data, lo, m, hi-1)
   112	
   113		// Invariants are:
   114		//	data[lo] = pivot (set up by ChoosePivot)
   115		//	data[lo <= i < a] = pivot
   116		//	data[a <= i < b] < pivot
   117		//	data[b <= i < c] is unexamined
   118		//	data[c <= i < d] > pivot
   119		//	data[d <= i < hi] = pivot
   120		//
   121		// Once b meets c, can swap the "= pivot" sections
   122		// into the middle of the slice.
   123		pivot := lo
   124		a, b, c, d := lo+1, lo+1, hi, hi
   125		for {
   126			for b < c {
   127				if data.Less(b, pivot) { // data[b] < pivot
   128					b++
   129				} else if !data.Less(pivot, b) { // data[b] = pivot
   130					data.Swap(a, b)
   131					a++
   132					b++
   133				} else {
   134					break
   135				}
   136			}
   137			for b < c {
   138				if data.Less(pivot, c-1) { // data[c-1] > pivot
   139					c--
   140				} else if !data.Less(c-1, pivot) { // data[c-1] = pivot
   141					data.Swap(c-1, d-1)
   142					c--
   143					d--
   144				} else {
   145					break
   146				}
   147			}
   148			if b >= c {
   149				break
   150			}
   151			// data[b] > pivot; data[c-1] < pivot
   152			data.Swap(b, c-1)
   153			b++
   154			c--
   155		}
   156	
   157		n := min(b-a, a-lo)
   158		swapRange(data, lo, b-n, n)
   159	
   160		n = min(hi-d, d-c)
   161		swapRange(data, c, hi-n, n)
   162	
   163		return lo + b - a, hi - (d - c)
   164	}
   165	
   166	func quickSort(data Interface, a, b, maxDepth int) {
   167		for b-a > 7 {
   168			if maxDepth == 0 {
   169				heapSort(data, a, b)
   170				return
   171			}
   172			maxDepth--
   173			mlo, mhi := doPivot(data, a, b)
   174			// Avoiding recursion on the larger subproblem guarantees
   175			// a stack depth of at most lg(b-a).
   176			if mlo-a < b-mhi {
   177				quickSort(data, a, mlo, maxDepth)
   178				a = mhi // i.e., quickSort(data, mhi, b)
   179			} else {
   180				quickSort(data, mhi, b, maxDepth)
   181				b = mlo // i.e., quickSort(data, a, mlo)
   182			}
   183		}
   184		if b-a > 1 {
   185			insertionSort(data, a, b)
   186		}
   187	}
