package variables

import (
  "regexp"
  "sort"
)

type Variable struct {
  Statements []*Statement
  InUse bool
}

func (v *Variable) setInUse() {
  v.InUse = true
}

func (v *Variable) setNotInUse() {
  v.InUse = false
}

func NewVariable(statement *Statement) *Variable {
  return &Variable{Statements: []*Statement{statement}, InUse: false,}
}

type Statement struct {
  Text string
  Num int
}

func NewStatement(text string, num int) *Statement {
  return &Statement{
    Text: text,
    Num: num,
  }
}

func NewVariableMap() *map[string]*Variable {
  return &map[string]*Variable{}
}

func GetStatements(newStatement *Statement, variableMap *map[string]*Variable) ([]string, []string) {
  var usedVariables []string
  var currentStatement *Statement

  statementQueue := []*Statement{newStatement}
  statementList := []*Statement{newStatement}

  for len(statementQueue) > 0 {
    currentStatement = statementQueue[0]

    for variableName, variableStruct := range *variableMap {
      if !variableStruct.InUse && variableInStatement(variableName, currentStatement.Text) {
        statementQueue = append(statementQueue, variableStruct.Statements...)
        statementList = append(variableStruct.Statements, statementList...)
        (*variableMap)[variableName].setInUse()
        usedVariables = append(usedVariables, variableName)
      }
    }

    statementQueue = statementQueue[1:]
  }

	sortStatements(statementList)

	statementStrings := statementsToStrings(statementList)

  return statementStrings, usedVariables
}

func statementsToStrings(statementList []*Statement) []string {
  var statementStrings []string

  for _, statement := range statementList {
    statementStrings = append(statementStrings, statement.Text)
  }

  return statementStrings
}

func variableInStatement(variableName string, statement string) bool {
  searchRegexp, _:= regexp.Compile(`(^|\s+|\W)` + variableName + `(\s|\W|$)`)
  found := searchRegexp.MatchString(statement)
  return found
}

// Sorting
type byInputOrder []*Statement

func (s byInputOrder) Len() int {
  return len(s)
}

func (s byInputOrder) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s byInputOrder) Less(i, j int) bool {
  return s[i].Num < s[j].Num
}

func sortStatements(statements []*Statement) {
  sort.Sort(byInputOrder(statements))
}

func FindChangedVariable(codeString string) string {
  assignmentMatch, _ := regexp.Compile(`^\s*(\w+)\S*\s*\S?=`)

	variable := assignmentMatch.FindStringSubmatch(codeString)

  if len(variable) == 0 {
    declarationMatch, _ := regexp.Compile(`^\s*var\s+(\S+)\s`)
    variable = declarationMatch.FindStringSubmatch(codeString)
  }

  if len(variable) == 0 {
    panic("Invalid variable creation/alteration syntax")
  }

  return variable[1]
}

func ResetVariables(variableList []string, variableMap *map[string]*Variable) {
  for _, name := range variableList {
    (*variableMap)[name].setNotInUse()
  }
}


// Initialize a statementQueue to contain the input statement
// For each statement in queue (until queue is empty),
//   Iterate over all variable names
//     If variable not "inUse"
//       Search statement for Variable name
//       If found,
//       Append variable's statements to queue and set Variable to inUse
//       Append variable to array of usedVariables (to reset them to not inUse at end)
//    Prepend statement to statementList
//  Return statementList and usedVariables
