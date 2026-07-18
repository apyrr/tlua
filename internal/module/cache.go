package module

import (
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/packagejson"
)

type ModeAwareCache[T any] map[ModeAwareCacheKey]T

type moduleResolutionCacheKey struct {
	containingDirectory string
	moduleName          string
	resolutionMode      core.ResolutionMode
	redirectConfigName  string
}

type moduleResolutionCache struct {
	cache collections.SyncMap[moduleResolutionCacheKey, *ResolvedModule]
}

func (c *moduleResolutionCache) Get(key moduleResolutionCacheKey) (*ResolvedModule, bool) {
	return c.cache.Load(key)
}

func (c *moduleResolutionCache) Set(key moduleResolutionCacheKey, value *ResolvedModule) {
	c.cache.LoadOrStore(key, value)
}

type typeRefDirectiveResolutionCacheKey struct {
	containingDirectory             string
	typeReferenceName               string
	resolutionMode                  core.ResolutionMode
	redirectConfigName              string
	fromInferredTypesContainingFile bool
}

type typeRefDirectiveResolutionCache struct {
	cache collections.SyncMap[typeRefDirectiveResolutionCacheKey, *ResolvedTypeReferenceDirective]
}

func (c *typeRefDirectiveResolutionCache) Get(key typeRefDirectiveResolutionCacheKey) (*ResolvedTypeReferenceDirective, bool) {
	return c.cache.Load(key)
}

func (c *typeRefDirectiveResolutionCache) Set(key typeRefDirectiveResolutionCacheKey, value *ResolvedTypeReferenceDirective) {
	c.cache.Store(key, value)
}

type caches struct {
	packageJsonInfoCache *packagejson.InfoCache

	moduleResolutionCache           moduleResolutionCache
	typeRefDirectiveResolutionCache typeRefDirectiveResolutionCache
}

func newCaches(
	currentDirectory string,
	useCaseSensitiveFileNames bool,
	options *core.CompilerOptions,
) caches {
	return caches{
		packageJsonInfoCache: packagejson.NewInfoCache(currentDirectory, useCaseSensitiveFileNames),
	}
}

func getRedirectConfigName(redirect ResolvedProjectReference) string {
	if redirect == nil {
		return ""
	}
	return redirect.ConfigName()
}
