// Auto-generated by avdl-compiler v1.3.11 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/simple_fs.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type OpID [16]byte
type PathType int

const (
	PathType_LOCAL PathType = 0
	PathType_KBFS  PathType = 1
)

var PathTypeMap = map[string]PathType{
	"LOCAL": 0,
	"KBFS":  1,
}

var PathTypeRevMap = map[PathType]string{
	0: "LOCAL",
	1: "KBFS",
}

func (e PathType) String() string {
	if v, ok := PathTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Path struct {
	PathType__ PathType `codec:"PathType" json:"PathType"`
	Local__    *string  `codec:"local,omitempty" json:"local,omitempty"`
	Kbfs__     *string  `codec:"kbfs,omitempty" json:"kbfs,omitempty"`
}

func (o *Path) PathType() (ret PathType, err error) {
	switch o.PathType__ {
	case PathType_LOCAL:
		if o.Local__ == nil {
			err = errors.New("unexpected nil value for Local__")
			return ret, err
		}
	case PathType_KBFS:
		if o.Kbfs__ == nil {
			err = errors.New("unexpected nil value for Kbfs__")
			return ret, err
		}
	}
	return o.PathType__, nil
}

func (o Path) Local() string {
	if o.PathType__ != PathType_LOCAL {
		panic("wrong case accessed")
	}
	if o.Local__ == nil {
		return ""
	}
	return *o.Local__
}

func (o Path) Kbfs() string {
	if o.PathType__ != PathType_KBFS {
		panic("wrong case accessed")
	}
	if o.Kbfs__ == nil {
		return ""
	}
	return *o.Kbfs__
}

func NewPathWithLocal(v string) Path {
	return Path{
		PathType__: PathType_LOCAL,
		Local__:    &v,
	}
}

func NewPathWithKbfs(v string) Path {
	return Path{
		PathType__: PathType_KBFS,
		Kbfs__:     &v,
	}
}

type DirentType int

const (
	DirentType_FILE DirentType = 0
	DirentType_DIR  DirentType = 1
	DirentType_SYM  DirentType = 2
	DirentType_EXEC DirentType = 3
)

var DirentTypeMap = map[string]DirentType{
	"FILE": 0,
	"DIR":  1,
	"SYM":  2,
	"EXEC": 3,
}

var DirentTypeRevMap = map[DirentType]string{
	0: "FILE",
	1: "DIR",
	2: "SYM",
	3: "EXEC",
}

func (e DirentType) String() string {
	if v, ok := DirentTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Dirent struct {
	Time       Time       `codec:"time" json:"time"`
	Size       int        `codec:"size" json:"size"`
	Name       string     `codec:"name" json:"name"`
	DirentType DirentType `codec:"direntType" json:"direntType"`
}

type ErrorNum int
type OpenFlags int

const (
	OpenFlags_READ      OpenFlags = 0
	OpenFlags_REPLACE   OpenFlags = 1
	OpenFlags_EXISTING  OpenFlags = 2
	OpenFlags_WRITE     OpenFlags = 4
	OpenFlags_APPEND    OpenFlags = 8
	OpenFlags_DIRECTORY OpenFlags = 16
)

var OpenFlagsMap = map[string]OpenFlags{
	"READ":      0,
	"REPLACE":   1,
	"EXISTING":  2,
	"WRITE":     4,
	"APPEND":    8,
	"DIRECTORY": 16,
}

var OpenFlagsRevMap = map[OpenFlags]string{
	0:  "READ",
	1:  "REPLACE",
	2:  "EXISTING",
	4:  "WRITE",
	8:  "APPEND",
	16: "DIRECTORY",
}

func (e OpenFlags) String() string {
	if v, ok := OpenFlagsRevMap[e]; ok {
		return v
	}
	return ""
}

type Progress int
type SimpleFSListResult struct {
	Entries  []Dirent `codec:"entries" json:"entries"`
	Progress Progress `codec:"progress" json:"progress"`
}

type FileContent struct {
	Data     []byte   `codec:"data" json:"data"`
	Progress Progress `codec:"progress" json:"progress"`
}

type AsyncOps int

const (
	AsyncOps_LIST           AsyncOps = 0
	AsyncOps_LIST_RECURSIVE AsyncOps = 1
	AsyncOps_READ           AsyncOps = 2
	AsyncOps_WRITE          AsyncOps = 3
	AsyncOps_COPY           AsyncOps = 4
	AsyncOps_MOVE           AsyncOps = 5
	AsyncOps_REMOVE         AsyncOps = 6
)

var AsyncOpsMap = map[string]AsyncOps{
	"LIST":           0,
	"LIST_RECURSIVE": 1,
	"READ":           2,
	"WRITE":          3,
	"COPY":           4,
	"MOVE":           5,
	"REMOVE":         6,
}

var AsyncOpsRevMap = map[AsyncOps]string{
	0: "LIST",
	1: "LIST_RECURSIVE",
	2: "READ",
	3: "WRITE",
	4: "COPY",
	5: "MOVE",
	6: "REMOVE",
}

func (e AsyncOps) String() string {
	if v, ok := AsyncOpsRevMap[e]; ok {
		return v
	}
	return ""
}

type ListArgs struct {
	OpID OpID `codec:"opID" json:"opID"`
	Path Path `codec:"path" json:"path"`
}

type RemoveArgs struct {
	OpID OpID `codec:"opID" json:"opID"`
	Path Path `codec:"path" json:"path"`
}

type ReadArgs struct {
	OpID   OpID  `codec:"opID" json:"opID"`
	Path   Path  `codec:"path" json:"path"`
	Offset int64 `codec:"offset" json:"offset"`
	Size   int   `codec:"size" json:"size"`
}

type WriteArgs struct {
	OpID   OpID  `codec:"opID" json:"opID"`
	Path   Path  `codec:"path" json:"path"`
	Offset int64 `codec:"offset" json:"offset"`
}

type CopyArgs struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type MoveArgs struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type OpDescription struct {
	AsyncOp__       AsyncOps    `codec:"asyncOp" json:"asyncOp"`
	List__          *ListArgs   `codec:"list,omitempty" json:"list,omitempty"`
	ListRecursive__ *ListArgs   `codec:"listRecursive,omitempty" json:"listRecursive,omitempty"`
	Read__          *ReadArgs   `codec:"read,omitempty" json:"read,omitempty"`
	Write__         *WriteArgs  `codec:"write,omitempty" json:"write,omitempty"`
	Copy__          *CopyArgs   `codec:"copy,omitempty" json:"copy,omitempty"`
	Move__          *MoveArgs   `codec:"move,omitempty" json:"move,omitempty"`
	Remove__        *RemoveArgs `codec:"remove,omitempty" json:"remove,omitempty"`
}

func (o *OpDescription) AsyncOp() (ret AsyncOps, err error) {
	switch o.AsyncOp__ {
	case AsyncOps_LIST:
		if o.List__ == nil {
			err = errors.New("unexpected nil value for List__")
			return ret, err
		}
	case AsyncOps_LIST_RECURSIVE:
		if o.ListRecursive__ == nil {
			err = errors.New("unexpected nil value for ListRecursive__")
			return ret, err
		}
	case AsyncOps_READ:
		if o.Read__ == nil {
			err = errors.New("unexpected nil value for Read__")
			return ret, err
		}
	case AsyncOps_WRITE:
		if o.Write__ == nil {
			err = errors.New("unexpected nil value for Write__")
			return ret, err
		}
	case AsyncOps_COPY:
		if o.Copy__ == nil {
			err = errors.New("unexpected nil value for Copy__")
			return ret, err
		}
	case AsyncOps_MOVE:
		if o.Move__ == nil {
			err = errors.New("unexpected nil value for Move__")
			return ret, err
		}
	case AsyncOps_REMOVE:
		if o.Remove__ == nil {
			err = errors.New("unexpected nil value for Remove__")
			return ret, err
		}
	}
	return o.AsyncOp__, nil
}

func (o OpDescription) List() ListArgs {
	if o.AsyncOp__ != AsyncOps_LIST {
		panic("wrong case accessed")
	}
	if o.List__ == nil {
		return ListArgs{}
	}
	return *o.List__
}

func (o OpDescription) ListRecursive() ListArgs {
	if o.AsyncOp__ != AsyncOps_LIST_RECURSIVE {
		panic("wrong case accessed")
	}
	if o.ListRecursive__ == nil {
		return ListArgs{}
	}
	return *o.ListRecursive__
}

func (o OpDescription) Read() ReadArgs {
	if o.AsyncOp__ != AsyncOps_READ {
		panic("wrong case accessed")
	}
	if o.Read__ == nil {
		return ReadArgs{}
	}
	return *o.Read__
}

func (o OpDescription) Write() WriteArgs {
	if o.AsyncOp__ != AsyncOps_WRITE {
		panic("wrong case accessed")
	}
	if o.Write__ == nil {
		return WriteArgs{}
	}
	return *o.Write__
}

func (o OpDescription) Copy() CopyArgs {
	if o.AsyncOp__ != AsyncOps_COPY {
		panic("wrong case accessed")
	}
	if o.Copy__ == nil {
		return CopyArgs{}
	}
	return *o.Copy__
}

func (o OpDescription) Move() MoveArgs {
	if o.AsyncOp__ != AsyncOps_MOVE {
		panic("wrong case accessed")
	}
	if o.Move__ == nil {
		return MoveArgs{}
	}
	return *o.Move__
}

func (o OpDescription) Remove() RemoveArgs {
	if o.AsyncOp__ != AsyncOps_REMOVE {
		panic("wrong case accessed")
	}
	if o.Remove__ == nil {
		return RemoveArgs{}
	}
	return *o.Remove__
}

func NewOpDescriptionWithList(v ListArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_LIST,
		List__:    &v,
	}
}

func NewOpDescriptionWithListRecursive(v ListArgs) OpDescription {
	return OpDescription{
		AsyncOp__:       AsyncOps_LIST_RECURSIVE,
		ListRecursive__: &v,
	}
}

func NewOpDescriptionWithRead(v ReadArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_READ,
		Read__:    &v,
	}
}

func NewOpDescriptionWithWrite(v WriteArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_WRITE,
		Write__:   &v,
	}
}

func NewOpDescriptionWithCopy(v CopyArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_COPY,
		Copy__:    &v,
	}
}

func NewOpDescriptionWithMove(v MoveArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_MOVE,
		Move__:    &v,
	}
}

func NewOpDescriptionWithRemove(v RemoveArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_REMOVE,
		Remove__:  &v,
	}
}

type SimpleFSListArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Path Path `codec:"path" json:"path"`
}

type SimpleFSListRecursiveArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Path Path `codec:"path" json:"path"`
}

type SimpleFSReadListArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSCopyArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSCopyRecursiveArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSMoveArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSRenameArg struct {
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSOpenArg struct {
	OpID  OpID      `codec:"opID" json:"opID"`
	Dest  Path      `codec:"dest" json:"dest"`
	Flags OpenFlags `codec:"flags" json:"flags"`
}

type SimpleFSSetStatArg struct {
	Dest Path       `codec:"dest" json:"dest"`
	Flag DirentType `codec:"flag" json:"flag"`
}

type SimpleFSReadArg struct {
	OpID   OpID  `codec:"opID" json:"opID"`
	Offset int64 `codec:"offset" json:"offset"`
	Size   int   `codec:"size" json:"size"`
}

type SimpleFSWriteArg struct {
	OpID    OpID   `codec:"opID" json:"opID"`
	Offset  int64  `codec:"offset" json:"offset"`
	Content []byte `codec:"content" json:"content"`
}

type SimpleFSRemoveArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Path Path `codec:"path" json:"path"`
}

type SimpleFSStatArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSMakeOpidArg struct {
}

type SimpleFSCloseArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSCancelArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSCheckArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSGetOpsArg struct {
}

type SimpleFSWaitArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSInterface interface {
	// Begin list of items in directory at path
	// Retrieve results with readList()
	// Can be a single file to get flags/status
	SimpleFSList(context.Context, SimpleFSListArg) error
	// Begin recursive list of items in directory at path
	SimpleFSListRecursive(context.Context, SimpleFSListRecursiveArg) error
	// Get list of Paths in progress. Can indicate status of pending
	// to get more entries.
	SimpleFSReadList(context.Context, OpID) (SimpleFSListResult, error)
	// Begin copy of file or directory
	SimpleFSCopy(context.Context, SimpleFSCopyArg) error
	// Begin recursive copy of directory
	SimpleFSCopyRecursive(context.Context, SimpleFSCopyRecursiveArg) error
	// Begin move of file or directory, from/to KBFS only
	SimpleFSMove(context.Context, SimpleFSMoveArg) error
	// Rename file or directory, KBFS side only
	SimpleFSRename(context.Context, SimpleFSRenameArg) error
	// Create/open a file and leave it open
	// or create a directory
	// Files must be closed afterwards.
	SimpleFSOpen(context.Context, SimpleFSOpenArg) error
	// Set/clear file bits - only executable for now
	SimpleFSSetStat(context.Context, SimpleFSSetStatArg) error
	// Read (possibly partial) contents of open file,
	// up to the amount specified by size.
	// Repeat until zero bytes are returned or error.
	// If size is zero, read an arbitrary amount.
	SimpleFSRead(context.Context, SimpleFSReadArg) (FileContent, error)
	// Append content to opened file.
	// May be repeated until OpID is closed.
	SimpleFSWrite(context.Context, SimpleFSWriteArg) error
	// Remove file or directory from filesystem
	SimpleFSRemove(context.Context, SimpleFSRemoveArg) error
	// Get info about file
	SimpleFSStat(context.Context, Path) (Dirent, error)
	// Convenience helper for generating new random value
	SimpleFSMakeOpid(context.Context) (OpID, error)
	// Close OpID, cancels any pending operation.
	// Must be called after list/copy/remove
	SimpleFSClose(context.Context, OpID) error
	// Cancels a running operation, like copy.
	SimpleFSCancel(context.Context, OpID) error
	// Check progress of pending operation
	SimpleFSCheck(context.Context, OpID) (Progress, error)
	// Get all the outstanding operations
	SimpleFSGetOps(context.Context) ([]OpDescription, error)
	// Blocking wait for the pending operation to finish
	SimpleFSWait(context.Context, OpID) error
}

func SimpleFSProtocol(i SimpleFSInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.SimpleFS",
		Methods: map[string]rpc.ServeHandlerDescription{
			"simpleFSList": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSListArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSListArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSListArg)(nil), args)
						return
					}
					err = i.SimpleFSList(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSListRecursive": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSListRecursiveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSListRecursiveArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSListRecursiveArg)(nil), args)
						return
					}
					err = i.SimpleFSListRecursive(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSReadList": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSReadListArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSReadListArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSReadListArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSReadList(ctx, (*typedArgs)[0].OpID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSCopy": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSCopyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSCopyArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSCopyArg)(nil), args)
						return
					}
					err = i.SimpleFSCopy(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSCopyRecursive": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSCopyRecursiveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSCopyRecursiveArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSCopyRecursiveArg)(nil), args)
						return
					}
					err = i.SimpleFSCopyRecursive(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSMove": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSMoveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSMoveArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSMoveArg)(nil), args)
						return
					}
					err = i.SimpleFSMove(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSRename": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSRenameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSRenameArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSRenameArg)(nil), args)
						return
					}
					err = i.SimpleFSRename(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSOpen": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSOpenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSOpenArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSOpenArg)(nil), args)
						return
					}
					err = i.SimpleFSOpen(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSSetStat": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSSetStatArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSSetStatArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSSetStatArg)(nil), args)
						return
					}
					err = i.SimpleFSSetStat(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSRead": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSReadArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSReadArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSReadArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSRead(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSWrite": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSWriteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSWriteArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSWriteArg)(nil), args)
						return
					}
					err = i.SimpleFSWrite(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSRemove": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSRemoveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSRemoveArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSRemoveArg)(nil), args)
						return
					}
					err = i.SimpleFSRemove(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSStat": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSStatArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSStatArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSStatArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSStat(ctx, (*typedArgs)[0].Path)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSMakeOpid": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSMakeOpidArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSMakeOpid(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSClose": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSCloseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSCloseArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSCloseArg)(nil), args)
						return
					}
					err = i.SimpleFSClose(ctx, (*typedArgs)[0].OpID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSCancel": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSCancelArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSCancelArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSCancelArg)(nil), args)
						return
					}
					err = i.SimpleFSCancel(ctx, (*typedArgs)[0].OpID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSCheck": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSCheckArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSCheckArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSCheckArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSCheck(ctx, (*typedArgs)[0].OpID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSGetOps": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSGetOpsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSGetOps(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"simpleFSWait": {
				MakeArg: func() interface{} {
					ret := make([]SimpleFSWaitArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SimpleFSWaitArg)
					if !ok {
						err = rpc.NewTypeError((*[]SimpleFSWaitArg)(nil), args)
						return
					}
					err = i.SimpleFSWait(ctx, (*typedArgs)[0].OpID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type SimpleFSClient struct {
	Cli rpc.GenericClient
}

// Begin list of items in directory at path
// Retrieve results with readList()
// Can be a single file to get flags/status
func (c SimpleFSClient) SimpleFSList(ctx context.Context, __arg SimpleFSListArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSList", []interface{}{__arg}, nil)
	return
}

// Begin recursive list of items in directory at path
func (c SimpleFSClient) SimpleFSListRecursive(ctx context.Context, __arg SimpleFSListRecursiveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSListRecursive", []interface{}{__arg}, nil)
	return
}

// Get list of Paths in progress. Can indicate status of pending
// to get more entries.
func (c SimpleFSClient) SimpleFSReadList(ctx context.Context, opID OpID) (res SimpleFSListResult, err error) {
	__arg := SimpleFSReadListArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSReadList", []interface{}{__arg}, &res)
	return
}

// Begin copy of file or directory
func (c SimpleFSClient) SimpleFSCopy(ctx context.Context, __arg SimpleFSCopyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCopy", []interface{}{__arg}, nil)
	return
}

// Begin recursive copy of directory
func (c SimpleFSClient) SimpleFSCopyRecursive(ctx context.Context, __arg SimpleFSCopyRecursiveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCopyRecursive", []interface{}{__arg}, nil)
	return
}

// Begin move of file or directory, from/to KBFS only
func (c SimpleFSClient) SimpleFSMove(ctx context.Context, __arg SimpleFSMoveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSMove", []interface{}{__arg}, nil)
	return
}

// Rename file or directory, KBFS side only
func (c SimpleFSClient) SimpleFSRename(ctx context.Context, __arg SimpleFSRenameArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSRename", []interface{}{__arg}, nil)
	return
}

// Create/open a file and leave it open
// or create a directory
// Files must be closed afterwards.
func (c SimpleFSClient) SimpleFSOpen(ctx context.Context, __arg SimpleFSOpenArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSOpen", []interface{}{__arg}, nil)
	return
}

// Set/clear file bits - only executable for now
func (c SimpleFSClient) SimpleFSSetStat(ctx context.Context, __arg SimpleFSSetStatArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSetStat", []interface{}{__arg}, nil)
	return
}

// Read (possibly partial) contents of open file,
// up to the amount specified by size.
// Repeat until zero bytes are returned or error.
// If size is zero, read an arbitrary amount.
func (c SimpleFSClient) SimpleFSRead(ctx context.Context, __arg SimpleFSReadArg) (res FileContent, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSRead", []interface{}{__arg}, &res)
	return
}

// Append content to opened file.
// May be repeated until OpID is closed.
func (c SimpleFSClient) SimpleFSWrite(ctx context.Context, __arg SimpleFSWriteArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSWrite", []interface{}{__arg}, nil)
	return
}

// Remove file or directory from filesystem
func (c SimpleFSClient) SimpleFSRemove(ctx context.Context, __arg SimpleFSRemoveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSRemove", []interface{}{__arg}, nil)
	return
}

// Get info about file
func (c SimpleFSClient) SimpleFSStat(ctx context.Context, path Path) (res Dirent, err error) {
	__arg := SimpleFSStatArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSStat", []interface{}{__arg}, &res)
	return
}

// Convenience helper for generating new random value
func (c SimpleFSClient) SimpleFSMakeOpid(ctx context.Context) (res OpID, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSMakeOpid", []interface{}{SimpleFSMakeOpidArg{}}, &res)
	return
}

// Close OpID, cancels any pending operation.
// Must be called after list/copy/remove
func (c SimpleFSClient) SimpleFSClose(ctx context.Context, opID OpID) (err error) {
	__arg := SimpleFSCloseArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSClose", []interface{}{__arg}, nil)
	return
}

// Cancels a running operation, like copy.
func (c SimpleFSClient) SimpleFSCancel(ctx context.Context, opID OpID) (err error) {
	__arg := SimpleFSCancelArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCancel", []interface{}{__arg}, nil)
	return
}

// Check progress of pending operation
func (c SimpleFSClient) SimpleFSCheck(ctx context.Context, opID OpID) (res Progress, err error) {
	__arg := SimpleFSCheckArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCheck", []interface{}{__arg}, &res)
	return
}

// Get all the outstanding operations
func (c SimpleFSClient) SimpleFSGetOps(ctx context.Context) (res []OpDescription, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSGetOps", []interface{}{SimpleFSGetOpsArg{}}, &res)
	return
}

// Blocking wait for the pending operation to finish
func (c SimpleFSClient) SimpleFSWait(ctx context.Context, opID OpID) (err error) {
	__arg := SimpleFSWaitArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSWait", []interface{}{__arg}, nil)
	return
}
