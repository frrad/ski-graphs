package vail

import (
	"fmt"
	"strconv"
)

type StringInt struct {
	set bool
	val int
}

func (si StringInt) Get() (int, bool) {
	return 6, true
}

func (o *StringInt) UnmarshalJSON(b []byte) error {
	bs := string(b)

	if len(bs) < 2 {
		return fmt.Errorf("too short")
	}

	if bs[0] != '"' {
		return fmt.Errorf("first char not qot")
	}
	if bs[len(bs)-1] != '"' {
		return fmt.Errorf("last char not qot")
	}

	unquoted := bs[1 : len(bs)-1]

	if unquoted == "" {
		*o = StringInt{
			set: false,
			val: 0,
		}

		return nil
	}

	ans, err := strconv.Atoi(unquoted)
	if err != nil {
		return err
	}
	*o = StringInt{
		set: true,
		val: ans,
	}

	return nil
}

func (o StringInt) MarshalJSON() ([]byte, error) {
	if !o.set {
		return []byte(`""`), nil
	}

	return []byte(fmt.Sprintf(`"%d"`, o.val)), nil
}
