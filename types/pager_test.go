package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagerValidate(t *testing.T) {
	t.Run("Should return error by invalid offset", func(t *testing.T) {
		pager := Pager{
			Offset: -10,
			Limit:  -20,
		}
		err := pager.Validate()
		assert.Error(t, err)
	})

	t.Run("Should return error by invalid limit", func(t *testing.T) {
		pager := Pager{
			Offset: 1,
			Limit:  -20,
		}
		err := pager.Validate()
		assert.Error(t, err)
	})
}

func TestPagerResultGetResult(t *testing.T) {
	pagerResult := PagerResult{}
	pager := Pager{
		Offset: 0,
		Limit:  10,
	}
	pages := pagerResult.GetResult(&pager, 0, []string{})
	assert.Equal(t, int64(0), pages.Next)
	assert.Equal(t, int64(0), pages.Previous)
	assert.Equal(t, int64(0), pages.Total)
	assert.Empty(t, pages.Data)
}
