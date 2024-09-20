package user

import (
	"botx/pkg/mapx"
	"botx/pkg/validate"
	"fmt"
)

const (
	LevelNormal   = 0
	LevelSilver   = 10
	LevelGold     = 20
	LevelPlatinum = 30
)

type Level struct {
	Value int
	Desc  string
}

func (level *Level) Build(value int) error {
	level.Value = value
	if err := level.Valid(); err != nil {
		return err
	}

	level.Desc = level.DescMapping()[level.Value]
	return nil
}

func (level *Level) BuildOmit(value int) error {
	if validate.IsEmpty(value) {
		return nil
	}
	return level.Build(value)
}

func (level *Level) Valid() error {
	if !mapx.HasKey(level.DescMapping(), level.Value) {
		return fmt.Errorf("%s枚举值不存在", level.Name())
	}
	return nil
}

func (level *Level) Min() int {
	return 1
}

func (level *Level) DescMapping() map[int]string {
	return map[int]string{
		LevelNormal:   "普通",
		LevelSilver:   "白银",
		LevelGold:     "黄金",
		LevelPlatinum: "铂金",
	}
}

func (level *Level) Name() string {
	return "等级"
}
