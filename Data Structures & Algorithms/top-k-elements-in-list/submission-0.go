import(
	"slices"
	"cmp"
)
type pair struct {
	num int
	count int
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int);
	for i := 0 ; i < len(nums) ; i++ {
		freqMap[nums[i]]++;
	}

	pairs := make([]pair, 0, len(freqMap));

	for key , value := range freqMap{
		pairs = append(pairs , pair{num:key , count:value});
	}

	//descending order
	slices.SortFunc(pairs , func(a , b pair)int{
		return cmp.Compare(b.count , a.count);
	})

	result := make([]int, k);

	for i:= 0 ; i < k ;i++{
		result[i] = pairs[i].num
	}


	return result;


}
