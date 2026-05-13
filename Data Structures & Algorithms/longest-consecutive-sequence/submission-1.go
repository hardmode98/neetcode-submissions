func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	notebook := make(map[int]struct{})
	for _, value := range nums {
		notebook[value] = struct{}{}
	}

	maxVal := 0

	// Tip: Iterate over the map keys instead of the 'nums' array. 
	// This skips duplicate numbers, saving slightly more time.
	for num := range notebook {

		// OPTIMIZATION: Only start counting if 'num' is the start of a sequence.
		// If num-1 exists, this is a middle/end piece, so we skip it.
		if _, hasPrev := notebook[num-1]; !hasPrev {
			
			chainLength := 1
			nextNum := num

			// Now we count upward from the start of the sequence
			for {
				if _, hasNext := notebook[nextNum+1]; hasNext {
					chainLength++
					nextNum++
				} else {
					maxVal = max(maxVal, chainLength)
					break
				}
			}
		}
	}

	return maxVal
}