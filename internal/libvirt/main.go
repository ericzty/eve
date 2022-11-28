package libvirt

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/ericzty/eve/internal/controllers"
)

func InitHVs(HV *controllers.HV) (err error) {
	c, err := net.DialTimeout("tcp", HV.IP.String()+":"+strconv.Itoa(HV.Port), 3*time.Second)

	if err != nil {
		HV.Status = "Offline"
		return fmt.Errorf("Failed to dial to libvirt tcp socket: %v", err)
	}

	l := libvirt.New(c)

	if err := l.Connect(); err != nil {
		HV.Status = "Offline"
		return fmt.Errorf("Failed to communicate with libvirt: %v", err)
	}

	v, err := l.Version()

	if err != nil {
		HV.Status = "Offline"
		return fmt.Errorf("Failed to get libvirt version: %v", err)
	}

	defer l.Disconnect()
	defer c.Close()

	HV.Status = "Online"
	HV.Version = v

	return
}
