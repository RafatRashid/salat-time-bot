package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

func LogInfo(args ...any) {
	message := coalesceArguments(args, make([]string, 0))
	slog.Info(strings.Join(message, " "))
}

func LogError(args ...any) {
	message := coalesceArguments(args, make([]string, 0))
	slog.Error(strings.Join(message, " "))
}

func coalesceArguments(args []any, message []string) []string {
	for _, arg := range args {
		switch inferredVal := arg.(type) {
		case string:
			message = append(message, inferredVal)

		default:
			jsonArg := ToJson(inferredVal)
			message = append(message, jsonArg)
		}
	}
	return message
}

func ToJson(payload any) string {
	bytes, err := json.Marshal(payload)
	if err != nil {
		slog.Error(fmt.Sprintf("error on marshalling: %s", err.Error()))
	}
	return string(bytes)
}

func RecoverPanic() {
	if r := recover(); r != nil {
		slog.Error(fmt.Sprintf("%v", r))
	}
}

func ToPtr[T any](v T) *T {
	s := &v
	return s
}

func ConvertToType[T any](objectLiteral any) T {
	jsonStream, _ := json.Marshal(objectLiteral)
	var unmarshalled T
	if err := json.Unmarshal(jsonStream, &unmarshalled); err != nil {
		LogError("error on manual type conversion: ", err.Error())
	}
	return unmarshalled
}
