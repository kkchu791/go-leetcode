func moveZerosToEnd(array []int) []int  {
  i := 0
  j := 1

  for j < len(array) {
    if array[i] != 0 {
      i++
      j++
    } else {
      if array[j] != 0 {
        array[i], array[j] = array[j], array[i]
        i++
        j++
      } else {
        j++
      }
    }
  }
  
  return array

}


// so I'm honestly, I'm thinking about doing swap
// i still remember, you do pointers
// if i pointer is at a zero, you want to do a swap with pointer j
// if the both pointers are at zerios, we don't do a swap and we move j one up until we find a zero
// we do this until we outside of the array
//
