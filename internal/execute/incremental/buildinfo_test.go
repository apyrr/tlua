package incremental

import (
	"testing"

	"github.com/apyrr/tlua/internal/diagnostics"
	"gotest.tools/v3/assert"
)

// Development builds share one version string, so a build info written by an
// earlier build passes the version check. A stored diagnostic whose message key
// this compiler no longer has (a reworded message) cannot be replayed -- it used
// to panic with "Unknown diagnostic message" -- so its file is checked again.
// The check reaches keys nested in message chains and related information.
func TestHasOnlyKnownMessages(t *testing.T) {
	t.Parallel()
	known := diagnostics.Identifier_expected.Key()
	nested := func(key diagnostics.Key) []*BuildInfoDiagnostic {
		return []*BuildInfoDiagnostic{{
			MessageKey:         known,
			MessageChain:       []*BuildInfoDiagnostic{{MessageKey: known}},
			RelatedInformation: []*BuildInfoDiagnostic{{MessageKey: key}},
		}}
	}
	assert.Assert(t, hasOnlyKnownMessages(nested(known)))
	assert.Assert(t, !hasOnlyKnownMessages(nested("A_message_that_was_reworded_9999")))
}
