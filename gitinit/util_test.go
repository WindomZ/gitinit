package gitinit

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestValidGitAddress(t *testing.T) {
	assert.False(t, ValidGitAddress("gitinit"))
	assert.False(t, ValidGitAddress("WindomZ/gitinit"))
	assert.False(t, ValidGitAddress("github.com/WindomZ/gitinit"))
	assert.False(t, ValidGitAddress("https://github.com/WindomZ/gitinit"))
	assert.True(t, ValidGitAddress("https://github.com/WindomZ/gitinit.git"))
	assert.True(t, ValidGitAddress("git@github.com:WindomZ/gitinit.git"))
}
