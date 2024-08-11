func maxprofit(prices []int) int {
	maxprofit := 0
	buypointer := 0
	sellpointer := 1

	while sellpointer < len(prices) {
		if prices[buypointer] >= prices[sellpointer] {
			buyerpointer = sellpointer
			sellpointer = buypointer + 1
		} else {
               		currprofit = prices[sellpointer] - prices[buypointer]
			if currprofit > maxprofit {
				maxprofit = currprofit
			}

			sellpointer++
	        }
	}

	return maxprofit
}


