package main

import (
	"context"
	"github.com/CodingJzy/bahamut_example/model"
	"github.com/CodingJzy/bahamut_example/processor"
	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
)

func main() {

	// 模型管理者：由一个map组成，value是一个elemental.ModelManager接口
	modelManagers := map[int]elemental.ModelManager{
		0: model.Manager(),
	}

	// 实例化bahamut之前设置一些选项，比如profilingServer，restServer、healthServer、pushServer、认证、超时、模型等
	options := []bahamut.Option{
		bahamut.OptRestServer("0.0.0.0:5011"),
		bahamut.OptModel(modelManagers),
	}
	// New一个bahamut对象
	bs := bahamut.New(options...)

	// 循环所有的模型标识对象，然后一一注册到处理器。
	for _, ident := range model.AllIdentities() {
		p, ok := processor.AllProcess[ident.Name]
		if ok {
			_ = bs.RegisterProcessor(p, ident)
		}
	}

	// 运行bahamut
	bs.Run(context.Background())
}
