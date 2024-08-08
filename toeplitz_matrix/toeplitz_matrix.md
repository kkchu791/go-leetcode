// what are the conditions, you can't really go too far out
// how do we know wheter r and c represent some element arr[r][c] in the array? Nested loop allows us to do that
// how many different left to right diagonal are there? Three
// if we wanted to solve for one dimensional array [1, 1, 1, 1], how might we solve this problem? We could loop through the array and compare each each element to the first element.
// what is the time and complexity? since its nested for loop, O(r *c) and O(1)

// we can loop through and make sure it matches the first element

//[[1,2,3,4],
 //[5,1,2,3],
 //[6,5,1,2]]

 // one hint is like you only do one small check in a loop in order to make sure its valid
 // and be careful of the border numbers, because  when they are at the border, its almost like they are always valid.
 
 // the check you need to do is current must equal bottom right
 // and you just need to loop through the whole array
 // one note is that you want to loop up to a point before the border numbers, it'll save some time but its not necessary to loop to to the border
 // but you can also add a check, if current and bottom right equals OR bottom right equals undefined, than its valid, otherwise you can
 // immediately return false
