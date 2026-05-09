func hasDuplicate(nums []int) bool {
    // 1. Early exit: a set with 0 or 1 elements cannot have duplicates
    if len(nums) < 2 {
        return false
    }

    // 2. Pre-allocate memory: if you know the potential size, 
    // provide it to make() to avoid multiple re-allocations.
    seen := make(map[int]struct{}, len(nums))

    for _, n := range nums {
        // 3. The "Comma OK" check
        if _, exists := seen[n]; exists {
            return true
        }
        // 4. Use struct{}{} to save memory
        seen[n] = struct{}{}
    }

    return false
}