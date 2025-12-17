/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package svc

import (
	"imooc.com/easy-chat/apps/im/immodels"
	"imooc.com/easy-chat/apps/im/ws/internal/config"
	"imooc.com/easy-chat/apps/task/mq/mqclient"
)

type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel
	mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		MsgChatTransferClient: mqclient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		ChatLogModel:          immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
