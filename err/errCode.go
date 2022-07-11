package err

const (
	E_NONE   errCode = 0
	E_SOME   errCode = 1000
	E_DB     errCode = 2000
	E_WORKER errCode = 3000
	E_HTTP   errCode = 4000
	E_UNDEF  errCode = 9000
)

var msgs = map[errCode]string{
	E_NONE:   "none",
	E_SOME:   "some",
	E_DB:     "with db",
	E_WORKER: "with worker",
	E_HTTP:   "with http",
	E_UNDEF:  "uncown error",
}
