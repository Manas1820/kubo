//go:build (linux || darwin || freebsd || netbsd || openbsd) && !nofuse
// +build linux darwin freebsd netbsd openbsd
// +build !nofuse

package mfs

import (
	core "github.com/ipfs/kubo/core"
	mount "github.com/ipfs/kubo/fuse/mount"
)

// Mount mounts MFS at a given location, and returns a mount.Mount instance.
func Mount(ipfs *core.IpfsNode, mountpoint string) (mount.Mount, error) {
	cfg, err := ipfs.Repo.Config()
	if err != nil {
		return nil, err
	}
	allowOther := cfg.Mounts.FuseAllowOther
	fsys := NewFileSystem(ipfs)
	return mount.NewMount(ipfs.Process, fsys, mountpoint, allowOther)
}
