package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistFunc(t *testing.T) {
	tests := []map[string][]float64{
		{"x": []float64{1, 3, 0, 8}, "y": []float64{5, 2, 0, 3}},
		{"x": []float64{2, 4, 0, 1}, "y": []float64{0, 2, 2, 3}},
		{"x": []float64{0, 3, 1, 8}, "y": []float64{5, 0, 0, 3}},
	}

	// TEST1: binary.
	distFunc := DistFunc("binary")
	want := []float64{0, 0.5, 0.75}
	_, err := distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Binary distance is not correct")
	}

	// TEST2: canberra.
	distFunc = DistFunc("canberra")
	want = []float64{1.32, 2.83, 3.45}
	_, err = distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Canberra distance is not correct")
	}

	// TEST3: jaccard.
	distFunc = DistFunc("jaccard")
	want = []float64{0.63, 0.73, 0.83}
	_, err = distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Jaccard distance is not correct")
	}

	// TEST4: manhattan.
	distFunc = DistFunc("manhattan")
	want = []float64{10, 8, 14}
	_, err = distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Manhattan distance is not correct")
	}

	// TEST5: maximum.
	distFunc = DistFunc("maximum")
	want = []float64{5, 2, 5}
	_, err = distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Maximum distance is not correct")
	}

	// TEST5: euclidean.
	distFunc = DistFunc("euclidean")
	want = []float64{6.48, 4, 7.75}
	_, err = distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Euclidean distance is not correct")
	}

	// TEST5: default = euclidean.
	distFunc = DistFunc("")
	want = []float64{6.48, 4, 7.75}
	_, err = distFunc([]float64{1, 2}, []float64{0, 3, 5})
	assert.NotNil(t, err, "Vectors of different length should return an error")
	for i, test := range tests {
		dist, testErr := distFunc(test["x"], test["y"])
		assert.Nil(t, testErr, "Valid input vectors should not return an error")
		assert.InDeltaf(t, want[i], dist, 0.01, "Euclidean distance is not correct")
	}
}
