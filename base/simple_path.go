package base

/*
special char: ' [ ] * .
wildcard: '
*/

type SimplePath []any

func ParseSimplePath(p string) (SimplePath, error) {

}

type SimpleNodeType int

const (
	SimpleNodeTypeObjectKey SimpleNodeType = -iota
	SimpleNodeTypeIndex
	SimpleNodeTypeIndexStar
	SimpleNodeTypeObjectKeyStar
)

func nextToken(s []byte) (SimpleNodeType, string, []byte, error) {
	if s[0] == '[' {
		if len(s) == 1 {
			return SimpleNodeTypeObjectKey, "[", nil, nil
		}
		if s[1] == '*'
	}
}
