package earthshaker

type (
	//定义消息
	Message interface {
	}

	//读写超时 等事件
	Event interface {
	}

	//与通道相关的数据
	Attachment interface {
	}

	//ActiveHandler
	//InboundHandler
	//OutboundHandler
	//ExceptionHandler
	//InactiveHandler
	//EventHandler
	Handler interface {
	}
)
