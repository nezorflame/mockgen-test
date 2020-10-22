//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package $GOPACKAGE -destination mocks_test.go github.com/nezorflame/mockgen-test/internal/client Producer

package service
