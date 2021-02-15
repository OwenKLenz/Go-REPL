package packages

import (
  "fmt"
  "strings"
  "regexp"
)

func ParseForPackages(codeString string, currentPackages []string) ([]string, bool) {
  validImportMatcher, _ := regexp.Compile(`^import ["][\w/]+["](, ["][\w/]+["])*$`)
  importFinder, _ := regexp.Compile(`"([\w/]+)"`)
  foundPackages := []string{}
  newPackages := []string{}
  packagesFound := false

  if validImportMatcher.Match([]byte(codeString)) {
    foundPackages = importFinder.FindAllString(codeString, -1)
    packagesFound = true
  }

  for _, pack := range foundPackages {
    found := false
    for _, currentPack := range currentPackages {
      if pack == currentPack {
        found = true
      }
    }

    if !found {
      newPackages = append(newPackages, pack)
    } else {
      fmt.Printf("Package %s is already imported\n", pack)
    }
  }

  return newPackages, packagesFound
}

func FindUsedPackages(packages []string, statements []string) []string {
  usedPackages := []string{}
  joinedStatements := strings.Join(statements, "\n")

  if !(strings.Contains(joinedStatements, "\"fmt\"")) {
    usedPackages = append(usedPackages, "\"fmt\"")
  }

  for _, pack := range packages {
    fmt.Println(pack)
    match, _ := regexp.MatchString(fmt.Sprintf(`(^|[\W])%s[.]`, pack[1:len(pack) - 1]), joinedStatements)

    if match {
      usedPackages = append(usedPackages, pack)
    }
  }

  return usedPackages
}
