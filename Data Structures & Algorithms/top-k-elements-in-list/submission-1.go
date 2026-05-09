

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int);
	for i := 0 ; i < len(nums) ; i++ {
		freqMap[nums[i]]++;
	}

	highrise := make([][]int , len(nums) + 1);

	for key , value := range freqMap{
		highrise[value] = append(highrise[value], key);
	}


	result := []int{};
	for i := len(highrise) - 1 ; i >= 0 ; i-- {
		for _ , num := range highrise[i]{
			result = append(result, num);
			if(len(result) == k){
				return result
			}
		}
	}

	return result;


}
