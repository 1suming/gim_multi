package room

import (
	"fmt"
	"gim/internal/logic/domain/room/repo"
	"gim/pkg/protocol/pb"
	"gim/pkg/util"
	"testing"
	"time"
)

func Test_service_DelExpireMessage(t *testing.T) {
	err := Service.DelExpireMessage(1)
	fmt.Println(err)
}

func Test_service_List(t *testing.T) {
	msgs, err := repo.MessageRepo.List(1, 1)
	fmt.Println(err)
	fmt.Println(msgs)
}

func Test_service_AddMessage(t *testing.T) {
	for i := 1; i <= 20; i++ {
		err := Service.AddMessage(1, &pb.Message{
			Seq:      int64(i),
			SendTime: util.UnixMilliTime(time.Now()),
		})
		fmt.Println(i, err)
		time.Sleep(time.Second)
	}

	err := Service.DelExpireMessage(1)
	fmt.Println(err)
}
