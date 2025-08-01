package simplepath

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type PathNode struct {
	NodeType NodeType
	Value    string
}

type SimplePath []PathNode

func Parse(s string) (SimplePath, error) {
	buf := []byte(s)
	needDot := false
	var p SimplePath
	for len(buf) > 0 {
		if needDot {
			buf = buf[1:]
		}
		needDot = true
		var nodeType NodeType
		var val string
		var err error
		nodeType, val, buf, err = nextToken(buf)
		if err != nil {
			return nil, err
		}
		p = append(p, PathNode{nodeType, val})
	}
	return p, nil
}

func nextToken(s []byte) (NodeType, string, []byte, error) {
	if len(s) == 0 {
		return NodeTypeInvalid, "", nil, errors.New("unexpected end of path")
	}
	if s[0] == '[' {
		return nextIndex(s)
	} else {
		return nextKey(s)
	}
}

func nextIndex(s []byte) (NodeType, string, []byte, error) {
	index := bytes.IndexByte(s, ']')
	if index == -1 {
		return NodeTypeInvalid, "", nil, errors.New("] expected")
	}
	if index == 2 && s[1] == '*' {
		return NodeTypeIndexStar, "", s[index+1:], nil
	}
	index, err := strconv.Atoi(string(s[1:index]))
	if err != nil {
		return NodeTypeInvalid, "", nil, err
	}
	if index < 0 {
		return NodeTypeInvalid, "", nil, errors.New("index out of range")
	}
	return NodeTypeIndex, strconv.Itoa(index), s[index+1:], nil
}

func nextKey(s []byte) (NodeType, string, []byte, error) {
	lenS := len(s)
	if lenS >= 1 {
		if s[0] == '*' {
			if lenS == 1 || s[1] == '.' {
				return NodeTypeKeyStar, "", s[1:], nil
			} else {
				return NodeTypeInvalid, "", nil, errors.New("key star expected")
			}
		}
	}
	key := bytes.NewBuffer(nil)
	i := 0
	for ; i < len(s); i++ {
		if s[i] == '.' {
			break
		}
		if s[i] == '\'' {
			if i+1 >= len(s) {
				return NodeTypeInvalid, "", nil, errors.New("invalid escape sequence")
			}
			key.WriteByte(s[i])
			i++
			continue
		}
		key.WriteByte(s[i])
	}
	return NodeTypeKey, key.String(), s[i:], nil
}

func (s SimplePath) AppendIndex(index int) SimplePath {
	return append(s, PathNode{
		NodeType: NodeTypeIndex,
		Value:    strconv.Itoa(index),
	})
}

func (s SimplePath) AppendKey(key string) SimplePath {
	return append(s, PathNode{
		NodeType: NodeTypeKey,
		Value:    key,
	})
}

func (s SimplePath) Match(target SimplePath) bool {
	if len(s) != len(target) {
		return false
	}
	for i, srcNode := range s {
		dstNode := target[i]
		if srcNode != dstNode {
			return false
		}
	}
	return true
}

func (s SimplePath) Get(data any) (any, bool) {
	if len(s) == 0 {
		return data, true
	}
	node := s[0]
	switch node.NodeType {
	case NodeTypeKey:
		mapVal, ok := data.(map[string]any)
		if !ok {
			return nil, false
		}
		data, ok = mapVal[node.Value]
		if !ok {
			return nil, false
		}
		return s[1:].Get(data)
	case NodeTypeIndex:
		arrayVal, ok := data.([]any)
		if !ok {
			return nil, false
		}
		index, _ := strconv.Atoi(node.Value)
		if index < 0 || index >= len(arrayVal) {
			return nil, false
		}
		return s[1:].Get(arrayVal[index])
	default:
		return nil, false
	}
}

func escapeKey(key string) string {
	buf := bytes.NewBuffer(nil)
	for _, ch := range key {
		switch ch {
		case '\'', '*', '[', ']', '.':
			buf.WriteByte('\'')
			buf.WriteRune(ch)
		default:
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func (s SimplePath) String() string {
	buf := bytes.NewBuffer(nil)
	for i, node := range s {
		if i != 0 {
			buf.WriteByte('.')
		}
		switch node.NodeType {
		case NodeTypeInvalid:
			return "invalid_path"
		case NodeTypeIndex:
			_, _ = fmt.Fprintf(buf, "[%s]", node.Value)
		case NodeTypeKey:
			buf.WriteString(escapeKey(node.Value))
		case NodeTypeKeyStar:
			buf.WriteString("*")
		case NodeTypeIndexStar:
			buf.WriteString("[*]")
		}
	}
	return buf.String()
}

type NodeType int

const (
	NodeTypeInvalid NodeType = iota
	NodeTypeIndex
	NodeTypeKey
	NodeTypeIndexStar
	NodeTypeKeyStar
)
