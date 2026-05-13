 // ---------------------------------------------------------------
    // IMPLEMENT THIS METHOD
    //
    // For every price in spotPrices, return the next chronological price
    // that is strictly LESS than it. If no later price is smaller, put -1
    // in that slot. Return an int[] of the same length as spotPrices.
    //
    // Example: [4, 8, 5, 2, 7]  ->  [2, 5, 2, -1, -1]
    // ---------------------------------------------------------------
    public static int[] nextSmallerPrices(int[] spotPrices) {
        // The stack holds INDICES of prices that are still waiting for an
        // answer — i.e. we haven't yet seen a smaller price after them.
        //
        // Why indices and not values? Because values aren't unique. For
        // [3, 3, 1, 3, 3, 0], a map keyed by value would overwrite the
        // answer for the first pair of 3s when the second pair gets resolved
        // by 0 — all four 3s would end up pointing to 0, even though the
        // first pair's correct answer is 1.
        //
        // Indices ARE unique, so each waiting position gets written exactly
        // once into its own slot in `result`. We compare values to decide
        // whether to pop (spotPrices[stack.peek()] > spotPrices[i]) but
        // store indices so we know which result slot to fill.
        Deque<Integer> stack = new ArrayDeque<>();
        int[] result = new int[spotPrices.length];
        Arrays.fill(result, -1);

        for (int i = 0; i < spotPrices.length; i++) {
            // Current price resolves every waiting index whose value is bigger.
            while (!stack.isEmpty() && spotPrices[stack.peek()] > spotPrices[i]) {
                result[stack.pop()] = spotPrices[i];
            }
            // Index i is now waiting for its own next-smaller price.
            stack.push(i);
        }
        // Anything still on the stack never found a smaller price — stays -1.
        return result;
    }
