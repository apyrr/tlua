package tspath

import (
	"slices"
	"strings"
)

// tlua sources are .tlua, declarations .d.tlua, and emit output .lua. The Go
// identifiers keep their upstream names so ported code stays greppable.
// (.tsx is deliberately unchanged for now.)
const (
	ExtensionTs          = ".tlua"
	ExtensionTsx         = ".tsx"
	ExtensionDts         = ".d.tlua"
	ExtensionJs          = ".lua"
	ExtensionJsx         = ".jsx"
	ExtensionJson        = ".json"
	ExtensionTsBuildInfo = ".tsbuildinfo"
)

var (
	SupportedDeclarationExtensions                 = []string{ExtensionDts}
	SupportedTSImplementationExtensions            = []string{ExtensionTs, ExtensionTsx}
	supportedTSExtensionsForExtractExtension       = []string{ExtensionDts, ExtensionTs, ExtensionTsx}
	AllSupportedExtensions                         = [][]string{{ExtensionTs, ExtensionTsx, ExtensionDts, ExtensionJs, ExtensionJsx}}
	SupportedTSExtensions                          = [][]string{{ExtensionTs, ExtensionTsx, ExtensionDts}}
	SupportedTSExtensionsFlat                      = []string{ExtensionTs, ExtensionTsx, ExtensionDts}
	SupportedJSExtensions                          = [][]string{{ExtensionJs, ExtensionJsx}}
	SupportedJSExtensionsFlat                      = []string{ExtensionJs, ExtensionJsx}
	AllSupportedExtensionsWithJson                 = slices.Concat(AllSupportedExtensions, [][]string{{ExtensionJson}})
	SupportedTSExtensionsWithJson                  = slices.Concat(SupportedTSExtensions, [][]string{{ExtensionJson}})
	SupportedTSExtensionsWithJsonFlat              = slices.Concat(SupportedTSExtensionsFlat, []string{ExtensionJson})
	ExtensionsNotSupportingExtensionlessResolution = []string{}
)

func ExtensionIsTs(ext string) bool {
	// The last clause matches custom declaration extensions (".d.<inner>.tlua").
	return ext == ExtensionTs || ext == ExtensionTsx || ext == ExtensionDts ||
		len(ext) > len(".d.")+len(ExtensionTs) && strings.HasPrefix(ext, ".d.") && strings.HasSuffix(ext, ExtensionTs)
}

var extensionsToRemove = []string{ExtensionDts, ExtensionTs, ExtensionJs, ExtensionTsx, ExtensionJsx, ExtensionJson}

func RemoveFileExtension(path string) string {
	// Remove any known extension even if it has more than one dot
	for _, ext := range extensionsToRemove {
		if strings.HasSuffix(path, ext) {
			return path[:len(path)-len(ext)]
		}
	}

	return path
}

func TryGetExtensionFromPath(p string) string {
	for _, ext := range extensionsToRemove {
		if FileExtensionIs(p, ext) {
			return ext
		}
	}
	return ""
}

func RemoveExtension(path string, extension string) string {
	return path[:len(path)-len(extension)]
}

func FileExtensionIsOneOf(path string, extensions []string) bool {
	for _, ext := range extensions {
		if FileExtensionIs(path, ext) {
			return true
		}
	}
	return false
}

func TryExtractTSExtension(fileName string) string {
	for _, ext := range supportedTSExtensionsForExtractExtension {
		if FileExtensionIs(fileName, ext) {
			return ext
		}
	}
	return ""
}

func HasTSFileExtension(path string) bool {
	return FileExtensionIsOneOf(path, SupportedTSExtensionsFlat)
}

func HasImplementationTSFileExtension(path string) bool {
	return FileExtensionIsOneOf(path, SupportedTSImplementationExtensions) && !IsDeclarationFileName(path)
}

func HasJSFileExtension(path string) bool {
	return FileExtensionIsOneOf(path, SupportedJSExtensionsFlat)
}

func HasJSONFileExtension(path string) bool {
	return FileExtensionIs(path, ExtensionJson)
}

func IsDeclarationFileName(fileName string) bool {
	return GetDeclarationFileExtension(fileName) != ""
}

func ExtensionIsOneOf(ext string, extensions []string) bool {
	return slices.Contains(extensions, ext)
}

func GetDeclarationFileExtension(fileName string) string {
	base := GetBaseFileName(fileName)
	for _, ext := range SupportedDeclarationExtensions {
		if strings.HasSuffix(base, ext) {
			return ext
		}
	}
	if strings.HasSuffix(base, ExtensionTs) {
		index := strings.Index(base, ".d.")
		if index >= 0 {
			return base[index:]
		}
	}
	return ""
}

func GetDeclarationEmitExtensionForPath(path string) string {
	switch {
	case FileExtensionIsOneOf(path, []string{ExtensionTs, ExtensionTsx, ExtensionJs, ExtensionJsx}):
		return ExtensionDts
	default:
		ext := GetAnyExtensionFromPath(path, nil, false)
		if ext != "" {
			return ".d" + ext + ExtensionTs
		}
		return ExtensionDts
	}
}

// ChangeAnyExtension changes the extension of a path to the provided extension if it has one of the provided extensions.
//
// ChangeAnyExtension("/path/to/file.ext", ".js", ".ext") === "/path/to/file.js"
// ChangeAnyExtension("/path/to/file.ext", ".js", ".ts") === "/path/to/file.ext"
// ChangeAnyExtension("/path/to/file.ext", ".js", [".ext", ".ts"]) === "/path/to/file.js"
func ChangeAnyExtension(path string, ext string, extensions []string, ignoreCase bool) string {
	pathext := GetAnyExtensionFromPath(path, extensions, ignoreCase)
	if pathext != "" {
		result := path[:len(path)-len(pathext)]
		if ext == "" {
			return result
		}
		if strings.HasPrefix(ext, ".") {
			return result + ext
		}
		return result + "." + ext
	}
	return path
}

func ChangeExtension(path string, newExtension string) string {
	return ChangeAnyExtension(path, newExtension, extensionsToRemove, false /*ignoreCase*/)
}

// Like `changeAnyExtension`, but declaration file extensions are recognized
// and replaced starting from the `.d`.
//
//	changeAnyExtension("file.d.ts", ".js") === "file.d.js"
//	changeFullExtension("file.d.ts", ".js") === "file.js"
func ChangeFullExtension(path string, newExtension string) string {
	declarationExtension := GetDeclarationFileExtension(path)
	if declarationExtension != "" {
		ext := newExtension
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		return path[:len(path)-len(declarationExtension)] + ext
	}
	return ChangeExtension(path, newExtension)
}

func GetPossibleOriginalInputExtensionForExtension(path string) []string {
	// Handle any custom .d.x.tlua extension (e.g., .d.json.tlua -> .json)
	if ext := GetDeclarationFileExtension(path); ext != "" && ext != ExtensionDts {
		inner := ext[len(".d.") : len(ext)-len(ExtensionTs)]
		return []string{"." + inner}
	}
	return []string{ExtensionTsx, ExtensionTs, ExtensionJsx, ExtensionJs}
}
