//go:build !race
// +build !race

package session

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/argoproj/argo-cd/v3/util/settings"
)

func TestRandomPasswordVerificationDelay(t *testing.T) {
	// !race:
	// `SessionManager.VerifyUsernamePassword` uses bcrypt to prevent against time-based attacks
	// and verify the hashed password; however this is a CPU intensive algorithm that is made
	// significantly slower due to data race detection being enabled, which breaks through
	// the maximum time limit required by `TestRandomPasswordVerificationDelay`.

	var sleptFor time.Duration
	settingsMgr := settings.NewSettingsManager(t.Context(), getKubeClient(t, "password", true), "argocd")
	mgr := newSessionManager(settingsMgr, getProjLister(), NewUserStateStorage(nil))
	mgr.verificationDelayNoiseEnabled = true
	mgr.sleep = func(d time.Duration) {
		sleptFor = d
	}
	for i := 0; i < 10; i++ {
		sleptFor = 0
		start := time.Now()
		require.NoError(t, mgr.VerifyUsernamePassword("admin", "password"))
		totalDuration := time.Since(start) + sleptFor
		assert.GreaterOrEqual(t, totalDuration.Nanoseconds(), verificationDelayNoiseMin.Nanoseconds())
		assert.LessOrEqual(t, totalDuration.Nanoseconds(), verificationDelayNoiseMax.Nanoseconds())
	}
}
