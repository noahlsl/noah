package resultx

import (
	"github.com/noahlsl/noah/consts"
	"sync"
)

var (
	defaultErrCode = consts.ErrCodeSysBadRequest
	defaultErr     = consts.ErrSysBadRequest
)

type ErrManger struct {
	sync.RWMutex
	errMap  map[string]map[int]error
	codeMap map[string]map[string]int
}

func NewErrManger() *ErrManger {

	errMap := make(map[string]map[int]error)
	codeMap := make(map[string]map[string]int)

	cnMap := map[int]error{
		consts.ErrCodeSysBadRequest:   consts.ErrSysBadRequest,
		consts.ErrCodeSysTokenExpired: consts.ErrSysTokenExpired,
		consts.ErrCodeSysAuthFailed:   consts.ErrSysAuthFailed,
		consts.ErrCodeRequestLimit:    consts.ErrRequestLimit,
		consts.ErrCodeTimeout:         consts.ErrTimeout,
		consts.ErrCodeIPLimit:         consts.ErrIPLimit,
		consts.ErrCodeImageSizeLimit:  consts.ErrImageSizeLimit,
		consts.ErrCodeImageSuffix:     consts.ErrImageSuffix,
	}

	errMap[consts.CN] = cnMap

	for lang, m := range errMap {
		codeMap[lang] = make(map[string]int)
		langMap := map[string]int{}
		for e, c := range m {
			langMap[c.Error()] = e
		}
		codeMap[lang] = langMap
	}

	return &ErrManger{
		errMap:  errMap,
		codeMap: codeMap,
	}
}

func (e *ErrManger) GetCode(err error, lang string) int {
	e.RLock()
	defer e.RUnlock()

	if lang == "" {
		lang = consts.CN
	}

	m, ok := e.codeMap[lang]
	if !ok {
		return defaultErrCode
	}

	code, ok := m[err.Error()]
	if !ok {
		return defaultErrCode
	}

	return code
}

func (e *ErrManger) GetErr(code int, lang string) error {
	e.Lock()
	defer e.Unlock()

	if lang == "" {
		lang = consts.CN
	}

	m, ok := e.errMap[lang]
	if !ok {
		return defaultErr
	}

	er, ok := m[code]
	if !ok {
		return defaultErr
	}

	return er
}
