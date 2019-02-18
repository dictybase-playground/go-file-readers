package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTxtAndGetIDs(t *testing.T) {
	r, err := readTxtAndGetIDs("../data/stalk.txt")
	if err != nil {
		t.Errorf("unable to read txt file %s", err)
	}
	assert := assert.New(t)
	assert.Equal(len(r), 44, "should match length of txt document")
	assert.Equal(r[0], "DDANAT:0000093", "should match first item of txt document")
	assert.Equal(r[len(r)-1], "DDANAT:0000053", "should match last item of txt document")
}

func TestConvertToBool(t *testing.T) {
	f := convertToBool("No")
	r := convertToBool("Yes")
	assert := assert.New(t)
	assert.False(f, "should convert No to false")
	assert.True(r, "should convert Yes to true")
}

func TestReadAndParseJSON(t *testing.T) {
	r, err := readAndParseJSON("../data/genes.json", "DDB_G0270756")
	if err != nil {
		t.Errorf("unable to read and parse json file %s", err)
	}
	assert := assert.New(t)
	assert.Equal(r.ID, "DDB_G0270756", "should match the passed in ID")
	assert.Equal(r.Type, "genes", "should match the type from the json file")
	assert.Equal(r.Attributes.SeqID, "DDB0232428", "should match the seqid from the json file")
	assert.Equal(r.Attributes.BlockID, "DDB0232428", "should match the block_id from the json file")
	assert.Equal(r.Attributes.Source, "dictyBase Curator", "should match the source from the json file")
	assert.Equal(r.Attributes.Start, 4446395, "should match the start from the json file")
	assert.Equal(r.Attributes.End, 4449415, "should match the end from the json file")
	assert.Equal(r.Attributes.Strand, "-", "should match the strand from the json file")
}
