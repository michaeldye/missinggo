package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anacrolix/missinggo"
	"github.com/anacrolix/missinggo/itertools"
)

func TestEmptyBitmap(t *testing.T) {
	var bm Bitmap
	assert.False(t, bm.Contains(0))
	bm.Remove(0)
	it := itertools.NewIterator(&bm)
	assert.Panics(t, func() { it.Value() })
	assert.False(t, it.Next())
}

func bitmapSlice(bm *Bitmap) (ret []int) {
	sl := itertools.IterableAsSlice(bm)
	missinggo.CastSlice(&ret, sl)
	return
}

func TestSimpleBitmap(t *testing.T) {
	bm := new(Bitmap)
	assert.EqualValues(t, []int(nil), bitmapSlice(bm))
	bm.Add(0)
	assert.True(t, bm.Contains(0))
	assert.False(t, bm.Contains(1))
	bm.Add(3)
	assert.True(t, bm.Contains(0))
	assert.True(t, bm.Contains(3))
	assert.EqualValues(t, []int{0, 3}, bitmapSlice(bm))
	bm.Remove(0)
	assert.EqualValues(t, []int{3}, bitmapSlice(bm))
}

func TestSub(t *testing.T) {
	var left, right Bitmap
	left.Add(2, 5, 4)
	right.Add(3, 2, 6)
	assert.Equal(t, []int{4, 5}, Sub(&left, &right).ToSortedSlice())
	assert.Equal(t, []int{3, 6}, Sub(&right, &left).ToSortedSlice())
}

func TestSubUninited(t *testing.T) {
	var left, right Bitmap
	assert.EqualValues(t, []int(nil), Sub(&left, &right).ToSortedSlice())
}

func TestAddRange(t *testing.T) {
	var bm Bitmap
	bm.AddRange(21, 26)
	bm.AddRange(9, 14)
	bm.AddRange(11, 16)
	bm.Remove(12)
	assert.EqualValues(t, []int{9, 10, 11, 13, 14, 15, 21, 22, 23, 24, 25}, bm.ToSortedSlice())
	bm.Clear()
	bm.AddRange(3, 7)
	bm.AddRange(0, 3)
	bm.AddRange(2, 4)
	bm.Remove(3)
	assert.EqualValues(t, []int{0, 1, 2, 4, 5, 6}, bm.ToSortedSlice())
}
