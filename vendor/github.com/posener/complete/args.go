package complete

import (
	"os"
	"path/filepath"
)

// Args describes command line arguments
type Args struct {
	// All lists of all arguments in command line (not including the command itself)
	All []string
	// Completed lists of all completed arguments in command line,
	// If the last one is still being typed - no space after it,
	// it won't appear in this list of arguments.
	Completed []string
	// Last argument in command line, the one being typed, if the last
	// character in the command line is a space, this argument will be empty,
	// otherwise this would be the last word.
	Last string
	// LastCompleted is the last argument that was fully typed.
	// If the last character in the command line is space, this would be the
	// last word, otherwise, it would be the word before that.
	LastCompleted string
}

// Directory gives the directory of the current written
// last argument if it represents a file name being written.
// in case that it is not, we fall back to the current directory.
func (a Args) Directory() string {
	if info, err := os.Stat(a.Last); err == nil && info.IsDir() {
		return fixPathForm(a.Last, a.Last)
	}
	dir := filepath.Dir(a.Last)
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		return "./"
	}
	return fixPathForm(a.Last, dir)
}

func newArgs(line []string) Args {
	completed := removeLast(line[1:])
	return Args{
		All:           line[1:],
		Completed:     completed,
		Last:          last(line),
		LastCompleted: last(completed),
	}
}

func (a Args) from(i int) Args {
	if i > len(a.All) {
		i = len(a.All)
	}
	a.All = a.All[i:]

	if i > len(a.Completed) {
		i = len(a.Completed)
	}
	a.Completed = a.Completed[i:]
	return a
}

func removeLast(a []string) []string {
	if len(a) > 0 {
		return a[:len(a)-1]
	}
	return a
}

func last(args []string) (last string) {
	if len(args) > 0 {
		last = args[len(args)-1]
	}
	return
}
