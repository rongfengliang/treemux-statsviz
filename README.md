#  statsviz for treemux


## how to using


```code
go get github.com/rongfengliang/statsviz-pprof
router := treemux.New()
pprof.RouterpprofRegister(router)

view result:

/debug/statsviz

```