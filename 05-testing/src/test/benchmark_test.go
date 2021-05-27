package test

import (
    "testing"
    "math/rand"
    "btree/src/btree"
    "btree/src/sortedSlice"
)

const maxUint = ^uint(0) 
const maxInt = int(maxUint >> 1) 

func Benchmark_Tree_Random1000Values(b *testing.B) {
    count := 1000
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertBenchmarkTree(b, rand, count)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertBenchmarkTree(b, rand, count)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertBenchmarkTree(b, rand, count)
    })
}

func Benchmark_Tree_Random10000Values(b *testing.B) {
    count := 10000
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertBenchmarkTree(b, rand, count)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertBenchmarkTree(b, rand, count)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertBenchmarkTree(b, rand, count)
    })
}

func Benchmark_Tree_LowCardinality10000Values(b *testing.B) {
    count := 10000
    cardinality := 4
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertBenchmarkTree(b, rand, count, cardinality)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertBenchmarkTree(b, rand, count, cardinality)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertBenchmarkTree(b, rand, count, cardinality)
    })
}

func Benchmark_Tree_Random1000Values_WithDeletions(b *testing.B) {
    count := 1000
    deletePctChance := 33
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertDeleteBenchmarkTree(b, rand, count, deletePctChance)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertDeleteBenchmarkTree(b, rand, count, deletePctChance)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertDeleteBenchmarkTree(b, rand, count, deletePctChance)
    })
}

func Benchmark_Tree_Random10000Values_WithDeletions(b *testing.B) {
    count := 10000
    deletePctChance := 33
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertDeleteBenchmarkTree(b, rand, count, deletePctChance)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertDeleteBenchmarkTree(b, rand, count, deletePctChance)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertDeleteBenchmarkTree(b, rand, count, deletePctChance)
    })
}

func Benchmark_Slice_Random1000Values(b *testing.B) {
    count := 1000
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertBenchmarkSlice(b, rand, count)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertBenchmarkSlice(b, rand, count)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertBenchmarkSlice(b, rand, count)
    })
}

func Benchmark_Slice_Random10000Values(b *testing.B) {
    count := 10000
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertBenchmarkSlice(b, rand, count)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertBenchmarkSlice(b, rand, count)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertBenchmarkSlice(b, rand, count)
    })
}

func Benchmark_Slice_LowCardinality10000Values(b *testing.B) {
    count := 10000
    cardinality := 4
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertBenchmarkSlice(b, rand, count, cardinality)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertBenchmarkSlice(b, rand, count, cardinality)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertBenchmarkSlice(b, rand, count, cardinality)
    })
}

func Benchmark_Slice_Random1000Values_WithDeletions(b *testing.B) {
    count := 1000
    deletePctChance := 33
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertDeleteBenchmarkSlice(b, rand, count, deletePctChance)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertDeleteBenchmarkSlice(b, rand, count, deletePctChance)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertDeleteBenchmarkSlice(b, rand, count, deletePctChance)
    })
}

func Benchmark_Slice_Random10000Values_WithDeletions(b *testing.B) {
    count := 10000
    deletePctChance := 33
    b.Run("Seed=0", func (b *testing.B) {
        rand := rand.New(rand.NewSource(0))
        runInsertDeleteBenchmarkSlice(b, rand, count, deletePctChance)
    })
    b.Run("Seed=108409", func (b *testing.B) {
        rand := rand.New(rand.NewSource(108409))
        runInsertDeleteBenchmarkSlice(b, rand, count, deletePctChance)
    })
    b.Run("Seed=-9999", func (b *testing.B) {
        rand := rand.New(rand.NewSource(-9999))
        runInsertDeleteBenchmarkSlice(b, rand, count, deletePctChance)
    })
}

func createValuesToInsert(rand *rand.Rand, count int, maxRand ...int) []int {
    values := make([]int, count)
    max := maxInt
    if len(maxRand) > 0 {
        max = maxRand[0]
    }
    for i := 0; i < count; i++ {
        values[i] = rand.Intn(max)
    }
    return values
}

func runInsertBenchmarkTree(b *testing.B, rand *rand.Rand, count int, maxRand ...int) {
    values := createValuesToInsert(rand, count, maxRand...)
    b.ResetTimer()
    tree := new(btree.AvlTree)
    for _, v := range values {
        tree.Insert(v)
    }
}

func runInsertDeleteBenchmarkTree(b *testing.B, rand *rand.Rand, count int, deletePctChance int, maxRand ...int) {
    values := createValuesToInsert(rand, count, maxRand...)
    b.ResetTimer()
    tree := new(btree.AvlTree)
    for _, v := range values {
        tree.Insert(v)
    }
    for _, v := range values {
        if rand.Intn(100) < deletePctChance {
            tree.Delete(v)
        }
    }
}

func runInsertBenchmarkSlice(b *testing.B, rand *rand.Rand, count int, maxRand ...int) {
    values := createValuesToInsert(rand, count, maxRand...)
    b.ResetTimer()
    slice := new(sortedSlice.SortedSlice)
    for _, v := range values {
        slice.Insert(v)
    }
}

func runInsertDeleteBenchmarkSlice(b *testing.B, rand *rand.Rand, count int, deletePctChance int, maxRand ...int) {
    values := createValuesToInsert(rand, count, maxRand...)
    b.ResetTimer()
    slice := new(sortedSlice.SortedSlice)
    for _, v := range values {
        slice.Insert(v)
    }
    for _, v := range values {
        if rand.Intn(100) < deletePctChance {
            slice.Delete(v)
        }
    }
}