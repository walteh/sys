package mount

// Constants for mount(2).
const (
	MS_RDONLY      = 0x1
	MS_NOSUID      = 0x2
	MS_NODEV       = 0x4
	MS_NOEXEC      = 0x8
	MS_SYNCHRONOUS = 0x10
	MS_REMOUNT     = 0x20
	MS_MANDLOCK    = 0x40
	MS_DIRSYNC     = 0x80
	MS_NOATIME     = 0x400
	MS_NODIRATIME  = 0x800
	MS_BIND        = 0x1000
	MS_MOVE        = 0x2000
	MS_REC         = 0x4000

	MS_POSIXACL    = 0x10000
	MS_UNBINDABLE  = 0x20000
	MS_PRIVATE     = 0x40000
	MS_SLAVE       = 0x80000
	MS_SHARED      = 0x100000
	MS_RELATIME    = 0x200000
	MS_KERNMOUNT   = 0x400000
	MS_I_VERSION   = 0x800000
	MS_STRICTATIME = 0x1000000

	MS_MGC_VAL = 0xC0ED0000
	MS_MGC_MSK = 0xffff0000
)

const (
	// RDONLY will mount the file system read-only.
	RDONLY = MS_RDONLY

	// NOSUID will not allow set-user-identifier or set-group-identifier bits to
	// take effect.
	NOSUID = MS_NOSUID

	// NODEV will not interpret character or block special devices on the file
	// system.
	NODEV = MS_NODEV

	// NOEXEC will not allow execution of any binaries on the mounted file system.
	NOEXEC = MS_NOEXEC

	// SYNCHRONOUS will allow I/O to the file system to be done synchronously.
	SYNCHRONOUS = MS_SYNCHRONOUS

	// DIRSYNC will force all directory updates within the file system to be done
	// synchronously. This affects the following system calls: create, link,
	// unlink, symlink, mkdir, rmdir, mknod and rename.
	DIRSYNC = MS_DIRSYNC

	// REMOUNT will attempt to remount an already-mounted file system. This is
	// commonly used to change the mount flags for a file system, especially to
	// make a readonly file system writeable. It does not change device or mount
	// point.
	REMOUNT = MS_REMOUNT

	// MANDLOCK will force mandatory locks on a filesystem.
	MANDLOCK = MS_MANDLOCK

	// NOATIME will not update the file access time when reading from a file.
	NOATIME = MS_NOATIME

	// NODIRATIME will not update the directory access time.
	NODIRATIME = MS_NODIRATIME

	// BIND remounts a subtree somewhere else.
	BIND = MS_BIND

	// RBIND remounts a subtree and all possible submounts somewhere else.
	RBIND = MS_BIND | MS_REC

	// UNBINDABLE creates a mount which cannot be cloned through a bind operation.
	UNBINDABLE = MS_UNBINDABLE

	// RUNBINDABLE marks the entire mount tree as UNBINDABLE.
	RUNBINDABLE = MS_UNBINDABLE | MS_REC

	// PRIVATE creates a mount which carries no propagation abilities.
	PRIVATE = MS_PRIVATE

	// RPRIVATE marks the entire mount tree as PRIVATE.
	RPRIVATE = MS_PRIVATE | MS_REC

	// SLAVE creates a mount which receives propagation from its master, but not
	// vice versa.
	SLAVE = MS_SLAVE

	// RSLAVE marks the entire mount tree as SLAVE.
	RSLAVE = MS_SLAVE | MS_REC

	// SHARED creates a mount which provides the ability to create mirrors of
	// that mount such that mounts and unmounts within any of the mirrors
	// propagate to the other mirrors.
	SHARED = MS_SHARED

	// RSHARED marks the entire mount tree as SHARED.
	RSHARED = MS_SHARED | MS_REC

	// RELATIME updates inode access times relative to modify or change time.
	RELATIME = MS_RELATIME

	// STRICTATIME allows to explicitly request full atime updates.  This makes
	// it possible for the kernel to default to relatime or noatime but still
	// allow userspace to override it.
	STRICTATIME = MS_STRICTATIME

	// mntDetach = MNT_DETACH
)
