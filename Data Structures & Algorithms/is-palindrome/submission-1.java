class Solution {
    public boolean isPalindrome(String s) {
        int left = 0;
        s = s.trim().toLowerCase();
        int right = s.length()-1;

        boolean valid = true;

        while(left < right){

            if(!isComparable(s.charAt(left))){
                left++;
                continue;
            }

             if(!isComparable(s.charAt(right))){
                right--;
                continue;
            }

            if(!(s.charAt(left) == s.charAt(right))){
               break; 
            }else{
                left++;
                right--;
            }


        }
        

        return left >= right;


    }

    public boolean isComparable(char i){
        if((i>='0' && i<= '9') || (i >= 'A' && i<= 'Z') || (i >= 'a' && i<= 'z')){
            return true;
        }else{
            return false;
        }
    }

}
