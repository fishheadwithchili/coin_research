package processing_ingredients

type Algo interface {
	// GetCurrentValueToInt64 get current value after algorithm
	// param previousResults:the previous algorithm results
	// param currentOriginValue: the current origin value before algorithm
	// return current result after algorithm
	GetCurrentValueToInt64(previousResults []int64, currentOriginValue int64) int64
	// GetReferenceCount the count of previous results, return const
	GetReferenceCount() int
}
