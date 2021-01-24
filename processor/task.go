package processor

import (
	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
	"github.com/CodingJzy/bahamut_example/model"
)

func init() {
	p := &TaskProcessor{}
	AllProcess[model.TaskIdentity.Name] = p
}

type TaskProcessor struct {
	err    error
	Output interface{}
	events []*elemental.Event
}

func (p *TaskProcessor) ProcessRetrieveMany(ctx bahamut.Context) error {
	ctx.SetOutputData("hello task 查询所有")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *TaskProcessor) ProcessRetrieve(ctx bahamut.Context) error {
	ctx.SetOutputData("hello task 查询单个")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *TaskProcessor) ProcessCreate(ctx bahamut.Context) error {
	ctx.SetOutputData("hello task 创建")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *TaskProcessor) ProcessUpdate(ctx bahamut.Context) error {
	ctx.SetOutputData("hello task 更新")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *TaskProcessor) ProcessDelete(ctx bahamut.Context) error {
	ctx.SetOutputData("hello task 删除")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *TaskProcessor) ProcessPatch(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *TaskProcessor) ProcessInfo(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
