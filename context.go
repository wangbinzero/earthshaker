package earthshaker

type (
	HandlerContext interface {
		Channel() Channel
		Handler() Handler
		Write(message Message)
		Close(err error)
		Trigger(event Event)
		Attachment() Attachment
		SetAttachment(Attachment)
	}

	//活跃
	ActiveContext interface {
		HandlerContext
		HandleActive()
	}

	//入栈
	InboundContext interface {
		HandlerContext
		HandleRead(message Message)
	}

	//出栈
	OutboundContext interface {
		HandlerContext
		HandleWrite(message Message)
	}

	//异常
	ExceptionContext interface {
		HandlerContext
		HandleException(e Exception)
	}

	//不活跃
	InactiveContext interface {
		HandlerContext
		HandleInactive(e Exception)
	}

	//事件
	EventContext interface {
		HandlerContext
		HandleEvent(event Event)
	}
)
