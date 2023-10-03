# Sorting Algorithms

## Table of Content

| Sorting Algorithm          | Difficulty Level     |
|---------------------------|----------------------|
| Bingo Sort Algorithm      | Very Easy            |
| Pigeonhole Sort           | Very Easy            |
| Cycle Sort                | Easy                 |
| Comb Sort                 | Easy                 |
| Bubble Sort               | Easy                 |
| Gnome Sort                | Easy                 |
| Odd-Even Sort / Brick Sort| Easy                 |
| Cocktail Sort             | Easy                 |
| Bitonic Sort              | Moderate             |
| Pancake Sorting           | Moderate             |
| Sleep Sort                | Moderate             |
| Structure Sorting         | Moderate             |
| Tag Sort (To get both sorted and original) | Moderate |
| Selection Sort            | Moderate             |
| Insertion Sort            | Moderate             |
| Tree Sort                 | Moderate             |
| Stooge Sort               | Moderate             |
| Heap Sort                 | Challenging          |
| ShellSort                 | Challenging          |
| Merge Sort                | Challenging          |
| Quick Sort                | Challenging          |
| TimSort                   | Challenging          |
| 3-way Merge Sort          | Challenging          |
| Counting Sort             | Very Challenging     |
| Radix Sort                | Very Challenging     |
| Bucket Sort               | Very Challenging     |

## Table of Complexity Comparison

Here's a table comparing various sorting algorithms based on their complexity in markdown format:

| Sorting Algorithm            | Best Case    | Worst Case   | Average Case | Memory | Stable | Method Used              |
| ---------------------------- | ------------ | ------------ | ------------ | ------ | ------ | ------------------------ |
| Selection Sort               | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Selection                |
| Bubble Sort                  | O(n)         | O(n^2)       | O(n^2)       | O(1)   | Yes    | Swapping                 |
| Insertion Sort               | O(n)         | O(n^2)       | O(n^2)       | O(1)   | Yes    | Insertion                |
| Merge Sort                   | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging                  |
| Quick Sort                   | O(n log n)   | O(n^2)       | O(n log n)   | O(log n) | No   | Partitioning             |
| Heap Sort                    | O(n log n)   | O(n log n)   | O(n log n)   | O(1)   | No     | Selection (Heapify)      |
| Counting Sort                | O(n + k)     | O(n + k)     | O(n + k)     | O(k)   | Yes    | Counting                 |
| Radix Sort                   | O(nk)        | O(nk)        | O(nk)        | O(n + k) | Yes   | Distribution             |
| Bucket Sort                  | O(n^2)       | O(n^2)       | O(n^2)       | O(n)   | Yes    | Distribution             |
| Bingo Sort Algorithm         | O(n^2)       | O(n^2)       | O(n^2)       | O(n)   | No     | Combining and Sorting    |
| ShellSort                    | O(n log^2 n) | O(n log^2 n) | O(n log^2 n) | O(1)   | No     | Insertion with Gaps      |
| TimSort                      | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging and Insertion    |
| Comb Sort                    | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Swapping with Gaps       |
| Pigeonhole Sort              | O(n + N)     | O(n^2)       | O(n + N)     | O(N)   | Yes    | Distribution             |
| Cycle Sort                   | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Cyclic Swapping          |
| Cocktail Sort                | O(n)         | O(n^2)       | O(n^2)       | O(1)   | Yes    | Bidirectional Bubble Sort|
| Strand Sort                  | O(n^2)       | O(n^3)       | O(n^2)       | O(n)   | No     | Merging                  |
| Bitonic Sort                 | O(n log^2 n) | O(n log^2 n) | O(n log^2 n) | O(log n) | Yes   | Comparisons and Swaps    |
| Pancake Sorting              | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Flipping                 |
| BogoSort or Permutation Sort | O((n+1)!)    | O(∞)         | O((n+1)!)    | O(1)   | No     | Random Permutations      |
| Gnome Sort                   | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | Yes    | Insertion with BackStep  |
| Sleep Sort – The King of Laziness | O(n)     | O(n)         | O(n)         | O(n)   | No     | Sleep and WakeUp         |
| Structure Sorting            | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging and Insertion    |
| Stooge Sort                  | O(n^(log 3)) | O(n^(log 3)) | O(n^(log 3)) | O(1)   | No     | Recursive Trisection     |
| Tag Sort (To get both sorted and original) | O(n log n) | O(n log n) | O(n log n) | O(n)   | No | Tagging and Sorting     |
| Tree Sort                    | O(n log n)   | O(n^2)       | O(n log n)   | O(n)   | Yes    | Binary Search Tree       |
| Odd-Even Sort / Brick Sort   | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | Yes    | Odd-Even Comparison      |
| 3-way Merge Sort            | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging (3-way)         |

## 1. Bingo Sort

First find the Min value inside the array, First Loop through the array from 0 which is starting out index and loop till the
end of unsorted Array then that value will be our min, basically we are sorting based on index so we will called it minIndex
Now inside that loop we will write another loop that iterate through the array again but this time it will be the second value of the array then we will compare second value of the array with the minIndex value if it's smaller than we set that second value at minIndex, after that we will swap the elements, like swapping the position at each and every index

1. **Initialize MinIndex:** Start by initializing a variable called `minIndex` to 0. This will represent the index of the minimum value in the array.

2. **Outer Loop:** Begin an outer loop that will iterate through the array from index 0 (the start of the unsorted portion of the array) to the end of the unsorted portion. This loop will determine the `minIndex` and the element to be swapped.

3. **Inner Loop:** Inside the outer loop, start an inner loop. This loop will iterate through the array again but starting from the second element (index + 1). The goal here is to compare each element with the current `minIndex` value and potentially update `minIndex`.

4. **Compare and Update `minIndex`:** In the inner loop, compare the element at the current index with the element at `minIndex`. If the element is smaller than the element at `minIndex`, update `minIndex` to the current index.

5. **Swap Elements:** After finding the `minIndex`, exit the inner loop. Swap the element at the current outer loop index with the element at `minIndex`. This places the smallest unsorted element in its correct position.

6. **Repeat:** Continue the outer loop, incrementing the index by 1 each time. This will find the next smallest element and place it in the next position in the sorted portion of the array.

7. **Finish:** Once the outer loop has iterated through the entire array, the array will be sorted in ascending order. All elements will be in their correct positions.

This sorting algorithm effectively finds the minimum value in the array based on index positions and swaps elements to sort the array.

```go
func Bingo(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		// Find the index of the minimum unsorted element
		minIndex := findMinUnsorted(arr, i, n)

		// Swap the minimum unsorted element with the current element (i)
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

// Helper function to find the index of the minimum unsorted element in the array
func findMinUnsorted(arr []int, startIndex, endIndex int) int {
	minIndex := startIndex

	for i := startIndex + 1; i < endIndex; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}
```

Here's the time complexity, best case, worst case, average case, and auxiliary space complexity of the `Bingo` sorting algorithm:

| Complexity Type            | Complexity      |
|----------------------------|-----------------|
| Time Complexity            | O(n^2)          |
| Best Case Time Complexity  | O(n^2)          |
| Worst Case Time Complexity | O(n^2)          |
| Average Case Time Complexity | O(n^2)        |
| Space Complexity           | O(1) (In-Place) |

Explanation:

- **Time Complexity**: The time complexity of Bingo Sort is O(n^2), where "n" is the number of elements in the array. This is because it involves two nested loops, where the outer loop runs "n" times, and the inner loop runs up to "n" times in the worst case, resulting in a quadratic time complexity.

- **Best Case Time Complexity**: In the best case, the time complexity is still O(n^2) because the algorithm always compares and potentially swaps elements, even if the array is already sorted.

- **Worst Case Time Complexity**: In the worst case, the time complexity is also O(n^2) when the input array is in reverse order, as the algorithm will make a significant number of comparisons and swaps.

- **Average Case Time Complexity**: On average, the time complexity remains O(n^2) because the algorithm does not employ any special strategies to reduce the number of comparisons or swaps.

- **Space Complexity**: The space complexity of Bingo Sort is O(1) because it operates in-place, meaning it does not require additional memory that grows with the size of the input. It only uses a constant amount of additional memory for variables like `minIndex` and `i`.
