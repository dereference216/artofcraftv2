package utils

import (
	"strings"
	"time"

	"github.com/gookit/color"
)

func PrintError(header, msg string) {
	color.Printf("<fg=F13030>[</><fg=CE76BE>"+time.Now().Format("15:04:05")+"</><fg=F13030>]</> <fg=F13030>[</><fg=CE3333>ERROR %s</><fg=F13030>]</> <fg=BCB3B1>-></> <fg=CB4141>%s</>\n", header, msg)
}
func (conf *Config) PrintWarn(header, msg string) {
	if conf.Debug {
		color.Printf("<fg=F13030>[</><fg=CE76BE>"+time.Now().Format("15:04:05")+"</><fg=F13030>]</> <fg=F13030>[</><fg=CE3333>WARN %s</><fg=F13030>]</> <fg=BCB3B1>-></> <fg=CB4141>%s</>\n", header, msg)
	}
}

func (conf *Config) PrintDebug(header, msg string) {
	if conf.Debug {
		color.Printf("<fg=F3D62E>[</><fg=CE76BE>"+time.Now().Format("15:04:05")+"</><fg=F3D62E>]</> <fg=F3D62E>[</><fg=E3F00B>DEBUG %s</><fg=F3D62E>]</> <fg=BCB3B1>-></> <fg=BCB3B1>%s</>\n", header, msg)
	}
}

func (conf *Config) PrintGen(header, msg string) {
	color.Printf("<fg=F3D62E>[</><fg=CE76BE>"+time.Now().Format("15:04:05")+"</><fg=F3D62E>]</> <fg=2EF331>[</><fg=2EF331>%s</><fg=2EF331>]</> <fg=BCB3B1>-></> <fg=F3D62E>%s</>\n", header, msg)
}

func (conf *Config) PrintNetworkError(stat int, header, msg string) {
	color.Printf("<fg=F13030>[</><fg=CE76BE>"+time.Now().Format("15:04:05")+"</><fg=F13030>]</> <fg=F13030>[</><fg=CE3333>%d</><fg=F13030>]</> <fg=F13030>[</><fg=CE3333>%s</><fg=F13030>]</> <fg=BCB3B1>-></> <fg=CB4141>%s</>\n", stat, header, strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(msg, "\t", ""), "\n", ""), " ", ""))
}
