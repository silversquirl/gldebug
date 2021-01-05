package gldebug

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

// MessageCallback outputs an OpenGL debug message.
// It is designed for use with gl.DebugMessageCallback.
func MessageCallback(source, gltype, id, severity uint32, length int32, message string, userParam unsafe.Pointer) {
	var sourceStr, typeStr, severityStr string

	switch source {
	case gl.DEBUG_SOURCE_API:
		sourceStr = "API"
	case gl.DEBUG_SOURCE_WINDOW_SYSTEM:
		sourceStr = "Window System"
	case gl.DEBUG_SOURCE_SHADER_COMPILER:
		sourceStr = "Shader Compiler"
	case gl.DEBUG_SOURCE_THIRD_PARTY:
		sourceStr = "Third Party"
	case gl.DEBUG_SOURCE_APPLICATION:
		sourceStr = "Application"
	case gl.DEBUG_SOURCE_OTHER:
		sourceStr = "Other"
	}

	switch gltype {
	case gl.DEBUG_TYPE_ERROR:
		typeStr = "Error"
	case gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR:
		typeStr = "Deprecated Behaviour"
	case gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR:
		typeStr = "Undefined Behaviour"
	case gl.DEBUG_TYPE_PORTABILITY:
		typeStr = "Portability"
	case gl.DEBUG_TYPE_PERFORMANCE:
		typeStr = "Performance"
	case gl.DEBUG_TYPE_MARKER:
		typeStr = "Marker"
	case gl.DEBUG_TYPE_PUSH_GROUP:
		typeStr = "Push Group"
	case gl.DEBUG_TYPE_POP_GROUP:
		typeStr = "Pop Group"
	case gl.DEBUG_TYPE_OTHER:
		typeStr = "Other"
	}

	switch severity {
	case gl.DEBUG_SEVERITY_HIGH:
		severityStr = "High"
	case gl.DEBUG_SEVERITY_MEDIUM:
		severityStr = "Medium"
	case gl.DEBUG_SEVERITY_LOW:
		severityStr = "Low"
	case gl.DEBUG_SEVERITY_NOTIFICATION:
		severityStr = "Notification"
	}

	msg := fmt.Sprintf("(%d) source: %s, type: %s, severity: %s, message: %s", id, sourceStr, typeStr, severityStr, message)
	if severity == gl.DEBUG_SEVERITY_HIGH {
		panic(msg)
	} else {
		log.Println("GL debug:", msg)
	}
}
