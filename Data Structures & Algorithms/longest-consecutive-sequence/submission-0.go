
func longestConsecutive(nums []int) int {

	notebook := make(map[int]struct{});
	maxVal := 0;

	for _ , value := range nums {
		notebook[value] = struct{}{};
	}

	for _ , num := range nums {

		chainLength := 1;
		nextNum := num;
		for true {
			
			_ , exists  := notebook[nextNum+1];

			if exists {
				chainLength++;
				nextNum++;
			}else{
				maxVal = max(maxVal , chainLength);
				break;
			}
 		

		}
	}

	return maxVal;

}
