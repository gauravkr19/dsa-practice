// function mergeSort(arr, left, right):
//     if left >= right:
//         return  # Base case: Single element is already sorted

//	   mid = (left + right) / 2  # Find the middle point

/* 	Question: why not left == right
However, using left >= right is just a safer defensive check:
Normally, the recursive calls will always follow left == right at the base case.
> But in some edge cases (like if the function is misused or incorrect arguments are passed), left > right could occur.
> left >= right prevents infinite recursion in such cases.
*/

//     # Recursively sort left half
//     mergeSort(arr, left, mid)

//     # Recursively sort right half
//     mergeSort(arr, mid + 1, right)

//     # Merge sorted halves
//     merge(arr, left, mid, right)

// function merge(arr, left, mid, right):
//     temp = []
//     i = left      # Pointer for left subarray
//     j = mid + 1   # Pointer for right subarray

//     # Merge elements from both halves into temp in sorted order
//     while i <= mid and j <= right:
//         if arr[i] <= arr[j]:
//             temp.append(arr[i])
//             i += 1
//         else:
//             temp.append(arr[j])
//             j += 1

//     # Copy any remaining elements from left subarray
//     while i <= mid:
//         temp.append(arr[i])
//         i += 1

//     # Copy any remaining elements from right subarray
//     while j <= right:
//         temp.append(arr[j])
//         j += 1

//     # Copy back the sorted temp array into original array
//     for k in range(left, right + 1):
//         arr[k] = temp[k - left]

// Explanation
// 1. Divide Phase
// The array is recursively divided into two halves until each subarray has only one element.
// [38, 27, 43, 3, 9, 82, 10]
// First Split:
// Left:  [38, 27, 43, 3]
// Right: [9, 82, 10]
// Further Splitting:
// Left:  [38, 27]   [43, 3]
// Right: [9, 82]    [10]

// Left:  [38] [27]  [43] [3]
// Right: [9] [82]   [10]
// Now, every subarray has only one element (base case).

// 2. Conquer Phase

// MergeSort(arr, 0, 6)
// ├── MergeSort(arr, 0, 3)
// │   ├── MergeSort(arr, 0, 1)
// │   │   ├── MergeSort(arr, 0, 0) → returns
// │   │   ├── MergeSort(arr, 1, 1) → returns
// │   │   ├── merge(arr, 0, 0, 1) → merges [38, 27] → [27, 38]
// │   ├── MergeSort(arr, 2, 3)
// │   │   ├── MergeSort(arr, 2, 2) → returns
// │   │   ├── MergeSort(arr, 3, 3) → returns
// │   │   ├── merge(arr, 2, 2, 3) → merges [43, 3] → [3, 43]
// │   ├── merge(arr, 0, 1, 3) → merges [27, 38] and [3, 43] → [3, 27, 38, 43]
// ├── MergeSort(arr, 4, 6)
// │   ├── MergeSort(arr, 4, 5)
// │   │   ├── MergeSort(arr, 4, 4) → returns
// │   │   ├── MergeSort(arr, 5, 5) → returns
// │   │   ├── merge(arr, 4, 4, 5) → merges [9, 82] → [9, 82]
// │   ├── MergeSort(arr, 6, 6) → returns
// │   ├── merge(arr, 4, 5, 6) → merges [9, 82] and [10] → [9, 10, 82]
// ├── merge(arr, 0, 3, 6) → merges [3, 27, 38, 43] and [9, 10, 82] → [3, 9, 10, 27, 38, 43, 82]
