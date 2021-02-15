package main

import (
  // "strings"
  "fmt"
  "regexp"
  // term "github.com/nsf/termbox-go"
)

func scanForVariableChanges(codeString string) bool {
  fullVarDeclaration, _ := regexp.Compile(`^var\s+\w+(,\s+\w+)*\s+\w+($|(\s+=\s+.+$))`)
  walrusVarDeclaration, _ := regexp.Compile(`^\w+(\s*,\s*\w+\s*)*\s*:=\s*\w+($|((\s*,\s*\w+\s*)*$))`)
  reassignment, _ := regexp.Compile(`^\w+(,\s*\w+)*\s*=\s*.+`)

  if fullVarDeclaration.MatchString(codeString) ||
		 walrusVarDeclaration.MatchString(codeString) ||
		 reassignment.MatchString(codeString) {
    return true
  } else {
    return false
  }
}

Array of variable declarations
	- each declaration needs a 

1. Whenever a variable declaration occurs,
func findUsedVariables(variables []string



// map[string] []string {
//   "a": []
// }

// type Variable struct {

// }

// [a := 1,
//  a = a + 2,
 

// 1. Require single variable declarations
//   - Parse them into a map with the variable's value stored
//     - If walrus declaration, set map[varName] = declarationstring
//     - If var x type declaration, set value to 
//   - Whenever a "x = y" statement is sent, if map[x] is set,  
//   - Print variables to file by iterating over the map's keys/values

// 2. Multivariable assignment
//   - Extract


// text := thing
// text, thing := one, two
// var thing string

func main() {
  // term.Init()

	// for i := 0; i < 3; i++ {
  //   // ev := term.PollEvent();
  //   fmt.Println("yes")
  //   // fmt.Println(ev.Type);
  // }

  // var one, two, three int
  // one, two:=1,2
  // fmt.Println(one, two)
  // fmt.Println(strings.TrimSpace("  123 "))
  // fmt.Println(strings.Split("hi there world", " "))
  fmt.Println(scanForVariableDeclarations("one,two = 1,2"))
}
