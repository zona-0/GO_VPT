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

func SortByReleveance() {
	var i, j int
	var temp JobListing

	for i = 0; i < jobCount-1; i++ {
		for j = 0; j < jobCount-i-1; j++ {
			if jobListings[j].Relevance < jobListings[j+1].Relevance {
				temp = jobListings[j]
				jobListings[j] = jobListings[j+1]
				jobListings[j+1] = temp
			}
		}
	}
}
