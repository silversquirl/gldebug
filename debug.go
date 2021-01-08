package gldebug

import (
	"fmt"
	"log"

	"github.com/vktec/gll"
)

// MessageCallback outputs an OpenGL debug message.
// It is designed for use with glDebugMessageCallback.
func MessageCallback(source, type_, id, severity uint32, length int32, message string) {
	var sourceStr, typeStr, severityStr string

	switch source {
	case gll.DEBUG_SOURCE_API:
		sourceStr = "API"
	case gll.DEBUG_SOURCE_WINDOW_SYSTEM:
		sourceStr = "Window System"
	case gll.DEBUG_SOURCE_SHADER_COMPILER:
		sourceStr = "Shader Compiler"
	case gll.DEBUG_SOURCE_THIRD_PARTY:
		sourceStr = "Third Party"
	case gll.DEBUG_SOURCE_APPLICATION:
		sourceStr = "Application"
	case gll.DEBUG_SOURCE_OTHER:
		sourceStr = "Other"
	}

	switch type_ {
	case gll.DEBUG_TYPE_ERROR:
		typeStr = "Error"
	case gll.DEBUG_TYPE_DEPRECATED_BEHAVIOR:
		typeStr = "Deprecated Behaviour"
	case gll.DEBUG_TYPE_UNDEFINED_BEHAVIOR:
		typeStr = "Undefined Behaviour"
	case gll.DEBUG_TYPE_PORTABILITY:
		typeStr = "Portability"
	case gll.DEBUG_TYPE_PERFORMANCE:
		typeStr = "Performance"
	case gll.DEBUG_TYPE_MARKER:
		typeStr = "Marker"
	case gll.DEBUG_TYPE_PUSH_GROUP:
		typeStr = "Push Group"
	case gll.DEBUG_TYPE_POP_GROUP:
		typeStr = "Pop Group"
	case gll.DEBUG_TYPE_OTHER:
		typeStr = "Other"
	}

	switch severity {
	case gll.DEBUG_SEVERITY_HIGH:
		severityStr = "High"
	case gll.DEBUG_SEVERITY_MEDIUM:
		severityStr = "Medium"
	case gll.DEBUG_SEVERITY_LOW:
		severityStr = "Low"
	case gll.DEBUG_SEVERITY_NOTIFICATION:
		severityStr = "Notification"
	}

	msg := fmt.Sprintf("(%d) source: %s, type: %s, severity: %s, message: %s", id, sourceStr, typeStr, severityStr, message)
	if severity == gll.DEBUG_SEVERITY_HIGH {
		panic(msg)
	} else {
		log.Println("GL debug:", msg)
	}
}
