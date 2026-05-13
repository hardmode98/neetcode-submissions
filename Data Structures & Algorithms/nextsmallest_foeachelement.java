 public static int[] nextSmallerPrices(int[] spotPrices) {
        
        HashMap<Integer , Integer> map = new HashMap<>();
        Deque<Integer> stack = new ArrayDeque<>();
        int[] result = new int[spotPrices.length];
        Arrays.fill(result, -1);

        for (int i =0 ; i< spotPrices.length ; i++){
            
            while (!stack.isEmpty() && (spotPrices[stack.peek()] > spotPrices[i])){
                result[stack.pop()] = spotPrices[i];
            }

            stack.push(i);
        }


        return result;
    }
