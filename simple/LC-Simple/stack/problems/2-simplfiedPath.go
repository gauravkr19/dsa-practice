// You are given an absolute path for a Unix-style file system, which always begins with a slash '/'.
// Your task is to transform this absolute path into its simplified canonical path.

// The rules of a Unix-style file system are as follows:

// A single period '.' represents the current directory.
// A double period '..' represents the previous/parent directory.
// Multiple consecutive slashes such as '//' and '///' are treated as a single slash '/'.
// Any sequence of periods that does not match the rules above should be treated as a valid directory or file name.
// For example, '...' and '....' are valid directory or file names.

// The simplified canonical path should follow these rules:
// The path must start with a single slash '/'.
// Directories within the path must be separated by exactly one slash '/'.
// The path must not end with a slash '/', unless it is the root directory.
// The path must not have any single or double periods ('.' and '..') used to denote current or parent directories.
// Return the simplified canonical path.

package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	simplifiedPath := []string{}
	dirs := strings.Split(path, "/")

	fmt.Println(dirs)

	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		}
		if dir != ".." {
			simplifiedPath = append(simplifiedPath, dir)
		} else if len(simplifiedPath) > 0 {
			simplifiedPath = simplifiedPath[:len(simplifiedPath)-1]
		}
	}
	fmt.Println(simplifiedPath)
	return "/" + strings.Join(simplifiedPath, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))
}
