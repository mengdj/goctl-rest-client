// Package rest
// @file:utility.go
// @description:
// @date: 02/24/2023
// @version:1.0.0
// @author: mengdj<mengdj@outlook.com>
package rest

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/lang"
	nurl "net/url"
	"strings"
)

func buildFormQuery(u *nurl.URL, val map[string]interface{}) nurl.Values {
	query := u.Query()
	for k, v := range val {
		query.Add(k, fmt.Sprint(v))
	}
	return query
}

func fillPath(u *nurl.URL, val map[string]interface{}) error {
	used := make(map[string]lang.PlaceholderType)
	fields := strings.Split(u.Path, slash)

	for i := range fields {
		field := fields[i]
		if len(field) > 0 && field[0] == colon {
			name := field[1:]
			ival, ok := val[name]
			if !ok {
				return fmt.Errorf("missing path variable %q", name)
			}
			value := fmt.Sprint(ival)
			if len(value) == 0 {
				return fmt.Errorf("empty path variable %q", name)
			}
			fields[i] = value
			used[name] = lang.Placeholder
		}
	}

	if len(val) != len(used) {
		for key := range used {
			delete(val, key)
		}

		var unused []string
		for key := range val {
			unused = append(unused, key)
		}

		return fmt.Errorf("more path variables are provided: %q", strings.Join(unused, ", "))
	}

	u.Path = strings.Join(fields, slash)
	return nil
}
