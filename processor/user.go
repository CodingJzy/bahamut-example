package processor

import (
	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
	"github.com/CodingJzy/bahamut_example/model"
)

func init() {
	p := &UserProcessor{}
	AllProcess[model.UserIdentity.Name] = p
}

type UserProcessor struct {
	err    error
	Output interface{}
	events []*elemental.Event
}

func (p *UserProcessor) ProcessRetrieveMany(ctx bahamut.Context) error {
	ctx.SetOutputData("hello user 查询所有")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *UserProcessor) ProcessRetrieve(ctx bahamut.Context) error {
	ctx.SetOutputData("hello user 查询单个")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *UserProcessor) ProcessCreate(ctx bahamut.Context) error {
	ctx.SetOutputData("hello user 创建")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *UserProcessor) ProcessUpdate(ctx bahamut.Context) error {
	ctx.SetOutputData("hello user 更新")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *UserProcessor) ProcessDelete(ctx bahamut.Context) error {
	ctx.SetOutputData("hello user 删除")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *UserProcessor) ProcessPatch(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *UserProcessor) ProcessInfo(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
