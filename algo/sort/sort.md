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

### Space & Time Complexity

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

## Pigeon Sort

First is to find the Min and Max value inside the Unsorted Array then Calculate the Range of possible value by subtracting the Min Value inside the Unsorted Array by Max value Calculation of the range is because This determine the Amount of bucket needed for the Collection of Values inside the array then create an bucket array (pigeon holes) adn place the value of the original array inside that bucket then Iterate though the range of values and based on the the bucket hole index we will start adding the value inside the original array because it will be getting sorted starting at the beginning and do it step by step while emptying the pigeon hole bucket until nothing left

1. **Find the Min and Max Values**

- Begin by finding the minimum and maximum values within the unsorted array. You can do this by iterating through the array and keeping track of the minimum and maximum values.

2. **Calculate the Range of Possible Values**

- Calculate the range of possible values by subtracting the minimum value found in step 1 from the maximum value. This range represents the span of values that can be present in the input array.

3. **Create an Array of Pigeonholes (Buckets)**

- Create an array of pigeonholes (buckets) to accommodate the values within the range of possible values. The number of pigeonholes should be at least equal to the range of values calculated in step 2.

4. **Distribute Values into Pigeonholes**

- Iterate through the unsorted array, and for each element:
  - Calculate its position within the pigeonholes array based on its value within the range. This can be done by subtracting the minimum value from the element's value.
  - Place the element into the corresponding pigeonhole based on this position. If multiple elements have the same value, you can use a data structure (e.g., a linked list) within each pigeonhole to maintain their order.

5. **Collection Phase (Sorting)**

- Begin the collection phase by iterating through the pigeonholes array in order, from the smallest possible value to the largest.
- For each non-empty pigeonhole:
  - Collect the elements back into the original array, maintaining their order.
  - Continue collecting elements from the pigeonhole until it is empty.
- Repeat this process for all pigeonholes in ascending order.

6. **Final Result**

- After completing the collection phase, the original array will be sorted in ascending order.

The Pigeonhole Sort algorithm sorts elements by distributing them into buckets (pigeonholes) based on their values within a known range, and then collecting them back in ascending order. The critical steps involve calculating the range, creating pigeonholes, distributing values, and carefully collecting elements from the pigeonholes to achieve the desired sorting result.

```go
func Pigeon(array []int) ([]int, error) {
	if len(array) == 0 {
		return nil, errors.New("The Provided Array is Empty")
	}
	// Finding min and max with the help of helper function
	minValue, maxValue, minMaxErr := findMinMax(array)
	if minMaxErr != nil {
		return nil, minMaxErr
	}
	// Range of possible value - because This determine the Amount of bucket needed for the Collection
	var rangeOfValue int = maxValue - minValue + 1
	// we will create pigeon Holes based on range of value for possible collection of the Values
	pigeonHolesBucket := make([]int, rangeOfValue)
	// Iterate through the values of the original array
	for _, num := range array {
		// adding those values into the relative position
		// num-minValue := calculates the index of the pigeonhole where the current element should be placed.
		// and by incrementing ++ we will effectively place the each and every value inside the bucket
		pigeonHolesBucket[num-minValue]++
	}

	/// This part is collecting the element from the Bucket and sorting
	//  it and placing it back to original array as Sorted

	// index is to keep track index at value of the Original Array
	var index int = 0
	// this loop iterate through the range of Values inside the Pigeon Hole in Ascending Order (small to Large)
	for i := 0; i < rangeOfValue; i++ {
		// It collect the smallest element from the bucket till pigeon hole is empty
		for pigeonHolesBucket[i] > 0 {
			// array[index] := is the starting position of the Original array
			// i := is the index of the position of the current pigeon Hole Bucket
			// minValue := minimum value in the Original array
			// By adding (i + minValue) minVal to i, we "de-normalize" the value, converting
			// it back to its original value within the range.
			array[index] = i + minValue
			// Increment the Index for preparing the next element
			index++
			// Decrement the values from the bucket to original array till it's empty
			pigeonHolesBucket[i]--
		}
	}

	return array, nil
}

// Finding the min and max value inside the Unsorted Array
func findMinMax(array []int) (int, int, error) {
	// base case - if our array has nothing that It will return Nothing
	if len(array) == 0 {
		return 0, 0, errors.New("The Array is Empty: Nothing Found")
	}

	// Lets Assume that that the First element in the array is the min and also max

	var maxValue int = array[0]
	var minValue int = array[0]
	// We iterate though each and every value inside the array and compare it and assign the correct
	// and corresponding value to min and max
	for _, num := range array {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}

	return minValue, maxValue, nil
}
```

### Space & Time Complexity

| Complexity Type          | Worst Case    | Best Case     | Average Case  | Space Complexity |
|--------------------------|---------------|---------------|---------------|-------------------|
| Time Complexity          | O(n + N)      | O(n + N)      | O(n + N)      | O(N)              |
| Space Complexity         | O(N)          | O(N)          | O(N)          | O(N)              |

Explanation:

- **Time Complexity**:
  - In the worst case, Pigeonhole Sort has a time complexity of O(n + N), where 'n' is the number of elements in the input array, and 'N' is the range of possible values within the array.
  - In the best case and average case, Pigeonhole Sort still has a time complexity of O(n + N). This is because the distribution and collection phases both depend on the range 'N' and the number of elements 'n' in the array.

- **Space Complexity**:
  - The space complexity of Pigeonhole Sort is O(N) because it requires additional space to store the pigeonholes (buckets), and the number of pigeonholes is determined by the range 'N' of possible values.
  - This space complexity remains the same in both the worst case, best case, and average case.

Keep in mind that Pigeonhole Sort's efficiency heavily depends on the range of values 'N.' If 'N' is significantly larger than 'n' (the number of elements to be sorted), the algorithm may become inefficient due to the large number of empty pigeonholes, resulting in increased space usage and potentially slower sorting times.

## Cycle Sort

1. **Loop Through the Array**: Begin by setting up a loop to iterate through the unsorted array. This loop will help you process each element in the array.

2. **Select the Current Item**: For each iteration, select the first element in the array as the 'item' that you want to place in its correct sorted position. This element will also serve as your starting point.

3. **Find the Correct Position**: Initialize a 'position' variable to the current starting position (which is the index of the 'item' within the array). Then, loop through the remaining elements in the array, starting from the second element. Compare each element with the 'item' to determine the correct position for 'item.' If an element is smaller than 'item,' increment the 'position' variable.

4. **Skip If Already Sorted**: Check if 'item' is already in its correct sorted position by comparing 'position' with the initial starting position. If they are the same, it means 'item' is already in the right place, so you can move on to the next item.

5. **Find the Next Available Position**: To avoid overwriting existing elements in the array, find the next available position for 'item.' This involves looping through the array until you find a position where 'item' can be placed without overwriting any values.

6. **Swap the Elements**: Once you've identified the correct position and the next available position for 'item,' perform a swap. Swap 'item' with the element at the 'position.' This step effectively moves 'item' to its correct sorted position.

7. **Continue Swapping**: To ensure that 'item' reaches its final sorted position, continue swapping elements until 'item' cycles back to its initial starting position within the array. This step may involve multiple swaps.

```go
func Cycle(array []int) []int {
    // Loop through the unsorted array
    for initialPosition := 0; initialPosition < len(array)-1; initialPosition++ {
        // Select the first element as the 'item' to be placed in its correct position
        item := array[initialPosition]
        // Initialize 'position' as the current starting position for 'item'
        position := initialPosition

        // Find the correct position for the 'item'
        // by comparing it with elements starting from the second element
        for i := initialPosition + 1; i < len(array); i++ {
            // If the current element is smaller than 'item', increment 'position'
            if array[i] < item {
                position++
            }
        }

        // If 'item' is already in its correct position, move on to the next item
        // Ensures that smaller elements are already in their correct positions
        if position == initialPosition {
            continue
        }

        // Find the next available position for 'item' to avoid overwriting elements
        for item == array[position] {
            position++
        }

        // Swap 'item' with the element at 'position', placing 'item' in its correct position
        array[position], item = item, array[position]

        // Continue swapping until 'item' reaches its final sorted position
        for position != initialPosition {
            // Reset 'position' to the initial position
            position = initialPosition

            // Find the next available position for 'item'
            for i := initialPosition + 1; i < len(array); i++ {
                if array[i] < item {
                    position++
                }
            }

            // Find the next available position for 'item' to avoid overwriting elements
            for item == array[position] {
                position++
            }

            // Swap 'item' with the element at 'position' to continue its cycle
            array[position], item = item, array[position]
        }
    }
    return array
}
```

### Space & Time Complexity

| Complexity        | Worst Case    | Average Case  | Best Case     |
|-------------------|---------------|---------------|---------------|
| Time Complexity   | O(n^2)        | O(n^2)        | O(n^2)        |
| Space Complexity  | O(1)          | O(1)          | O(1)          |

Explanation:

- The worst-case time complexity is O(n^2) because in the worst case, each element may need to go through multiple swaps to reach its final sorted position.
- The average-case time complexity is also O(n^2) because it typically involves a significant number of comparisons and swaps.
- The best-case time complexity is O(n^2) as well because even if the array is already partially sorted, the algorithm still performs a substantial number of comparisons and swaps.
- The space complexity is constant O(1) because Cycle Sort is an in-place sorting algorithm that doesn't require additional memory allocation that depends on the input size.

## Comb Sort

Comb Sort is a sorting algorithm that combines the principles of Bubble Sort and the concept of dynamically changing gap sizes. This guide will walk you through the key steps of the Comb Sort algorithm, providing explanations along the way.

 1.**Initialization**

- Create an array of elements that you want to sort.

 2. **Swapped Flag**

- Initialize a boolean flag, 'swapped,' as 'true.' This flag will help us track whether any element swaps occurred during a pass.

 3. **Main Loop**

- Enter a loop that continues until 'swapped' is 'false' or the array length is reduced to one. You cannot swap a single-element array.

4. **Calculate Gap**

- Calculate the 'gap' value for this iteration using the 'getGap' function. The gap dynamically shrinks during each pass.

 5. **Reset Swapped Flag**

- Set the 'swapped' flag to 'false' at the start of each pass. We assume no swaps have occurred initially.

 6. **Iterate Through Array**

- Loop through the array, comparing and potentially swapping elements that are 'gap' positions apart.

 7. **Compare and Swap**

- If the element at the current 'index' is greater than the element at 'index+gap,' swap these elements.
- This step involves moving elements toward their correct positions in the sorted order.

 8. **Track Swaps**

- Set the 'swapped' flag to 'true' if any swapping occurred during this pass.

 9. **Gap Adjustment**

- Repeat the process with a smaller gap in the next iteration.
- The 'getGap' function calculates the new gap size, adapting it for better efficiency.

 10. **Repeat Loop**

- Continue the main loop until 'swapped' is 'false' or the array length is reduced to one.

 11. **Sorting Completion**

- After the loop finishes, the array is sorted.

### The Gap Factor

The "gap" is a crucial concept in the Comb Sort algorithm, and it plays a significant role in the algorithm's performance. It determines how far apart elements are compared and swapped during each pass of the algorithm. The gap is important for the following reasons:

**1. Influence on Sorting Efficiency:**

- The gap size directly impacts the algorithm's sorting efficiency. A larger gap allows for faster movement of elements across the array but may result in fewer comparisons and swaps during each pass. A smaller gap, on the other hand, ensures more thorough comparisons but requires more passes to complete the sorting process.

**2. Adaptive Shrinkage:**

- One of the unique features of the Comb Sort algorithm is that the gap size starts large and progressively shrinks with each pass. This adaptive shrinkage is essential for the algorithm's efficiency because it combines the benefits of both large and small gap sizes.
- Initially, with a large gap, the algorithm can quickly move larger elements toward the end of the array, similar to the way Bubble Sort works. This reduces the total number of larger elements that need to be moved later.
- As the gap decreases, the algorithm performs more fine-grained comparisons and swaps, ensuring that smaller elements are placed in their correct positions.

**3. Finding an Optimal Gap:**

- Choosing an optimal gap size is essential for optimizing the performance of Comb Sort. The choice of gap and the shrink factor (typically around 1.3) can significantly affect the sorting speed.
- Experimenting with different gap sizes and shrink factors is often necessary to find the best combination for a specific dataset. Different datasets may benefit from different gap sizes, so it's not a one-size-fits-all parameter.

**4. Time Complexity Considerations:**

- The time complexity of Comb Sort depends on the choice of the gap size. In its worst-case scenario, if the gap is set as 1, Comb Sort degenerates into a Bubble Sort, resulting in a time complexity of O(n^2). However, with a properly chosen gap, it can achieve an average-case time complexity of approximately O(n*log(n)).

```go
func Comb(array []int) {
	// Get the length of the input array
	n := len(array)

	// Set the initial gap to the length of the array
	gap := n

	// Define the shrink factor for adjusting the gap
	shrinkFactor := 1.3

	// Initialize a flag to track whether any swaps occurred
	swapped := true

	// Continue looping until the gap is greater than 1 or swaps occurred
	for gap > 1 || swapped {
		// Calculate the gap using the shrink factor
		gap = int(float64(gap) / shrinkFactor)

		// Ensure that the gap is at least 1 to prevent division by zero
		if gap < 1 {
			gap = 1
		}

		// Initialize the swapped flag to false for this pass
		swapped = false

		// Compare and potentially swap elements with the calculated gap
		for index := 0; index < n-gap; index++ {
			if array[index] > array[index+gap] {
				// Swap elements if the condition is met
				array[index], array[index+gap] = array[index+gap], array[index]

				// Set the swapped flag to true to indicate a swap occurred
				swapped = true
			}
		}
	}
}
```

| Complexity         | Worst Case                              | Average Case                         | Best Case                                | Space Complexity |
|--------------------|----------------------------------------|-------------------------------------|------------------------------------------|------------------|
| Time Complexity    | O(n^2)                                   | O(n^2)                                | O(n*log(n)) [Optimal Gap Sequence]        | O(1)             |
| Explanation        | In the worst case, when the gap is set to 1 and the array is in reverse order or nearly so, Comb Sort becomes inefficient and takes quadratic time. | In most practical cases, Comb Sort's adaptive gap size reduction provides better performance than Bubble Sort but still exhibits quadratic time complexity. | In the best case, when an optimal gap sequence is used and data is partially sorted, Comb Sort approaches a time complexity of n*log(n). Achieving this best-case performance is challenging in practice. | Comb Sort is an in-place sorting algorithm, which means it uses a constant amount of extra memory for variables like the gap and loop indices. |

## Bubble Sort

1. Initialize `n` with the length of the input slice `arr`. `n` represents the number of elements in the slice.

2. Initialize a boolean variable `swapped` to `true`. This variable will be used to check whether any swaps were made in a pass.

3. Enter a `for` loop that continues as long as `swapped` is `true`. This loop ensures that the sorting process continues until no more swaps are made in a pass.

4. Within the loop, set `swapped` to `false`. This flag is initially true and will only be set to false if no swaps are made in the current pass.

5. Start another `for` loop that iterates through the slice from index 1 to `n-1`. We compare adjacent elements and swap them if they are in the wrong order.

6. If the comparison (`arr[i-1] > arr[i]`) is true, swap the elements using a parallel assignment. This is where the actual swapping of elements occurs.

7. Set `swapped` to `true` to indicate that a swap was made in this pass.

8. After the inner loop completes, one pass through the array is done. Now, decrement `n` to reduce the range of elements to consider in the next pass.

```go
func Bubble(array []int) []int {
	var swapped bool = true
	// entire length of the array
	var n int = len(array)
	// continues Loop until all the elements inside the array is swapped or Sorted
	for swapped {
		// we assume that the array is not sorted
		swapped = false
		// loop thorough the array at position second
		for i := 1; i < n; i++ {
			// array[i-1]: the first element of the array
			// array[i]: second element of the array
			// Checking if first element of the array is greater than the second element than we need to swap
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				// outer loop will continue till this condition reached
				swapped = true
			}
		}
		// After each pass, the largest element is at the end, so we reduce the range
		n--
	}
	return array
}

// Recursive Bubble

func RecursiveBubble(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}
	for i := 1; i < n; i++ {
		if array[i-1] > array[i] {
			array[i-1], array[i] = array[i], array[i-1]
		}
	}
	RecursiveBubble(array[:n-1])
	return array
}
```

### Space & Time Complexity

| Bubble Sort      | Best Case | Average Case | Worst Case  | Space Complexity |
|------------------|---------------------------|------------------------------|---------------------------|------------------|
| Bubble Sort      | O(n)                      | O(n^2)                       | O(n^2)                    | O(1)             |

Explanation:

- Best Case Time Complexity: O(n) - This is when the input array is already sorted, and Bubble Sort makes only one pass to confirm the array is sorted.
- Average Case Time Complexity: O(n^2) - This is the expected time complexity when elements are randomly ordered, and Bubble Sort makes multiple passes.
- Worst Case Time Complexity: O(n^2) - This occurs when the input array is sorted in reverse order, and Bubble Sort makes the maximum number of passes with many swaps.
- Space Complexity: O(1) - Bubble Sort is an in-place sorting algorithm, meaning it doesn't require additional memory proportional to the input size. It has a constant space complexity.
