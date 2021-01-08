// +build renderdoc

// Renderdoc API wrapper
package gldebug

/*
#cgo LDFLAGS: -ldl
#include <stddef.h>
#include <dlfcn.h>
#include "renderdoc.h"
RENDERDOC_API_1_4_1 *rdocInit(void) {
	void *rdoc = dlopen("librenderdoc.so", RTLD_NOW|RTLD_NOLOAD);
	if (!rdoc) return NULL;
	pRENDERDOC_GetAPI getAPI = (pRENDERDOC_GetAPI)dlsym(rdoc, "RENDERDOC_GetAPI");
	dlclose(rdoc);
	void *api;
	if (!getAPI(eRENDERDOC_API_Version_1_4_1, &api)) return NULL;
	return api;
}
void rdocStartCapture(RENDERDOC_API_1_4_1 *api) {
	api->StartFrameCapture(NULL, NULL);
}
void rdocEndCapture(RENDERDOC_API_1_4_1 *api) {
	api->EndFrameCapture(NULL, NULL);
}
*/
import "C"

var rdocAPI = C.rdocInit()

// Rdoc returns true if the program is running inside RenderDoc, false otherwise.
func Rdoc() bool {
	return rdocAPI != nil
}

// RdocStartCapture starts a capture if the program is running inside RenderDoc.
func RdocStartCapture() {
	if Rdoc() {
		C.rdocStartCapture(rdocAPI)
	}
}

// RdocStartCapture ends the current capture if the program is running inside RenderDoc.
func RdocEndCapture() {
	if Rdoc() {
		C.rdocEndCapture(rdocAPI)
	}
}
