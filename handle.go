package gonetty

type ChannelHandle interface {
}

type DefaultChannelHandleContext struct {
	name     string
	pipeline Pipeline
	h        ChannelHandle
	next     ChannelHandlerContext
	prev     ChannelHandlerContext
}

func (d DefaultChannelHandleContext) Handle() ChannelHandle {
	return d.h
}

func (d DefaultChannelHandleContext) FireChannelActive() ChannelInboundInvoker {
	var currentCtx = d.next
	for currentCtx != nil {
		switch currentCtx.handle().(type) {
		case InboundHandler:
		case OutboundHandler:

		}
		currentCtx = d.next
	}
	return nil
}

func (DefaultChannelHandleContext) FireChannelInactive() ChannelInboundInvoker {
	//TODO implement me
	panic("implement me")
}

func (DefaultChannelHandleContext) FireChannelRead() ChannelInboundInvoker {
	//TODO implement me
	panic("implement me")
}

func (DefaultChannelHandleContext) FireChannelReadComplete() ChannelInboundInvoker {
	//TODO implement me
	panic("implement me")
}

type ChannelHandlerContext interface {
	handle() ChannelHandle
	ChannelInboundInvoker
}
type ChannelInboundInvoker interface {
	FireChannelActive() ChannelInboundInvoker
	FireChannelInactive() ChannelInboundInvoker
	FireChannelRead() ChannelInboundInvoker
	FireChannelReadComplete() ChannelInboundInvoker
}
type Message interface {
}

type InboundHandler interface {
	ChannelHandle
	ChannelRead(ctx ChannelHandlerContext, message Message)
}
type OutboundHandler interface {
	ChannelHandle
	ChannelWrite(ctx ChannelHandlerContext, message Message)
}
