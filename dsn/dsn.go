package dsn

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type DSN map[string]string

func (dsn DSN) String() string {

	keys := make([]string, 0)

	for k, _ := range dsn {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	pairs := make([]string, 0)

	for _, k := range keys {

		pair := fmt.Sprintf("%s=%s", k, dsn[k])
		pairs = append(pairs, pair)
	}

	return strings.Join(pairs, " ")
}

func StringToDSN(str_dsn string) (DSN, error) {

	str_dsn = strings.Trim(str_dsn, " ")
	dsn := make(map[string]string)

	for _, str_pair := range strings.Split(str_dsn, " ") {

		pair := strings.Split(str_pair, "=")

		if len(pair) != 2 {
			return nil, errors.New("Invalid DSN string")
		}

		k := pair[0]
		v := pair[1]

		_, ok := dsn[k]

		if ok {
			msg := fmt.Sprintf("'%s' key already set", k)
			return nil, errors.New(msg)
		}

		dsn[k] = v
	}

	return dsn, nil
}

func StringToDSNWithKeys(str_dsn string, keys ...string) (DSN, error) {

	dsn, err := StringToDSN(str_dsn)

	if err != nil {
		return nil, err
	}

	for _, k := range keys {

		_, ok := dsn[k]

		if !ok {
			msg := fmt.Sprintf("DSN is missing '%s=' key", k)
			return nil, errors.New(msg)
		}
	}

	return dsn, nil
}
