package libvirt

import (
	"encoding/hex"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

func (l Libvirt) lookup(vmId uuid.UUID) (dom libvirt.Domain, err error) {
	vmIdHex, _ := hex.DecodeString(vmId.String())
	var libvirtUUID libvirt.UUID
	copy(libvirtUUID[:], vmIdHex[:])
	dom, err = l.conn.DomainLookupByUUID(libvirtUUID)
	return
}

func (l Libvirt) GetVMState(dom libvirt.Domain) (state int32, reason int32, err error) {
	state, reason, err = l.conn.DomainGetState(dom, 0)
	return
}