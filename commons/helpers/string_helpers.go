package helpers

import "strings"

func SeparaComando(comandoCompleto string) (comandoSeparado []string) {
	comandoSeparado = strings.Split(strings.TrimSuffix(comandoCompleto, "\n"), " ")
	return comandoSeparado
}
