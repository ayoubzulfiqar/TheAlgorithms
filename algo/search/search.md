# Searching Algorithms

Searching algorithms are a fundamental part of computer science and data retrieval. They are used to locate specific items within a collection of data, such as an array, list, or database. These algorithms play a crucial role in various applications, including information retrieval, data processing, and path-finding. Here's a Markdown table summarizing the advantages and disadvantages of searching algorithms:

| Algorithm                  | Advantages                                                                                                          | Disadvantages                                                                                                      |
|----------------------------|--------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| Linear Search              | - Simple and easy to understand.<br> - Suitable for small datasets.<br> - No need for data to be sorted.        | - Inefficient for large datasets (O(n) time complexity).<br> - Not optimal for sorted data.                      |
| Binary Search              | - Efficient for large sorted datasets (O(log n) time complexity).<br> - Reduces the search space by half.        | - Requires data to be sorted initially.<br> - Not suitable for unsorted data.<br> - May not find all occurrences. |
| Interpolation Search       | - Faster for uniformly distributed data.<br> - O(log log n) average time complexity for uniformly distributed data. | - Not effective for non-uniformly distributed data.<br> - Requires data to be sorted initially.                |
| Exponential Search         | - Useful when the size of the data is unknown.<br> - O(log n) time complexity for sorted data.                  | - Requires data to be sorted initially.<br> - Inefficient for very large datasets.                               |
| Jump Search                | - Efficient for sorted datasets.<br> - Reduces the number of comparisons.<br> - O(√n) time complexity.           | - Requires data to be sorted initially.<br> - Not suitable for very small or unsorted datasets.                   |
| Fibonacci Search           | - Efficient for sorted datasets.<br> - O(log n) time complexity.<br> - Reduces the search space systematically. | - Requires data to be sorted initially.<br> - Slightly more complex to implement.                              |
| Ternary Search             | - Useful for finding a single value in a sorted dataset.<br> - O(log3 n) time complexity.                        | - Requires data to be sorted initially.<br> - Not suitable for very large datasets.                             |
| Sentinel Linear Search     | - Simplifies the code by eliminating the need to check for array boundaries explicitly.<br> - Linear time complexity (O(n)). | - Still has linear time complexity.<br> - Not suitable for very large datasets.                                |
| Meta Binary Search         | - Improves the binary search algorithm's average case by reducing unnecessary comparisons.<br> - O(log n) time complexity for the average case. | - Slightly more complex to implement.<br> - May not always perform better in the worst case.                    |
| Ubiquitous Binary Search   | - Combines the advantages of binary search and linear search for specific scenarios.<br> - Better performance in certain situations. | - More complex to implement than binary or linear search.<br> - May not be suitable for all use cases.           |

We Will Implement the Generic Version of the Searching Algorithm Where It will work on different type of data e.g floats, ints, uint etc

## Numeric Types

Implementing Numeric Type of String is also possible and most of the Search Algorithms work with it, but Some f them (jump, interpolation) needed little bit modifications but others work perfectly

```go
type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
```

### Array vs Slice

I know the the difference of **Array** and **Slice** in GO, The thing that we are using inside the each and every search algorithm is basically a slice. So don't complain about it if the i name is as arr or array inside the function

```go

// This is Array - In Go array it is compulsory to have specified length which is 5 like it will have 5 possible values
// It always always have fix size
var array [5]int = [5]int{1, 2, 3, 4, 5}

// This is Slice -NO Specific Length needed
var	slice []int= []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// another way to make slice
var slice []int = make([]int, 5)

```

## Linear Search

Linear search is a simple searching algorithm that finds the position of a target element within a list or array. It works by sequentially checking each element in the list until a match is found or the entire list has been searched.

1. Start at the beginning of the list.
2. Compare the target element with the current element in the list.
3. If the current element matches the target element, return the index of the current element (or its position).
4. If the current element does not match the target element, move to the next element in the list.
5. Repeat steps 2-4 until the target element is found or the end of the list is reached.

If the target element is not found in the list, the linear search returns a "not found" indication.

Real-life example:
Imagine you have a shelf full of books, and you want to find a specific book with a particular title. You start at one end of the shelf and go through each book, one by one, checking the title. If you find the book you're looking for, you stop and note its position. If you've checked all the books without finding the one you're looking for, you conclude that it's not on the shelf.

### Code

This Type of Linear does not care about the multiple of duplicate values if present inside the array, Whatever the element it found or math at the beginning it wil return it,

```go
// Iterative
func Linear[T Numeric](array []T, target T) T {
	for i, value := range array {
		if value == target {
			fmt.Printf("Found %v at index %v\n", target, i)
			return T(rune(i))
		}
	}
	fmt.Printf("Element %v not found in the list\n", target)
	return T(rune(0))
}


// Example:

uintArr := []uint{1, 3, 5, 7, 8, 2, 4, 6, 9}                                 // array of uint
strArr := []string{"A", "B", "C", "D", "F", "G", "H"}                                            // array of string
floatArr := []float64{9.0, 7.0, 5.0, 11.0, 12.0, 2.0, 14.0, 3.0, 10.0, 6.0} // array of float
intArr := []int{-9, 7, 5, -11, 12, -2, 14, 3, -10, 6}                       // array of int

// Element You wanna find inside the any  array
element := "A" 
search.Linear(arr, element)


// Recursive

func RecursiveLinear[T Numeric](arr []T, target T) T {
	var search func(int) int
	search = func(index int) int {
		if index >= len(arr) {
			fmt.Printf("%v not found in the list\n", target)
			return -1
		}

		if arr[index] == target {
			fmt.Printf("Found %v at index %v\n", target, index)
			return index
		}

		return search(index + 1)
	}

	return T(rune(search(0)))
}
```

### Multi Threading Linear Search - Go-Routines

```go
func LinearSearch[T Numeric](arr []T, target T) T {

    // runtime.GOMAXPROCS(runtime.NumCPU())
	numThreads := 4 // Number of go-routines (adjust as needed)

	arrLen := len(arr)
	wg := sync.WaitGroup{}
	resultChan := make(chan int)

	// Divide the array into segments for each goroutine
	segmentSize := arrLen / numThreads
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		startIdx := i * segmentSize
		endIdx := (i + 1) * segmentSize
		if i == numThreads-1 {
			endIdx = arrLen // Handle the last segment
		}
		go func(arr []T, target T, start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				if arr[i] == target {
					resultChan <- i
					return
				}
			}
		}(arr, target, startIdx, endIdx)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect the results from go-routines
	foundIndex := -1
	for idx := range resultChan {
		if idx != -1 {
			foundIndex = idx
			break
		}
	}

	if foundIndex != -1 {
		fmt.Printf("Element %v found at index %v\n", target, foundIndex)
	} else {
		fmt.Printf("Element %v not found in the list.\n", target)
	}
	return T(rune(0))
}
```

### Space & Time Complexity

| Complexity Type   | Time Complexity    | Space Complexity   |
|-------------------|--------------------|--------------------|
| Worst-case        | O(n)               | O(1)               |
| Best-case         | O(1)               | O(1)               |
| Average-case      | O(n)               | O(1)               |

- **Worst-case time complexity:** In the worst-case scenario, the linear search has to go through the entire list or array of `n` elements to find the target element, making it O(n).

- **Best-case time complexity:** In the best-case scenario, the target element is found at the beginning of the list, and only one comparison is needed, making it O(1).

- **Average-case time complexity:** On average, when the target element is not necessarily at the beginning or end of the list, the algorithm may have to check about half of the elements in the list, so it's still O(n).

- **Space complexity:** The linear search algorithm only requires a constant amount of space for storing the variables used in the algorithm, making it O(1) in terms of space complexity.

Certainly, here are the advantages, drawbacks, and when to use Linear Search in markdown format:

### Advantages of Linear Search

- **Simplicity:** Linear search is a straightforward and easy-to-understand algorithm, making it accessible to programmers of all skill levels.
- **No requirement for a sorted list:** Linear search can be used on unsorted or randomly ordered lists, unlike some other search algorithms that require the list to be sorted.
- **Minimal memory usage:** It has a constant space complexity, meaning it doesn't require additional memory to perform the search.

### Drawbacks of Linear Search

- **Inefficient for large datasets:** Linear search has a time complexity of O(n), which makes it inefficient for large datasets. As the dataset grows, the time required for the search grows linearly.
- **Not suitable for real-time systems:** Linear search's linear time complexity can be problematic in real-time applications where response time is critical.
- **There are faster alternatives:** For large datasets and when frequent searches are needed, more efficient algorithms like binary search or hash tables provide faster search times.

### When to Use Linear Search

You should consider using linear search when:

- You are working with a small dataset where the efficiency of the search algorithm is not a significant concern.
- The list is unsorted or partially sorted, and you want to find an element without the need for sorting it first.
- You are implementing a simple search algorithm for educational purposes or for a quick, one-time search.
- The dataset is relatively small, and you only need to search it infrequently, so you prioritize code simplicity over optimization.

Keep in mind that for larger datasets or scenarios where search efficiency is critical, other search algorithms with better time complexity, like binary search or hash tables, may be more suitable.

## Sentinel Linear Search

Sentinel Linear Search is a variation of the standard Linear Search algorithm that optimizes the search process by eliminating one comparison inside the loop. This is done by placing a "sentinel" value (the element you're searching for) at the end of the array. The sentinel value guarantees that the loop will always find the element it's searching for, thus removing the need for a separate comparison inside the loop.

**Scenario:** Searching for a Book in a Library

Imagine you're a librarian in a large library. People come in and ask you to find a specific book in the library's collection. You could use a Linear Search approach, where you start at the first shelf, examine each book one by one until you find the book, or determine that it's not there.

However, to speed up the process, you decide to use a Sentinel Linear Search approach. Instead of searching through each book and making a separate comparison for each book, you create a card catalog, where you have a list of books sorted by their titles. You place a "sentinel" card at the end of the catalog, representing the book you're searching for.

Now, when someone asks for a book, you go to the catalog, flip through the cards from start to end until you find the book you're looking for (or the sentinel card). Because you're guaranteed to find the sentinel card, you don't need to make a separate comparison to check if you've found the book you're looking for.

In this analogy, the catalog represents the array, the sentinel card represents the sentinel value, and your search through the catalog is analogous to the Sentinel Linear Search algorithm. It demonstrates how this optimization can save time and effort in the search process, especially when searching through a large collection.

### Code

```go
func Sentinel[T Numeric](arr []T, key T) T {
	var n int = len(arr)
	var last T = arr[n-1]
	arr[n-1] = key
	var idx int = 0

	for arr[idx] != key {
		idx++
	}

	arr[n-1] = last

	if (idx < n-1) || (arr[n-1] == key) {
		fmt.Printf("Found: %v present at Index: %v\n", key, idx)
	} else {
		fmt.Println("Element not found")
	}
	return T(rune(0))
}

// Recursive
func RecursiveSentinel[T Numeric](arr []T, key T) T {
	var n int = len(arr)
	if n <= 0 {
		return T(rune(0))
	}

	if arr[n-1] == key {
		fmt.Printf("Found: %v present at Index:%v\n", key, n-1)
		return T(rune(n - 1))
	} else {
		fmt.Printf("Not Found: %v\n", key)

	}
	return RecursiveSentinel(arr[n:], key)
}
```

### Multi Threading Sentinel Linear Search

```go
func SentinelLinearSearch[T Numeric](arr []T, key T) T {
	n := len(arr)
	numThreads := 4 // You can adjust the number of threads based on your requirements.
	step := n / numThreads
	resultChan := make(chan T)
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go SentinelSearch(arr, key, i*step, (i+1)*step, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		return result
	}
	fmt.Printf("Not found %v\n", key)
	return T(rune(0))

}

func SentinelSearch[T Numeric](arr []T, key T, start int, end int, wg *sync.WaitGroup, resultChan chan T) {
	defer wg.Done()
	for i := start; i < end; i++ {
		if arr[i] == key {
			resultChan <- T(rune(i))
			fmt.Printf("Found: %v present at Index:%v\n", key, i)
			return
		}
	}
}
```

### Space & Time Complexity

| Complexity          | Worst Case    | Average Case | Best Case    |
|---------------------|--------------|--------------|--------------|
| Time Complexity     | O(n)         | O(n)         | O(1)         |
| Space Complexity    | O(1)         | O(1)         | O(1)         |

- **Time Complexity**:
  - **Worst Case (O(n))**: In the worst case, Sentinel Linear Search will examine all elements in the array before finding the target element or determining that it's not present. This scenario occurs when the target element is the last element or not present in the array.
  - **Average Case (O(n))**: On average, you may need to examine approximately half of the elements before finding the target element or determining it's not present. The sentinel value does reduce the number of comparisons inside the loop but doesn't change the linear growth with the input size.
  - **Best Case (O(1))**: The best-case scenario occurs when the target element is found at the very beginning of the array. In this case, only one comparison is needed, and the algorithm completes very quickly.

- **Space Complexity**:
  - **Worst Case (O(1))**: The algorithm uses a constant amount of additional space regardless of the input size. The only additional space used is for the sentinel element, which is independent of the array size.
  - **Average Case (O(1))**: Like the worst case, the space complexity is constant and doesn't depend on the input size.
  - **Best Case (O(1))**: The space complexity remains constant even in the best-case scenario. The sentinel element is placed at the end of the array, but it doesn't depend on the array size.

These complexities are based on the assumption that the data structure is an array and that the algorithm is implemented as described in the previous answers.

### Advantages of Sentinel Linear Search

1. **Reduced Comparisons:** Sentinel Linear Search guarantees that you'll find the element you're searching for. This means that you eliminate one comparison inside the loop, which can be advantageous in terms of efficiency, especially when comparisons are costly operations.

2. **Simplicity:** It's a straightforward and easy-to-understand algorithm. You don't need to keep track of whether you've reached the end of the array; the sentinel value takes care of that.

3. **Applicability:** Sentinel Linear Search is applicable to both arrays and lists, making it versatile for different data structures.

### Drawbacks of Sentinel Linear Search

1. **Modification of Data:** The algorithm requires modifying the data structure (e.g., adding a sentinel element). If the data structure cannot be modified, this approach is not suitable.

2. **Initial Setup:** There's an initial setup cost to place the sentinel value at the end of the array, which can be unnecessary in some scenarios.

3. **Memory Overhead:** Adding a sentinel element consumes additional memory. In situations where memory usage is a concern, this can be a drawback.

### When to Use Sentinel Linear Search

1. **Small to Medium-sized Data:** Sentinel Linear Search is suitable for relatively small to medium-sized data sets. For very large data sets, more efficient search algorithms like binary search or hash-based methods may be preferable.

2. **Unsorted Data:** When the data is unsorted, linear search is typically the most straightforward search method. The sentinel variant can improve its performance without the need for additional sorting.

3. **Minimizing Comparisons:** If comparisons are computationally expensive and you want to minimize the number of comparisons, the sentinel approach can be beneficial.

4. **Data Modification Is Allowed:** If you can modify the data structure by adding a sentinel element, Sentinel Linear Search can be a viable option. This is often the case when working with in-memory data structures.

Sentinel Linear Search can be a good choice for smaller data sets, unsorted data, and situations where reducing the number of comparisons is important. However, it's essential to consider the trade-offs, including the initial setup and additional memory overhead, when deciding whether to use it. For larger data sets or sorted data, other search algorithms may offer better performance.

## Binary Search

Binary search is an efficient algorithm for finding a specific target element within a sorted collection (e.g., an array or list). It works by repeatedly dividing the search range in half, effectively cutting down the search space by half with each iteration, until the target element is found or determined to not exist in the collection.

1. **Precondition**: The collection must be sorted in ascending order for binary search to work correctly.

2. **Initialization**:
   - Start with the entire collection.
   - Define two pointers, usually called "left" and "right," which represent the current search range. Initially, set the "left" pointer to the first element (index 0) and the "right" pointer to the last element (index n-1), where n is the size of the collection.

3. **Search Iteration**:
   - Calculate the middle index between the "left" and "right" pointers. You can do this by taking the average of the two pointers: `mid = (left + right) / 2`.
   - Compare the element at the middle index with the target element you are looking for.
     - If the element at the middle index is equal to the target, you've found the target, and the search is complete.
     - If the element at the middle index is greater than the target, update the "right" pointer to `mid - 1`. This restricts the search to the lower half of the current range.
     - If the element at the middle index is less than the target, update the "left" pointer to `mid + 1`. This restricts the search to the upper half of the current range.

4. **Repeat**: Continue the search iteration until the "left" pointer is greater than the "right" pointer, which means the search range is empty, and the target element is not in the collection. Alternatively, if the target element is found, the search will terminate successfully.

Binary search has a time complexity of O(log n) because it effectively halves the search space with each iteration. This makes it significantly more efficient than linear search (which has a time complexity of O(n)) for large collections. However, binary search is only applicable when the collection is already sorted, and it may not be the best choice for small collections or dynamic data structures like linked lists.

### Code

```go
func Binary[T Numeric](arr []T, key T) int {

	var left, right int = 0, len(arr) - 1

	for left <= right {

		// Check bounds
		if left >= len(arr) || right < 0 {
			return -1
		}

		var mid int = left + (right-left)/2

		// Handle mid out of bounds
		if mid >= len(arr) {
			mid = len(arr) - 1
		}

		if arr[mid] == key {
			fmt.Printf("Found key %v at index %v\n", key, mid)
			return mid
		}

		if arr[mid] < key {
			left = mid + 1
			// Check left in bounds
			if left < len(arr) && arr[left] == key {
				fmt.Printf("Found key %v at index %v\n", key, left)
				return left
			}
		} else {
			right = mid - 1
			// Check right in bounds
			if right >= 0 && arr[right] == key {
				fmt.Printf("Found key %v at index %v\n", key, right)
				return right
			}
		}
	}
	fmt.Printf("Key %v not found in the array\n", key)
	return -1
}
```

### Recursive

```go
func RecursiveBinarySearch[T Numeric](arr []T, key T) int {
	return binarySearch(arr, key, 0, len(arr))
}

func binarySearch[T Numeric](arr []T, key T, low, high int) int {
	if low >= high {
		return -1 // Element not found
	}

	var mid int = low + (high-low)/2

	if arr[mid] == key {
		return mid // Element found
	} else if arr[mid] > key {
		return binarySearch(arr, key, low, mid)
	} else {
		return binarySearch(arr, key, mid+1, high)
	}
}
```

### MultiThread Binary Search

```go
func BinarySearch[T Numeric](arr []T, target T) int {
	var numThreads int = 4 // setting default number of thread
	var wg sync.WaitGroup
	resultChan := make(chan int)

	subArraySize := len(arr) / numThreads

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		start := i * subArraySize
		end := start + subArraySize

		if i == numThreads-1 {
			end = len(arr)
		}
		// Slandered MultiThreaded Binary Search
		go func(arr []T, target T, left, right int, wg *sync.WaitGroup, result chan int) {
			defer wg.Done()

			for left <= right {
				mid := left + (right-left)/2

				if arr[mid] == target {
					result <- mid
					return
				} else if arr[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}

			result <- -1 // Element not found
		}(arr, target, start, end-1, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for index := range resultChan {
		if index != -1 {
			return index
		}
	}

	return -1
}
```

### Advantages of Binary Search

1. **Efficiency:** Binary search is an efficient search algorithm. It has a time complexity of O(log n), which means it can quickly search large datasets, especially when compared to linear search, which has a time complexity of O(n). This makes it highly suitable for large and sorted datasets.

2. **Deterministic:** Binary search is deterministic. It will consistently find the target element if it exists in the sorted dataset. This predictability is essential in critical applications where correctness is paramount.

3. **Low Memory Usage:** Binary search typically uses minimal memory as it doesn't require additional data structures like hash tables. It can be implemented in a memory-efficient manner.

### Drawbacks of Binary Search

1. **Requirement for Sorted Data:** Binary search requires the data to be sorted. If the dataset is not sorted, it needs to be sorted first, which can be an additional time-consuming operation.

2. **No Duplicate Handling:** The basic binary search algorithm doesn't handle duplicate elements well. It may return any occurrence of the target value, which is not necessarily the first or last occurrence. Special modifications are needed for handling duplicates.

3. **Not Suitable for Dynamic Data:** Binary search is not suitable for dynamic data structures that frequently change. If the data changes often, maintaining the sorted order can be costly.

### When to Use Binary Search

1. **Large Sorted Datasets:** Binary search shines when you have large sorted datasets, such as arrays, lists, or databases. It's especially useful when the dataset is too large for a linear search to be practical.

2. **Databases:** Binary search is commonly used in database management systems for quickly retrieving data from sorted indices.

3. **Embedded Systems:** In resource-constrained environments like embedded systems, where memory usage is critical, binary search is a good choice due to its low memory footprint.

4. **Scientific and Mathematical Applications:** Binary search is widely used in scientific and mathematical applications where searching for specific values, e.g., in numerical methods, is common.

5. **Cryptography:** Cryptographic algorithms often involve searching for specific keys or values in large sets, where binary search can offer performance benefits.

Binary search is a powerful algorithm when dealing with sorted data and is best suited for scenarios where efficiency and deterministic behavior are essential. However, its requirement for sorted data and limitations in handling duplicates should be considered when choosing it for a particular application.

## Meta Binary Search | One-Sided Binary Search

Meta Binary Search, also known as One-Sided Binary Search, is a variation of the binary search algorithm that is used to search an ordered list or array of elements. This algorithm is designed to reduce the number of comparisons needed to search the list for a given element.

The basic idea behind the Meta Binary Search algorithm is to efficiently search for an element in a sorted, but rotated array, which is also known as a circularly sorted array. The key insight of Meta Binary Search is that you can still perform a binary search on a rotated array without the need for un-rotation, making it more efficient than some other approaches.

Here are the key steps and concepts behind Meta Binary Search:

1. **Initialization:** Start with the entire array. Find the number of bits (`lg`) required to represent the largest possible index in the array. This value is determined using the logarithm of the array's size minus 1 (to account for zero-based indexing).

2. **Binary Search Iteration:** Begin an iterative process that simulates a binary search within the rotated array. This process involves selecting positions based on the bit representation of the current search index.

3. **Position Update:** At each step, update the position to a new index, effectively shifting the binary search window. This update is done by setting a bit (starting with the most significant bit) in the current position.

4. **Comparison:** Compare the element at the updated position with the search key. If they match, return the current position as the index of the key.

5. **Direction of Movement:** The direction in which you move (left or right) depends on the comparison between the element at the current position and the search key. If the element is less than or equal to the key, you move to the right; otherwise, you move to the left.

6. **Repeat or Terminate:** Repeat the above steps until you find the key or exhaust the search space. If the key is found, return its index; otherwise, return -1 to indicate that the key is not present in the array.

The key advantage of Meta Binary Search is that it maintains the benefits of binary search and doesn't require un-rotating the array. This makes it efficient for searching in circularly sorted arrays. However, it assumes that the array is sorted and may have duplicate elements.

### Code

```go
// Calculation of log2 of the largest value
// I don't wanna use math.Log2 because it take float64 as input and then we 
// need to convert to int - so i implemented my own log2 that work on every type
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Calculation of log2
func log2[I Integer](x I) I {
	if x <= 0 {
		return 0 // Invalid input; you can choose any suitable error value
	}
	var result I = 0
	for x > 1 {
		x >>= 1 // Right shift x by 1, which is equivalent to dividing by 2 =  x = x/T(2)
		result++
	}

	return result
}


// Meta Binary Search
// It search for a key in a sorted, but rotated array with duplicates.
// It takes advantage of bitwise operations and bit manipulation to efficiently locate the key in the array.

func MetaBinary[T Numeric](arr []T, key T) int {
	var n int = len(arr)       // Get the length of the array
	var lg int = log2(n-1) + 1 // Calculate the number of bits needed to represent the largest array index

	var pos int = 0            // Initialize the position variable to start searching from the first element
	for i := lg; i >= 0; i-- { // Iterate from the most significant bit to the least significant bit
		if arr[pos] == key { // If the current position contains the key, return the index
			return pos
		}

		var newPos int = pos | (1 << i) // Construct a new position by setting the i-th bit

		if newPos < n && arr[newPos] <= key { // If the new position is within bounds and its value is less than or equal to the key
			pos = newPos // Update the current position to the new position
		}
	}

	if arr[pos] == key { // Check if the key is found at the final position
		return pos
	}

	return -1 // If the key is not found, return -1
}
```

### Advantages of Meta Binary Search

1. **Efficiency**: Meta Binary Search is highly efficient for searching in a sorted, but rotated array. It offers a time complexity of O(log n), which is similar to traditional binary search.

2. **No Need for Array Un-rotation**: Unlike other search methods for rotated arrays that require unrotation, Meta Binary Search operates directly on the rotated array, saving time and space complexity. Unrotation can be an expensive operation.

3. **Minimal Space Complexity**: Meta Binary Search uses a minimal amount of extra memory for variables. It doesn't require additional data structures or copying the array.

### Drawbacks of Meta Binary Search

1. **Array Must Be Sorted**: The array must be sorted for Meta Binary Search to work correctly. If the array is not sorted, the algorithm may not produce accurate results.

2. **Limited Applicability**: Meta Binary Search is specifically designed for circularly sorted arrays. It may not be suitable for other types of data or search scenarios.

### When to Use Meta Binary Search

Meta Binary Search is a specialized algorithm designed for a specific use case. You should consider using it when:

1. **You Have a Circularly Sorted Array**: If you have a sorted array that has been rotated (circularly sorted), and you need to search for an element in it, Meta Binary Search is a good choice.

2. **Efficiency Matters**: When searching for an element in a rotated array, especially a large one, the efficiency of the search algorithm becomes important. Meta Binary Search provides an efficient solution with a time complexity of O(log n).

3. **You Want to Avoid Array Un-rotation**: If you want to avoid the overhead of un-rotating the array, which can be computationally expensive, Meta Binary Search allows you to search the rotated array directly.

## Ubiquitous Binary Search

The basic Idea behind the Ubiquitous Binary Search Where we reduce the number of comparison in each iteration from two to one

### Code

```go
func UbiquitousBinary[T Numeric](arr []T, key T) int {

	var low int = 0
	var high int = len(arr) - 1

	for low <= high {
		var mid int = high - (low+high)/2

		if arr[mid] < key {
			low = mid + 1
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			return mid // value found
		}
	}

	return -1 // value not found
}
```

## Uniform Binary Search

Uniform Binary Search (UBS) is a variant of the binary search algorithm that pre-computes the indices of the middle elements of the search space. This allows UBS to avoid having to perform arithmetic operations to calculate the middle element in each iteration, which can improve performance on some platforms.

The basic idea of UBS is to create a lookup table of the indices of the middle elements of the search space. This lookup table is created once, before the search begins. Once the lookup table is created, the search algorithm can simply look up the index of the middle element in the table and then perform the comparison.

UBS can be a useful optimization for binary search on platforms where arithmetic operations are relatively slow. However, it is important to note that UBS requires additional memory to store the lookup table. Therefore, UBS may not be suitable for all applications.

### Code

```go
// Create a lookup table of the indices of the middle elements of the search space.
func lookupTable[T Numeric](array []T) []T {
	var lookupTable []T = make([]T, len(array))

	// Calculate the middle element index for each element in the array.
	for i := range array {
		// mid := int(math.Floor(float64(i+1) / 2.0))
		var mid T = T((i + 1) / 2.0)

		lookupTable[i] = mid
	}

	return lookupTable
}

// Perform a uniform binary search on the given array.
func uniformBinarySearch[T Numeric](array []T, target T, lookupTable []T) int {
	var low int = 0
	var high int = len(array) - 1

	// While the low pointer is less than or equal to the high pointer, continue the search.
	for low <= high {
		// Get the index of the middle element from the lookup table.
		var mid int = int(lookupTable[low])

		// If the target element is equal to the middle element, return the index of the middle element.
		if array[mid] == target {
			return mid
		}

		// If the target element is less than the middle element, set the high pointer to the middle element minus one.
		if target < array[mid] {
			high = mid - 1
		} else {
			// Otherwise, set the low pointer to the middle element plus one.
			low = mid + 1
		}
	}

	// If the target element was not found, return -1.
	return -1
}

func Uniform[T Numeric](arr []T, key T) int {
	var lookupTable []T = lookupTable(arr)
	var index int = uniformBinarySearch(arr, key, lookupTable)
	if index != -1 {
		fmt.Printf("Found Key: %v at Index: %v", key, index)
	} else {
		fmt.Printf("Key: %v not Found", key)
	}
	return index
}
```

### Advantages

1. **Consistent Steps:** UBS ensures that the steps taken during the search are uniform, making it particularly useful when you want to bound the number of comparisons or ensure a predictable number of steps. This can be important in real-time or embedded systems.

2. **Simplicity:** UBS is straightforward to implement and can be easier to understand compared to regular binary search, which may require careful calculation of midpoints.

3. **Predictability:** In UBS, the pivot is determined by evenly dividing the search space, making it predictable and repeatable.

### Drawbacks

1. **Inefficient for Non-Uniform Distributions:** UBS can be less efficient than regular binary search for data that is not uniformly distributed. If the data is clustered or skewed, UBS may take more steps than necessary.

2. **Extra Space:** UBS typically requires an additional lookup table to store pivot information, which consumes memory. In some memory-constrained applications, this could be a drawback.

3. **Complexity for Dynamic Data:** UBS assumes that the data remains sorted and constant. If you have a dynamic dataset with frequent insertions or deletions, maintaining the lookup table can become complex and inefficient.

### When to Use Uniform Binary Search

The decision to use UBS depends on your specific use case and requirements:

1. **Predictable Performance:** If you need a search algorithm with predictable performance in terms of the number of comparisons or steps, UBS can be a good choice. This is particularly valuable in real-time systems or when you want to ensure a consistent response time.

2. **Uniformly Distributed Data:** UBS is most effective when the data you are searching through is uniformly distributed. If you can guarantee that your data adheres to this pattern, UBS can be a suitable choice.

3. **Memory Trade-Off:** If memory usage is not a significant concern and you prefer a simpler implementation, you might choose UBS over traditional binary search.

4. **Static or Infrequently Updated Data:** UBS is more suitable for datasets that don't change frequently. If your data is mostly static or updated infrequently, the overhead of maintaining the lookup table may be acceptable.

### Randomize Binary Search

Similar to Normal Binary Search but We will get the middle element Randomly
It is a Las Vegas randomized algorithm as it always finds the correct result.

### Code

```go
func getRandom(x int, y int) int {
	return x + (rand.Intn(y - x + 1))
}

func RandomBinary[T Numeric](arr []T, key T) int {
	var left, right int = 0, len(arr) - 1

	for left <= right {

		// Check bounds
		if left >= len(arr) || right < 0 {
			return -1
		}

		var mid int = getRandom(left, right)

		// Handle mid out of bounds
		if mid >= len(arr) {
			mid = len(arr) - 1
		}

		if arr[mid] == key {
			fmt.Printf("Found key %v at index %v\n", key, mid)
			return mid
		}

		if arr[mid] < key {
			left = mid + 1
			// Check left in bounds
			if left < len(arr) && arr[left] == key {
				fmt.Printf("Found key %v at index %v\n", key, left)
				return left
			}
		} else {
			right = mid - 1
			// Check right in bounds
			if right >= 0 && arr[right] == key {
				fmt.Printf("Found key %v at index %v\n", key, right)
				return right
			}
		}
	}
	fmt.Printf("Key %v not found in the array\n", key)
	return -1
}
```

## N-Base Modified Binary Search

 This algorithm is used to search for a target element in a sorted array, but it does so in a base-N fashion, where N is a predefined constant. The key idea is to optimize the search by skipping over parts of the array that are unlikely to contain the target element.

1. The code starts with an initialization section where it sets the base (N) and a constant (PowerOf2). The base, N, determines how the search is divided, and PowerOf2 is used as a constant.

2. The function responsible for searching the target element in the array. It takes the following parameters:
   - `arr`: The sorted integer array in which you want to find the target.
   - `base`: The size of the array.
   - `target`: The value you're searching for.

3. **Power of N Division**: The code divides the search into regions based on powers of N. It starts by finding the first power of N (greater than 1) that is greater than or equal to the array size (`M`). This ensures that the entire array is divided into regions.

4. **Main Loop**:
   - It uses two pointers, `i` and `step1`, to navigate the array.
   - `i` represents the current position in the array.
   - `step1` is initialized to the first power of N and is used to step through the array in base-N chunks.
   - The loop continues as long as `step1` is greater than 0.

5. **Optimizing the Search**:
   - At each step, it checks if the element at the index `i + step1` is less than or equal to the target. If it is, this means the target could potentially be in this region, and it proceeds to investigate further.
   - It uses another loop with `step2` to refine the search. It tries to find the maximum number of times it can step in the given region.
   - The code keeps updating the `i` based on how many times the power of N can be used in the current region.

6. **Target Found or Not Found**: After the loop completes, it checks whether the element at index `i` is equal to the target. If it is, the target is found, and the function returns the index. If not, it returns -1 to indicate that the target is not present in the array.

The Idea here is to reduce the number of comparisons required to find the target element by skipping over regions of the array that are guaranteed to be too small or too large to contain the target. This is achieved by dividing the array into regions based on powers of N and iteratively narrowing down the search space.

### Code

```go
func NumericBaseBinarySearch[T Numeric](arr []T, key T, base int) int {
	var N = 3
	var PowerOf2 = 4
	var i, step1, step2, times int

	// Find the first power of N greater than the array size
	for step1 = 1; step1 < base; step1 *= N {
	}

	for i = 0; step1 > 0; step1 /= N {
		// Each time a power can be used, count how many times it can be used
		if i+step1 < base && arr[i+step1] <= key {
			for times = 1; step2 > 0; step2 >>= 1 {
				if i+(step1*(times+step2)) < base && arr[i+(step1*(times+step2))] <= key {
					times += step2
				}
			}
			step2 = PowerOf2

			// Add to the final result how many times the power can be used
			i += times * step1
		}
	}

	// Return the index if the element is present in the array, else return -1
	if arr[i] == key {
		fmt.Printf("Key: %v Found at Index: %v", key, i)
		return i
	}
	fmt.Printf("Key: %v Not Found", key)
	return -1
}
```

## Ternary Search

Ternary search is a search algorithm that divides the search space into three parts instead of two in each iteration. This can lead to better performance for certain types of search problems.

Ternary search works by first calculating the middle three elements of the search space. The algorithm then compares the target element to these middle three elements and eliminates the middle part of the search space that does not contain the target element. The algorithm then repeats this process on the remaining two parts of the search space until the target element is found or the search space is empty.

Ternary search can be more efficient than binary search for certain types of search problems, such as searching for a value in a sorted list where the values are evenly spaced. This is because ternary search can eliminate more of the search space in each iteration.

1. First, we compare the key with the element at mid1. If found equal, we return mid1.
2. If not, then we compare the key with the element at mid2. If found equal, we return mid2.
3. If not, then we check whether the key is less than the element at mid1. If yes, then recur to the first part.
4. If not, then we check whether the key is greater than the element at mid2. If yes, then recur to the third part.
5. If not, then we recur to the second (middle) part.

### Code

```go
func Ternary[T Numeric](arr []T, key T) int {
	// Initialize low and high as the first and last indices of the array.
	var low int = 0
	var high int = len(arr) - 1

	// Perform a loop until low is less than or equal to high.
	for low <= high {
		// Calculate the first and second midpoints.
		var firstMid int = low + (high-low)/3
		var secMid int = high - (high-low)/3

		// If the key is found at the first midpoint, return its index.
		if arr[firstMid] == key {
			return firstMid
		}

		// If the key is found at the second midpoint, return its index.
		if arr[secMid] == key {
			return secMid
		}

		// If the key is smaller than the element at the first midpoint, update 'high.'
		if key < arr[firstMid] {
			high = firstMid - 1
		} else if key > arr[secMid] {
			// If the key is larger than the element at the second midpoint, update 'low.'
			low = secMid + 1
		} else {
			// If the key is between the elements at firstMid and secMid, adjust 'low' and 'high.'
			low = firstMid + 1
			high = secMid - 1
		}
	}

	// If the key is not found in the array, return -1.
	return -1
}

// Recursive

func ternarySearch[T Numeric](arr []T, key T, low int, high int) int {

	if low > high {
		fmt.Printf("Key: %v Not found\n", key)
		return -1
	}

	var mid1 int = low + (high-low)/3
	var mid2 int = high - (high-low)/3

	if arr[mid1] == key {
		fmt.Printf("Key: %vFound at index %v\n", key, mid1)
		return mid1
	}

	if arr[mid2] == key {
		fmt.Printf("Key: %v Found at index %v\n", key, mid2)
		return mid2
	}

	if key < arr[mid1] {
		return ternarySearch(arr, key, low, mid1-1)
	} else if key > arr[mid2] {
		return ternarySearch(arr, key, mid2+1, high)
	} else {
		return ternarySearch(arr, key, mid1+1, mid2-1)
	}

}

// Helper function
func RecursiveTernary(arr []int, key int) int {
	return ternarySearch(arr, key, 0, len(arr)-1)

}
```

### MultiThreaded Ternary Search

```go
func ParallelTernary[T Numeric](arr []T, key T) int {
	// Create a channel to communicate results between goroutines.
	var goroutines int = 4 // Default Adjust According to YOUR NEED
	resultChan := make(chan int, goroutines)
	var low int = 0
	var high int = len(arr) - 1

	// Function for a single search segment.
	searchSegment := func(start int, end int) {
		for i := start; i <= end; i++ {
			// Calculate the first and second midpoints for this segment.
			firstMid := i + (end-i)/3
			secMid := end - (end-i)/3

			// If the key is found at the first midpoint, send its index to the channel.
			if arr[firstMid] == key {
				resultChan <- firstMid
				return
			}

			// If the key is found at the second midpoint, send its index to the channel.
			if arr[secMid] == key {
				resultChan <- secMid
				return
			}

			// If the key is smaller than the element at the first midpoint, update 'end.'
			if key < arr[firstMid] {
				end = firstMid - 1
			} else if key > arr[secMid] {
				// If the key is larger than the element at the second midpoint, update 'start.'
				i = secMid + 1
			} else {
				// If the key is between the elements at firstMid and secMid, adjust 'start' and 'end.'
				i = firstMid + 1
				end = secMid - 1
			}
		}
		resultChan <- -1
	}

	// Divide the search space among goroutines.
	for i := 0; i < goroutines; i++ {
		segmentSize := (high - low + 1) / goroutines
		start := i * segmentSize
		end := start + segmentSize - 1

		// The last goroutine might cover extra elements if the division isn't exact.
		if i == goroutines-1 {
			end = high
		}

		// Start a goroutine to search its segment.
		go searchSegment(start, end)
	}

	// Wait for results from goroutines and close the result channel.
	for i := 0; i < goroutines; i++ {
		result := <-resultChan
		if result != -1 {
			return result
		}
	}
	close(resultChan)

	// If the key is not found in the array, return -1.
	return -1
}
```

### Advantages of Ternary Search

1. **Efficiency:** Ternary search is more efficient than linear search when working with sorted data. It narrows down the search space by dividing it into three parts with each iteration, reducing the number of comparisons needed to find the target element.

2. **Logarithmic Time Complexity:** Ternary search operates in O(log3 N) time complexity, which is significantly better than linear search (O(N)) and is comparable to binary search (O(log2 N).

3. **Fewer Comparisons:** Ternary search typically requires fewer comparisons than binary search, which can be advantageous in certain situations, especially when the search space is large.

### Drawbacks of Ternary Search

1. **Array Requirement:** Ternary search assumes that the data is stored in an array and that the array is sorted. If your data is not organized in this way, you'll need to pre-process it, which can be time-consuming and resource-intensive.

2. **Limited Usefulness:** Ternary search is most effective when searching for a single target value. If you need to find multiple occurrences or need range-based searches, other algorithms might be more appropriate.

3. **Complexity:** Ternary search can be more complex to implement and understand compared to binary search. The algorithm involves handling three partitions of the array, which can increase the chance of errors in the code.

### When to Use Ternary Search

1. **Sorted Data:** Ternary search is most suitable for sorted data. If your data is not sorted, you should sort it first, which could be a time-consuming operation. In cases where data is frequently updated, other search algorithms like binary search might be a better choice.

2. **Single Occurrence Search:** Ternary search is ideal for finding a single occurrence of a target value within the data. If you need to find all occurrences, a different search algorithm might be more appropriate.

3. **Balanced Search Space:** Ternary search is efficient when the search space is reasonably balanced (not heavily skewed to one side). If the data has a skewed distribution, other search algorithms like binary search might perform better.

4. **Logarithmic Complexity:** When dealing with large datasets, ternary search's logarithmic time complexity can be an advantage. It ensures that the search time doesn't grow as quickly as with linear search.

Ternary search is a valuable algorithm when dealing with sorted data and searching for a single target value within a reasonably balanced search space. However, it may not be the best choice for all scenarios, and the choice of search algorithm depends on the specific requirements of your application.

## Jump Search

Jump search is a search algorithm that works by jumping through the search space instead of searching it linearly. It is similar to binary search, but instead of dividing the search space in half in each iteration, it jumps a certain number of elements ahead.

Jump search is most efficient when the search space is sorted and the elements are evenly distributed. It is also efficient when the search space is very large.

1. Calculate the step size. The step size is typically the square root of the length of the array.
2. Initialize the current index to 0.
3. While the current index is less than the length of the array and the element at the current index is less than the target:
    - Jump ahead the step size.
4. If the current index is greater than or equal to the length of the array, then the target element is not in the array.
5. Linearly search the remaining elements of the array for the target element.
6. If the target element is found, return its index. Otherwise, return -1.

### Code

```go
func Jump[T Numeric](arr []T, key T) int {
	var n int = len(arr) // Get the length of the array.
	var step int = sqrt(n) // Calculate the jump step size based on the square root of n.

	var prev int = 0 // Initialize the 'prev' variable to 0.
	
	// The first loop finds the block where the element is or should be.
	for minStep := min(step, n) - 1; arr[minStep] < key; minStep = min(step, n) - 1 {
		prev = step
		step += sqrt(n)
		if prev >= n {
			return -1 // If 'prev' exceeds the length of the array, the key is not present.
		}
	}

	// The second loop performs a linear search within the block.
	for arr[prev] < key {
		prev++

		if prev == min(step, n) {
			return -1 // If 'prev' reaches the end of the block, the key is not found.
		}
	}

	if arr[prev] == key {
		return prev // If the key is found, return its index.
	}

	return -1 // If the key is not found in the array, return -1.
}

func min[T Numeric](x T, y T) T {
	if x < y {
		return x // Return the smaller of the two values.
	}
	return y
}

func sqrt[T Numeric](x T) T {
	// This function approximates the square root of x using the Newton-Raphson method.
	var z T = 1.0 // Initial guess for the square root.

	for i := 1; i <= 10; i++ { // Perform 10 iterations for a better approximation.
		z -= (z*z - x) / (2 * z) // Update the guess using the Newton-Raphson formula.
	}

	return z // Return the approximated square root.
}

```

### Multi Threaded Jump Search

```go
func jumpParallel[T Numeric](arr []T, key T, startIndex int, endIndex int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	var prev int = startIndex

	for ; prev < endIndex; prev++ {
		if arr[prev] == key {
			resultChan <- prev // Send the result to the channel if the key is found.
			return
		}
	}

	resultChan <- -1 // Send -1 to indicate that the key was not found in this range.
}

// JumpMultithreaded performs a multithreaded jump search on the array.
func ParallelJump[T Numeric](arr []T, key T) int {
	var n int = len(arr)
	numThreads := runtime.NumCPU() // Determine the number of CPU cores available.

	resultChan := make(chan int, numThreads) // Create a channel for collecting search results.
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		startIndex := (n * i) / numThreads     // Calculate the start index for the current partition.
		endIndex := (n * (i + 1)) / numThreads // Calculate the end index for the current partition.

		wg.Add(1)                                                        // Increment the WaitGroup counter for the new goroutine.
		go jumpParallel(arr, key, startIndex, endIndex, &wg, resultChan) // Launch a goroutine to search this partition.
	}

	wg.Wait()         // Wait for all goroutines to finish.
	close(resultChan) // Close the result channel to signal that no more results will be sent.

	// Collect results from the channel. The first non-negative result indicates the index of the key.
	for result := range resultChan {
		if result != -1 {
			return result
		}
	}

	return -1 // If the key is not found in any partition, return -1.
}
```

### Advantages

1. **Efficiency:** Jump search is efficient for large datasets, especially when compared to linear search. It requires fewer comparisons than linear search due to the "jumping" nature of the algorithm. This can result in faster searches for large sorted arrays.

2. **Deterministic:** Jump search is deterministic, meaning it will always find the element if it exists in the array. It does this by dividing the array into blocks and then performing a linear search within the selected block, guaranteeing that the element is found, if present.

3. **Simplicity:** The algorithm is relatively simple to implement, and it doesn't require additional data structures or complex operations.

### Drawbacks

1. **Requires Sorting:** Jump search requires that the input array is sorted. If the array is unsorted, you'll need to sort it first, which may be an additional computational overhead.

2. **Inefficient for Small Arrays:** For small arrays, the overhead of calculating the jump step and performing linear search within a block may make jump search less efficient than a simple linear search.

3. **Not Suitable for Dynamic Data:** Jump search is most useful when the array doesn't change frequently. If the array frequently changes (insertions or deletions), maintaining the sorted order may not be efficient.

## When to Use Jump Search

Jump search is useful in specific scenarios:

1. **Large Sorted Arrays:** Jump search is well-suited for searching within large, sorted arrays, where its efficiency in terms of the number of comparisons and the reduction in search time is evident.

2. **Static or Infrequently Updated Data:** If the data is relatively static or updated infrequently, and the primary operation is searching, jump search can be a good choice.

3. **Deterministic Search:** When you need a deterministic search algorithm that ensures the element's presence, and you don't want to rely on probabilistic algorithms like binary search or interpolation search.

4. **Memory Constraints:** If you're working with limited memory and can't afford to store additional data structures, jump search can be a viable option.

## Interpolation Search

Interpolation search is an algorithm used to find a specific element within a sorted array. It improves upon binary search by estimating the likely position of the target element based on its value relative to the values at the endpoints of the array. This estimation allows interpolation search to potentially narrow down the search range more quickly than binary search.

Here's the general approach to implement the interpolation search algorithm:

1. **Input Data:** The algorithm requires a sorted array and the target element you want to find within that array.

2. **Calculate the Estimated Position:** The key idea behind interpolation search is to estimate the position of the target element based on its value. It uses a linear interpolation formula to calculate this estimate. The formula is:

    ```go
    position = low + ((key - arr[low]) * (high - low)) / (arr[high] - arr[low])
    ```

    Where:
    - `position` is the estimated position of the key.
    - `low` is the starting index of the current search range.
    - `high` is the ending index of the current search range.
    - `key` is the value you're searching for.
    - `arr[low]` and `arr[high]` are the values at the respective endpoints of the search range.

3. **Compare and Narrow the Search Range:** Once you have the estimated position, compare the value at `arr[position]` with the target `key`. Based on this comparison:
    - If `arr[position]` is equal to `key`, you've found the target element, and you return the `position`.
    - If `arr[position]` is less than `key`, it means the target is likely to be in the right sub-array, so update `low` to `position + 1`.
    - If `arr[position]` is greater than `key`, it means the target is likely to be in the left sub-array, so update `high` to `position - 1`.

4. **Repeat:** Repeat steps 2 and 3 until one of the following conditions is met:
    - The target element is found, and you return its index.
    - The `low` index exceeds the `high` index, indicating that the target element is not present in the array. In this case, you return -1 to indicate that the element was not found.

5. **Final Result:** Return the index of the target element if it was found, or -1 if it's not in the array.

Interpolation search is effective for finding elements in sorted arrays, especially when the values in the array are uniformly distributed. However, it may not be as efficient for non-uniform data distributions, where the data is clustered in a specific range.

```go
func Interpolation[T Numeric](arr []T, key T) int {
	var low int = 0
	var high int = len(arr) - 1

	for low <= high && key >= arr[low] && key <= arr[high] {
		if low == high {
			// If the search range narrows down to a single element, check if it's the key.
			if arr[low] == key {
				return low
			}
			// If the single element is not the key, return -1 (key not found).
			return -1
		}

		// Calculate the position estimate using interpolation formula.
		var position int = low + ((int(key)-int(arr[low]))*(high-low))/(int(arr[high])-int(arr[low]))

		// Perform a bound check to ensure 'position' is within the array bounds.
		if position < low || position > high {
			return -1
		}

		if arr[position] == key {
			// If the key is found at the estimated position, return that position.
			return position
		}

		if arr[position] < key {
			// If the estimated value is less than the key, narrow the search range to the right.
			low = position + 1
		} else {
			// If the estimated value is greater than the key, narrow the search range to the left.
			high = position - 1
		}
	}

	// If the key is not found within the search range, return -1 (key not found).
	return -1
}

// Recursive

func RecursiveInterpolation[T Numeric](arr []T, key T, high int, low int) int {
	if low <= high && key >= arr[low] && key <= arr[high] {
		if low == high {
			if arr[low] == key {
				return low
			}
			return -1
		}

		var position int = low + ((int(key)-int(arr[low]))*(high-low))/(int(arr[high])-int(arr[low]))

		if arr[position] == key {
			return position
		}

		if arr[position] < key {
			return RecursiveInterpolation(arr, key, position+1, high)
		}

		return RecursiveInterpolation(arr, key, low, position-1)
	}

	return -1
}
```

### MultiThreaded Interpolation search

```go
func ParallelInterpolation[T Numeric](arr []T, key T) int {
	// Determine the number of CPU cores available.
	numThreads := runtime.NumCPU()

	// Initialize variables for the result and a mutex for synchronization.
	var result int
	var resultMutex sync.Mutex
	var wg sync.WaitGroup

	// Function to perform interpolation search within a partition of the array.
	searchPartition := func(partitionIndex int) {
		defer wg.Done()

		// Calculate the search range for the current partition based on the number of threads.
		low := (len(arr) * partitionIndex) / numThreads
		high := (len(arr)*(partitionIndex+1))/numThreads - 1

		localResult := -1 // Initialize a local result for this partition.

		// Interpolation search within the partition.
		for low <= high && key >= arr[low] && key <= arr[high] {
			if low == high {
				if arr[low] == key {
					localResult = low // Set the local result if the key is found in this partition.
				}
				break
			}

			// Calculate the estimated position using interpolation formula.
			var position int = low + ((int(key)-int(arr[low]))*(high-low))/(int(arr[high])-int(arr[low]))

			if arr[position] == key {
				localResult = position // Set the local result if the key is found at the estimated position.
				break
			}

			if arr[position] < key {
				low = position + 1 // Narrow the search range to the right.
			} else {
				high = position - 1 // Narrow the search range to the left.
			}
		}

		// Update the overall result using a mutex to ensure that only the first found result is considered.
		resultMutex.Lock()
		if localResult != -1 {
			result = localResult
		}
		resultMutex.Unlock()
	}

	wg.Add(numThreads) // Increment the wait group counter to track goroutines.

	// Launch multiple goroutines to search different partitions in parallel.
	for i := 0; i < numThreads; i++ {
		go searchPartition(i)
	}

	wg.Wait() // Wait for all goroutines to finish.

	return result // Return the final result.
}
```

### Advantages

1. **Efficiency for Uniform Distributions:** Interpolation search is highly efficient for uniformly distributed data. It can significantly reduce the number of comparisons required to find the target element in such cases.

2. **Deterministic and Reliable:** Interpolation search is a deterministic algorithm. It will always find the element if it exists in the array, as long as the data is evenly distributed. This reliability can be important in certain applications.

3. **Works Well with Large Data:** Interpolation search can be more efficient than other search algorithms, such as linear search, for large datasets, thanks to its ability to quickly narrow down the search range based on the estimated position.

4. **Adaptive Search:** Interpolation search adapts to the distribution of data and the position of the target element, which makes it suitable for scenarios where the data distribution varies.

### Drawbacks

1. **Inefficiency for Non-Uniform Distributions:** When the data is not evenly distributed and is clustered in specific ranges, interpolation search can perform poorly. In such cases, it may require more comparisons than other search algorithms, such as binary search.

2. **Complexity:** The interpolation formula can involve integer division and multiplication, which may introduce some computational overhead. In some cases, simple algorithms like binary search can outperform it.

3. **Precondition of Sorted Data:** Interpolation search requires the data to be sorted. If the data is not initially sorted, sorting it before applying the search algorithm can add significant time and space complexity.

### When to Use Interpolation Search

Interpolation search is most useful in specific scenarios:

1. **Uniform Data Distribution:** It excels when the data is uniformly distributed, where most values are evenly spread throughout the array.

2. **Large Sorted Arrays:** It is well-suited for searching within large, sorted arrays, where its efficiency in terms of the number of comparisons and reduction in search time is evident.

3. **Deterministic Search:** When you need a deterministic search algorithm that ensures the element's presence and don't want to rely on probabilistic algorithms.

4. **Dynamic Data with Varied Distributions:** It can be useful in scenarios where data distribution varies, as it can adapt to the distribution and potentially reduce the number of comparisons.

## Exponential Search

Exponential search, also known as "exponential binary search," is a search algorithm that aims to find the position of a specific element within a sorted array. It combines elements of linear and binary search techniques. The idea behind exponential search is to start with a small step size and double it with each iteration until a range is found that may contain the target element, and then perform a binary search within that range.

1. Start with an initial step size of 1 (or any small number) and initialize the index to 0.

2. Double the step size in each iteration, moving to indexes 1, 2, 4, 8, 16, and so on, until you find a value greater than or equal to the target element. At this point, you have identified a search range.

3. Perform a binary search within the identified range. This binary search can help locate the exact position of the target element within the array.

4. If the target element is found during the binary search, return its index; otherwise, return -1 to indicate that the element is not in the array.

The key idea is to quickly identify a broad range where the target element might be located, and then perform a more efficient binary search within that narrowed-down range.

Exponential search is particularly useful when you have a sorted array and don't know the approximate location of the target element. It ensures that you find the element efficiently by first estimating a range and then narrowing it down, combining the advantages of both linear and binary search techniques.

### Code

```go
func Exponential[T Numeric](arr []T, key T) int {
	// Get the length of the array.
	var n int = len(arr)

	// If the array is empty, return -1 (key not found).
	if n == 0 {
		return -1
	}

	// Initialize the 'index' variable to 1 to start exponential search.
	var index int = 1

	// Double the 'iter' value until it's within the array bounds and the element at 'iter' is less than 'key'.
	for index < n && arr[index] < key {
		index *= 2
	}

	// Define 'left' as half of the 'iter' and 'right' as the minimum of 'iter' and the last index in the array.
	var left int = index / 2
	var right int = min(index, n-1)

	// Perform binary search within the range defined by 'left' and 'right'.
	for left <= right {
		// Calculate the middle index.
		var mid int = (left + right) / 2

		// Check if the element at 'mid' is equal to 'key'.
		if arr[mid] == key {
			// If the element is found at 'mid', return its index.
			return mid
		} else if arr[mid] < key {
			// If the element at 'mid' is less than 'key', narrow the search range to the right.
			left = mid + 1
		} else {
			// If the element at 'mid' is greater than 'key', narrow the search range to the left.
			right = mid - 1
		}
	}

	// If the key is not found, return -1.
	return -1
}
```

### Advantages of Exponential Search

1. **Efficiency on Sorted Data:** Exponential search is particularly efficient on sorted data. It combines elements of both linear and binary search, making it ideal for scenarios where data is sorted and you don't have information about the range of values in advance.

2. **Reduced Number of Comparisons:** It reduces the number of comparisons when the target element is found, especially when the target element is closer to the beginning of the array.

3. **Simplicity:** It is relatively simple to implement and doesn't require recursion, making it an attractive option for certain situations.

### Drawbacks of Exponential Search

1. **Inefficient for Unsorted Data:** Exponential search is not suitable for unsorted data. Its performance greatly depends on data being sorted, and it doesn't provide any advantage in terms of efficiency if the data is not in ascending order.

2. **Overhead:** It incurs additional overhead due to its two-phase nature, which can be problematic if efficiency is a critical concern.

3. **Not Always the Fastest:** While it can be efficient in some cases, it's not always the fastest search algorithm. Binary search or interpolation search might be more suitable for specific situations.

### When to Use Exponential Search

1. **Large Sorted Data Sets:** Exponential search is most useful when you have a large, sorted dataset, and you want to find an element within it. In such cases, it can be a good choice because it narrows down the search range more quickly than linear search.

2. **Unknown Data Distribution:** When you have limited information about the distribution of the data, exponential search can be a good initial search to quickly identify the possible range for further searches.

3. **When Simplicity Is Preferred:** If you prefer a relatively simple search algorithm that works reasonably well on sorted data, exponential search is a good choice.

4. **Hybrid Algorithms:** Exponential search can be used in hybrid algorithms, where it serves as an initial search to narrow down the range and is then followed by a more efficient search algorithm, such as binary search, for the final refinement.

## Fibonacci Search

Fibonacci search is an efficient search algorithm used to find the position of a specific element in a sorted array. It's similar to binary search but uses a different approach to divide the search space. The key idea behind Fibonacci search is to use Fibonacci numbers to determine the positions to be checked.

### Idea Behind

1. Pre-Compute or Calculate Fibonacci Numbers

- Start by calculating a sequence of Fibonacci numbers. You can pre-compute them or calculate them on-the-fly in your algorithm.

2. Initialize Variables

- Initialize three variables: `fib1`, `fib2`, and `fib3`. Set them as follows:
  - `fib1` = 0
  - `fib2` = 1
  - `fib3` = fib1 + fib2

3. Determine the Search Space

- Find the smallest Fibonacci number that is greater than or equal to the length of the sorted array. This step ensures that you determine suitable positions for comparison within the array.
- Use a loop to update the Fibonacci numbers as follows:
  - While `fib3` is less than the length of the array:
    - Set `fib1` to the value of `fib2`.
    - Set `fib2` to the value of `fib3`.
    - Update `fib3` as the sum of `fib1` and `fib2`.

4. Initialize an Offset Variable

- Initialize an offset variable (e.g., `offset`) to -1. This variable will keep track of the previous position you compared during the search.

5. Search for the Target Element

- Use a loop to search for the target element `x` in the array. The loop should continue as long as `fib3` is greater than 1.
- Inside the loop:
  - Calculate the index `i` to be checked. You can use `min(offset + fib2, n - 1)` where `n` is the length of the array.
  - Compare the element at index `i` with `x`.
  - Depending on the comparison result, update the Fibonacci variables for a smaller or larger search segment.
  - If the element at index `i` is equal to `x`, you've found the target, and you can return its index.

6. Final Check for the Last Element

- After the loop, check if the last element in the array is equal to `x` to account for the possibility that `x` is located in the last position.
- If the last element is equal to `x`, return the index of the last element. Otherwise, return -1 to indicate that the element `x` is not in the array.

```go
func Fibonacci[T Numeric](arr []T, x T) int {
	// Calculate the length of the input array.
	var n int = len(arr)

	// If the array is empty, return -1.
	if n == 0 {
		return -1
	}

	// Initialize Fibonacci variables.
	var fib1 int = 0       // First Fibonacci number
	var fib2 int = 1       // Second Fibonacci number
	var fib3 = fib1 + fib2 // Third Fibonacci number

	// Find the smallest Fibonacci number that is greater than or equal to n.
	for fib3 < n {
		fib1 = fib2
		fib2 = fib3
		fib3 = fib1 + fib2
	}

	// Initialize an offset variable to -1.
	var offset int = -1

	// Search for the element x in the array using Fibonacci search.
	for fib3 > 1 {
		// Calculate the index to be checked.
		var i int = min(offset+fib2, n-1)

		if arr[i] < x {
			// Update Fibonacci variables for a smaller segment.
			fib3 = fib2
			fib2 = fib1
			fib1 = fib3 - fib2
			offset = i
		} else if arr[i] > x {
			// Update Fibonacci variables for a larger segment.
			fib3 = fib1
			fib2 = fib2 - fib1
			fib1 = fib3 - fib2
		} else {
			// Element x found at index i.
			return i
		}
	}

	// Check if the last element in the array is equal to x.
	if fib2 == 1 && arr[offset+1] == x {
		return offset + 1
	} else {
		// Element x was not found in the array.
		return -1
	}
}
```

### Advantages of Fibonacci Search

1. **Efficiency in Large Arrays:** Fibonacci search is particularly efficient for searching large arrays or lists. It reduces the number of comparisons required compared to linear search.

2. **Adaptive Search:** Fibonacci search dynamically adapts the search interval based on the Fibonacci sequence, making it effective for different-sized search spaces.

3. **Better than Linear Search:** It is significantly faster than linear search because it reduces the search space exponentially, while linear search has a linear search space reduction.

4. **Simple to Implement:** The algorithm is relatively simple to implement and understand. It doesn't require additional data structures, making it easy to code from scratch.

5. **Low Memory Overhead:** Fibonacci search uses a small number of variables and doesn't require additional memory overhead.

### Drawbacks of Fibonacci Search

1. **Preprocessing of Fibonacci Numbers:** To apply Fibonacci search, you need to either pre-compute Fibonacci numbers or calculate them on-the-fly. This preprocessing can add overhead if the array is extremely large.

2. **Limited Usage:** Fibonacci search is most effective when the array is large and the elements are uniformly distributed. It may not be the best choice for small arrays or arrays with irregular data distributions.

3. **Not as Fast as Binary Search:** While it's faster than linear search, Fibonacci search is not as fast as binary search for uniformly distributed data. Binary search is more efficient when the data is evenly divided.

4. **Complexity for Non-Numeric Data:** Adapting the Fibonacci search algorithm for non-numeric data can be more complex because it relies on numerical comparisons for index calculations.

### When to Use Fibonacci Search

Fibonacci search is a good choice under the following conditions:

1. **Large Sorted Arrays:** It's most effective when dealing with large sorted arrays where the overhead of binary search may be significant.

2. **Uniformly Distributed Data:** It works well when data is uniformly distributed within the array. In such cases, it can efficiently narrow down the search space.

3. **When Memory Usage Matters:** Fibonacci search is a good choice when you want an efficient search algorithm that doesn't require significant additional memory overhead.

4. **Situations Where Preprocessing is Feasible:** If you can efficiently pre-compute or calculate Fibonacci numbers in your application and store them for later use, it becomes a practical choice.

5. **In Situations Where Simplicity is Preferred:** Fibonacci search is relatively easy to understand and implement compared to more complex search algorithms.
