package processor

import (
	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
	"github.com/CodingJzy/bahamut_example/model"
)

func init() {
	p := &ListProcessor{}
	AllProcess[model.ListIdentity.Name] = p
}

type ListProcessor struct {
	err    error
	Output interface{}
	events []*elemental.Event
}

func (p *ListProcessor) ProcessRetrieveMany(ctx bahamut.Context) error {
	ctx.SetOutputData("hello list 查询所有")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *ListProcessor) ProcessRetrieve(ctx bahamut.Context) error {
	ctx.SetOutputData("hello list 查询单个")
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *ListProcessor) ProcessCreate(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *ListProcessor) ProcessUpdate(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *ListProcessor) ProcessDelete(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *ListProcessor) ProcessPatch(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
func (p *ListProcessor) ProcessInfo(ctx bahamut.Context) error {
	ctx.SetOutputData(p.Output)
	ctx.EnqueueEvents(p.events...)
	return p.err
}
