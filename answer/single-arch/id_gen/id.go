//专门生成唯一id
package id_gen

import (
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}
func Init(machineId uint16) {
	sonyMachineID = machineId
	settings := sonyflake.Settings{}
	settings.MachineID = getMachineID //回调函数，在初始化时会调用该函数来获取workerID
	sonyFlake = sonyflake.NewSonyflake(settings)
}

func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = errors.New("Must init first.")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
