// given an array of integers nums and integer target
// you want to do the two pass solution

// first pass you want to create a hash
// the key is the difference and the value is the index

// second pass is you loop through the list again
// in each iteration, check if the element is in the hash
// if it is, it means its the difference, and you can return a list
// with the value and the index


// i guess you can do it all in one pass
// so calculate the difference between the target - the current element
// you check your differenceMap, and you go is it there, if not you add in the {difference: currElementIndex}
// and then you keep checking
// if you get a match with the difference, you return the [val (prevIndex, i (currIndex)]
// so you only need to do one pass
