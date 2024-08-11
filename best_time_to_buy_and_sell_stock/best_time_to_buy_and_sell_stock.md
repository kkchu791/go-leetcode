
// the pattern of this problem is you want to use 2 pointer
// and loop until the sellingPointer get out of the prices.length
// if the buyer pointer is greater than the sell pointer, you want to switch the buyer pointer w/ the seller pointer
// since theres no possiblility to make a profit, and you've found a new lowest point
// else the buyer pointer is less than the sell pointer, you have potential to make a profit
// but you don't know if you have the maxProfit, so you have to keep the maxProfit in storage and always compare with the currPorfit to see if you have the greatest value
// But you aren't sure if there are any more profits that could be max so you have to incrememnt your seller pointer
// You will do this pattern, until the seller pointer is out of the prices range
// and then return MaxProfit
// If you didn't find any profit, you've defaulted maxProfit to be 0 so you'll return 0.
// Time complexity O(n), cuz you only have to do through each number once
// Space complexity is O(1), because you've only used a constant number of variables to store data
// What if there are duplicates
