Test name: `TestSanitizedStackTraceDefeatsVSCodeGenericSecretRegex`

# Unsanitized input:

````
goroutine 7 [running]:
runtime/debug.Stack()
	runtime/debug/stack.go:26 +0x5e
github.com/apyrr/tlua/internal/ls.(*LanguageService).getSignatureHelp(0x1)
	github.com/apyrr/tlua/internal/ls/signature.go:42 +0x10
github.com/apyrr/tlua/internal/ls.LookupKey(0x2)
	github.com/apyrr/tlua/internal/ls/keys.go:7 +0x10
github.com/apyrr/tlua/internal/ls.validateToken(0x3)
	github.com/apyrr/tlua/internal/ls/token.go:9 +0x10
github.com/apyrr/tlua/internal/ls.signRequest(0x4)
	github.com/apyrr/tlua/internal/ls/sig.go:11 +0x10
github.com/apyrr/tlua/internal/ls.setPwd(0x5)
	github.com/apyrr/tlua/internal/ls/pwd.go:13 +0x10
````

# Sanitized output:

````
(REDACTED FRAME)
	(REDACTED FRAME)
tlua|>internal|>ls.(*LanguageService).getSignatureHelp()
	tlua|>internal|>ls|>signatureX_X.go:42
tlua|>internal|>ls.LookupKeyX_X()
	tlua|>internal|>ls|>keys.go:7
tlua|>internal|>ls.validateTokenX_X()
	tlua|>internal|>ls|>tokenX_X.go:9
tlua|>internal|>ls.signRequest()
	tlua|>internal|>ls|>sigX_X.go:11
tlua|>internal|>ls.setPwdX_X()
	tlua|>internal|>ls|>pwdX_X.go:13
````
