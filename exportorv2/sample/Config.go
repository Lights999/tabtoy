// Generated by github.com/davyxu/tabtoy
// Version: 2.4.0
// DO NOT EDIT!!
package table

import (
	"fmt"
)

type ActorType int32

const (
	ActorType_Fighter ActorType = 0

	ActorType_Power ActorType = 21
)

type Config struct {

	// Sample
	Sample []*SampleDefine

	// Exp
	Exp []*ExpDefine
}

type Prop struct {
	HP int32

	AttackRate float32

	ExType ActorType
}

type SampleDefine struct {

	// 唯一ID
	ID int64

	// 名称
	Name string

	// 图标ID
	IconID int32

	// 攻击率
	NumericalRate float32

	// 物品id
	ItemID int32

	// BuffID
	BuffID []int32

	// 类型
	Type ActorType

	// 技能ID列表
	SkillID []int32

	// 单结构解析
	SingleStruct *Prop

	// 字符串结构
	StrStruct []*Prop
}

type ExpDefine struct {

	// 唯一ID
	Level int32

	// 经验值
	Exp int32

	// 布尔检查
	BoolChecker bool

	// 类型
	Type ActorType
}

var SampleByID = make(map[int64]*SampleDefine)

var SampleByName = make(map[string]*SampleDefine)

var ExpByLevel = make(map[int32]*ExpDefine)

func init() {

	RegisterIndexEntry("Sample", func(content interface{}) {

		config := content.(*Config)

		// Sample
		for _, def := range config.Sample {

			if _, ok := SampleByID[def.ID]; ok {
				panic(fmt.Sprintf("duplicate index in SampleByID: %v", def.ID))
			}

			if _, ok := SampleByName[def.Name]; ok {
				panic(fmt.Sprintf("duplicate index in SampleByName: %v", def.Name))
			}

			SampleByID[def.ID] = def
			SampleByName[def.Name] = def

		}

	})

	RegisterIndexEntry("Exp", func(content interface{}) {

		config := content.(*Config)

		// Exp
		for _, def := range config.Exp {

			if _, ok := ExpByLevel[def.Level]; ok {
				panic(fmt.Sprintf("duplicate index in ExpByLevel: %v", def.Level))
			}

			ExpByLevel[def.Level] = def

		}

	})
}
