package snowflake

import (
	"time"

	"campus-recruiting-api/configs"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init() (err error) {
	cfg := configs.Get().System
	var st time.Time
	st, err = time.Parse("2006-01-02", cfg.StartTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(cfg.MachineID)
	return err
}

func GenID() int64 {
	return node.Generate().Int64()
}
