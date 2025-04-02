package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	// User credentials
	// These should be replaced with actual values or securely managed
	username = "testuser"
	password = "Supersecurepassword1234."
	group    = "Administrators"
	// Debug flag
	// Set to true for viewing debug messages
	// Set to false for production or normal operation
	dbg = true
)

var (
	// Windows API functions
	modNetapi32                 = windows.NewLazySystemDLL("netapi32.dll")
	procNetUserAdd              = modNetapi32.NewProc("NetUserAdd")
	procNetLocalGroupAddMembers = modNetapi32.NewProc("NetLocalGroupAddMembers")
)

const (
	USER_PRIV_USER        = 1
	UF_SCRIPT             = 1
	UF_NORMAL_ACCOUNT     = 0x0200
	ERROR_SUCCESS         = 0
	ERROR_MEMBER_IN_ALIAS = 1378
)

type USER_INFO_1 struct {
	Name        *uint16
	Password    *uint16
	PasswordAge uint32 // Este campo podría evitar desalineación
	Priv        uint32
	HomeDir     *uint16
	Comment     *uint16
	Flags       uint32
	ScriptPath  *uint16
}

type LOCALGROUP_MEMBERS_INFO_3 struct {
	Lgrmi3DomainAndName *uint16
}

func main() {
	var err error

	err = createUser(username, password)
	if err != nil {
		debug(fmt.Sprintf("Error al crear el usuario: %v\n", err))
		return
	}
	debug(fmt.Sprintf("Usuario '%s' creado exitosamente.\n", username))
	err = addUserToGroup(username, group)
	if err != nil {
		debug(fmt.Sprintf("Error al agregar el usuario al grupo %s: %v\n", group, err))
		return
	}
	debug(fmt.Sprintf("Usuario '%s' agregado al grupo %s exitosamente.\n", username, group))
}

func createUser(username, password string) error {
	var parmErr uint32

	userInfo := &USER_INFO_1{
		Name:       windows.StringToUTF16Ptr(username),
		Password:   windows.StringToUTF16Ptr(password),
		Priv:       USER_PRIV_USER,
		Flags:      UF_NORMAL_ACCOUNT,
		HomeDir:    nil,
		Comment:    nil,
		ScriptPath: nil,
	}

	ret, _, _ := procNetUserAdd.Call(
		0,
		1,
		uintptr(unsafe.Pointer(userInfo)),
		uintptr(unsafe.Pointer(&parmErr)),
	)
	if ret != ERROR_SUCCESS {
		return fmt.Errorf("NetUserAdd failed with error code %d (parmErr: %d)", ret, parmErr)
	}

	return nil
}

func addUserToGroup(username, groupName string) error {
	groupNameUTF16 := windows.StringToUTF16Ptr(groupName)
	usernameUTF16 := windows.StringToUTF16Ptr(username)

	memberInfo := LOCALGROUP_MEMBERS_INFO_3{
		Lgrmi3DomainAndName: usernameUTF16,
	}

	ret, _, _ := procNetLocalGroupAddMembers.Call(
		0,
		uintptr(unsafe.Pointer(groupNameUTF16)),
		3,
		uintptr(unsafe.Pointer(&memberInfo)),
		1,
	)
	if ret != ERROR_SUCCESS && ret != ERROR_MEMBER_IN_ALIAS {
		return fmt.Errorf("NetLocalGroupAddMembers failed with error code %d", ret)
	}

	return nil
}

func debug(msg string) {
	// This function is a placeholder for debugging purposes.
	// You can implement logging or debugging as needed.
	if dbg {
		fmt.Println(msg)
	}
}
