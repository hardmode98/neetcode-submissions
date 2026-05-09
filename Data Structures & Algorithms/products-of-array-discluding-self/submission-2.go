func productExceptSelf(nums []int) []int {

	output := make([]int , len(nums));

	for i := 0 ; i < len(nums) ; i++ {
		product:= 1;

		for j:= 0 ; j < len(nums) ; j++{
			if i != j{
				product = product * nums[j];
			}
		}

		output[i] = product
	}

	return output;
}
