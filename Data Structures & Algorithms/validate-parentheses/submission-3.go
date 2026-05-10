func isValid(s string) bool {

	stack := []rune{};
	letters := []rune(s);

	pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

	for i := 0; i< len(s) ; i++{

		ltr := letters[i];


		
		
		
		if ltr == '(' || ltr == '{' || ltr == '[' {
			stack = append(stack , ltr);
		}else {

			if len(stack) == 0 {
                return false
            }

			if pairs[ltr] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1];
		}else{
			return false;
		}
		}

	}

	return len(stack) == 0;
}
