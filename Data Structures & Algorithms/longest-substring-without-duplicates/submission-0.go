func lengthOfLongestSubstring(s string) int {
	left , right := 0 ,0
	set := make(map[rune]struct{});
	lttrs := []rune(s);

	score := 0;

	for i := 0 ; i < len(lttrs) ; i++ {

		_ , exists := set[lttrs[i]];

		if exists {

			//i is the position that's the problem
			stillExists := true
			for stillExists {
				delete(set , lttrs[left]);
				left++;
				_ , isThere := set[lttrs[i]];
			  	stillExists = isThere;
			}

			set[lttrs[i]] = struct{}{};
			
		}else{
			right++;
			set[lttrs[i]] = struct{}{};
		}

		if len(set) > score {
				score = len(set) ;
			}
	}

	return score;
}
