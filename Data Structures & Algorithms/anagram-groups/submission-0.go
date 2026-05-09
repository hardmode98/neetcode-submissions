import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {

	groups := [][]string{}
	for i := 0 ; i < len(strs) ; i++{

		word := strs[i]
        foundGroup := false

		for j:= 0 ; j < len(groups) ; j++{
			for k:=0 ; k < len(groups[j]) ; k++ {
				if isAnagram(strs[i] , groups[j][k]){
					groups[j] = append(groups[j], strs[i]);
					foundGroup = true;
					break
				}
			}
		}

		// 2. If it didn't match any existing group, create a new one!
        if !foundGroup {
            groups = append(groups, []string{word})
        }
	}


	return groups;

}

func isAnagram(first string , second string) bool {
	if(len(first) != len(second)){
		return false;
	}

	first = sortString(first);
	second = sortString(second);

	if first == second {
		return true
	}

	return false;
}

//function sorts strings
func sortString(word string) string{
	chars := []rune(word);
	slices.Sort(chars);
	return string(chars);
}