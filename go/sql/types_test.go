/*
   Copyright 2016 GitHub Inc.
	 See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package sql

import (
	"testing"

	"github.com/openark/golib/log"
	"github.com/stretchr/testify/require"
)

func init() {
	log.SetLevel(log.ERROR)
}

func TestParseColumnList(t *testing.T) {
	names := "id,category,max_len"

	columnList := ParseColumnList(names)
	require.Equal(t, 3, columnList.Len())
	require.Equal(t, []string{"id", "category", "max_len"}, columnList.Names())
	require.Equal(t, 0, columnList.Ordinals["id"])
	require.Equal(t, 1, columnList.Ordinals["category"])
	require.Equal(t, 2, columnList.Ordinals["max_len"])
}

func TestGetColumn(t *testing.T) {
	names := "id,category,max_len"
	columnList := ParseColumnList(names)
	{
		column := columnList.GetColumn("category")
		require.NotNil(t, column)
		require.Equal(t, column.Name, "category")
	}
	{
		column := columnList.GetColumn("no_such_column")
		require.Nil(t, column)
	}
}

func TestBinaryToString(t *testing.T) {
	id := []uint8{0x1b, 0x99}
	col := make([]interface{}, 1)
	col[0] = id
	cv := ToColumnValues(col)

	require.Equal(t, "1b99", cv.StringColumn(0))
}

func TestConvertArgCharsetDecoding(t *testing.T) {
	latin1Bytes := []uint8{0x47, 0x61, 0x72, 0xe7, 0x6f, 0x6e, 0x20, 0x21}

	col := Column{
		Charset: "latin1",
		charsetConversion: &CharacterSetConversion{
			FromCharset: "latin1",
			ToCharset:   "utf8mb4",
		},
	}

	// Should decode []uint8
	str := col.convertArg(latin1Bytes, false)
	require.Equal(t, "Garçon !", str)
}
