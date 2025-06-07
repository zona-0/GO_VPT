package modules

// TODO: Insertion Sort for sorting job listings by salary
func SortBySalary() {
	var i, j, maxIdx int
	var temp JobListing

	for i = 0; i < jobCount-1; i++ {
		maxIdx = i

		for j = i + 1; j < jobCount; j++ {
			if jobListings[j].Salary > jobListings[maxIdx].Salary {
				maxIdx = j
			}
		}
		temp = jobListings[i]
		jobListings[i] = jobListings[maxIdx]
		jobListings[maxIdx] = temp
	}
}

// TODO: Selection sort for sorting job list by relevance
func SortByReleveance() {
	var i, j, maxIdx int
	var temp JobListing

	for i = 0; i < jobCount-1; i++ {
		maxIdx = i
		for j = i + 1; j < jobCount; j++ {
			if jobListings[j].Relevance > jobListings[maxIdx].Relevance {
				maxIdx = j
			}
		}
		if maxIdx != i {
			temp = jobListings[i]
			jobListings[i] = jobListings[maxIdx]
			jobListings[maxIdx] = temp
		}
	}
}
