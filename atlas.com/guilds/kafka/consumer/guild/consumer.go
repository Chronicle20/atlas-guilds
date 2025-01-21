package guild

import (
	"atlas-guilds/guild"
	consumer2 "atlas-guilds/kafka/consumer"
	"context"
	"github.com/Chronicle20/atlas-kafka/consumer"
	"github.com/Chronicle20/atlas-kafka/handler"
	"github.com/Chronicle20/atlas-kafka/message"
	"github.com/Chronicle20/atlas-kafka/topic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const consumerCommand = "guild_command"

func CommandConsumer(l logrus.FieldLogger) func(groupId string) consumer.Config {
	return func(groupId string) consumer.Config {
		return consumer2.NewConfig(l)(consumerCommand)(EnvCommandTopic)(groupId)
	}
}

func RequestCreateRegister(db *gorm.DB) func(l logrus.FieldLogger) (string, handler.Handler) {
	return func(l logrus.FieldLogger) (string, handler.Handler) {
		t, _ := topic.EnvProvider(l)(EnvCommandTopic)()
		return t, message.AdaptHandler(message.PersistentConfig(handleRequestCreateCommand(db)))
	}
}

func handleRequestCreateCommand(db *gorm.DB) func(l logrus.FieldLogger, ctx context.Context, c command[requestCreateBody]) {
	return func(l logrus.FieldLogger, ctx context.Context, c command[requestCreateBody]) {
		if c.Type != CommandTypeRequestCreate {
			return
		}

		_ = guild.RequestCreate(l)(ctx)(db)(c.CharacterId, c.Body.WorldId, c.Body.ChannelId, c.Body.MapId, c.Body.Name)
	}
}
