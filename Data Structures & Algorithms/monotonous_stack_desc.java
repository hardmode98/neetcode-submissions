public static int[] nextGreaterPrices(int[] historicalPrices, int[] queryPrices) {
        
        HashMap<Integer, Integer> map = new HashMap<>();
        Deque<Integer> stack = new ArrayDeque<>();
        int[] result = new int[queryPrices.length];


        for (int i =0 ; i < historicalPrices.length; i++) {

            if(stack.isEmpty()){
                stack.push(historicalPrices[i]);
                continue;
            }
            
             while (!stack.isEmpty() && (stack.peek() < historicalPrices[i])){
                    
                    map.put(stack.pop(), historicalPrices[i]);
                }

            stack.push(historicalPrices[i]);
            
            
            
        }

        for (int i =0 ; i < queryPrices.length ; i++){
            result[i] = map.getOrDefault(queryPrices[i], -1);
        }

         return result;
    }
