package core

import (
	"github.com/acmestack/godkits/assert"
	"testing"
)

func TestGetConfigFromFile(t *testing.T) {
	p := NewProperties()
	_ = p.GetConfigFromFile("props.properties")
	val, ok := p.Get("id")
	assert.IsTrue(t, ok)
	assert.Equal(t, 1, val)
	val2, ok2 := p.Get("name")
	assert.IsTrue(t, ok2)
	assert.Equal(t, "xiaoming", val2)
	val3, ok3 := p.Get("mm")
	assert.IsFalse(t, ok3)
	assert.IsTrue(t, val3 == "")
}
