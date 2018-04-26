// Package hclust contains methods for clustering an input matrix.
package hclust

// Distance generates a square matrix of distance values calculated between row
// vectors of an input matrix. Setting tranpose to true will calculate the distance matrix for
// column vectors instead. Distance metric options are: binary, Canberra, Euclidean, Jaccard,
// Manhattan and maximum.
func Distance(matrix [][]float64, metric string, transpose bool) (dist [][]float64) {
	// Get distance function.
	distFunc := DistFunc(metric)

	// Transpose matrix if requested.
	if transpose {
		matrix = Transpose(matrix)
	}

	// Init distance matrix.
	dim := len(matrix)
	dist = make([][]float64, dim) // Set row capacity.
	for i := range dist {
		dist[i] = make([]float64, dim) // Set column capacity.
	}

	// Iterate over rows.
	for i := range matrix {
		dist[i][i] = 0
		for j := i + 1; j < dim; j++ {
			elementDist, _ := distFunc(matrix[i], matrix[j])
			dist[i][j] = elementDist
			dist[j][i] = elementDist
		}
	}

	return
}
