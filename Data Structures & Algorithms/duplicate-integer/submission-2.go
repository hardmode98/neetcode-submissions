func hasDuplicate(nums []int) bool {
    
	if(len(nums) <= 1){
		 return false;
	}

	var freq = make(map[int]struct{});

	for i:= 0; i< len(nums); i++ {
		_ , exists := freq[nums[i]];
		if exists {
			return true;
		}else{
			freq[nums[i]] = struct{}{}
		}
	}

	return false;

}
