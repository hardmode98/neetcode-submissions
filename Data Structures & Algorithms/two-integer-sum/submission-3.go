func twoSum(nums []int , target int) []int {
	notebook := make(map[int]int);
	for index , value := range nums{
		compliment:= target -value;
		_ , exists := notebook[compliment];

		if !exists{
			notebook[value] = index;
		}else{
			return []int{ notebook[compliment],index};
		}
	}

	return []int{-1,-1}
}
